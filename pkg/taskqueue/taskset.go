package taskqueue

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type TaskHandler = func(ctx context.Context, task Task) error

type Set struct {
	handlers map[string][]TaskHandler

	streamName string
	c          *Client
	s          jetstream.Stream
	log        logr.Logger
	telemetry  *unitel.Telemetry
}

func (t *Set) RegisterHandler(type_ string, handler TaskHandler) {
	t.log.V(2).Info("Registering handler", "type", type_)

	if _, ok := t.handlers[type_]; !ok {
		t.handlers[type_] = []TaskHandler{}
	}
	t.handlers[type_] = append(t.handlers[type_], handler)
}

func (t *Set) Submit(ctx context.Context, task Task) error {
	s := t.telemetry.StartSpan(ctx, "queue.publish", "taskqueue/TaskSet.Submit").
		AddAttribute("messaging.destination.name", t.streamName)
	defer s.End()
	ctx = s.Context()

	s.AddAttribute("jobID", task.ID)

	data, err := json.Marshal(t.c.newTaskWrapper(ctx, task))
	if err != nil {
		return err
	}

	s.AddAttribute("messaging.message.body.size", len(data))

	// TODO: Refactor
	msg, err := t.c.js.PublishMsg(ctx, &nats.Msg{Subject: t.streamName, Data: data})
	if err != nil {
		return err
	}
	t.log.V(2).Info("Submitted task", "id", task.ID, "type", task.Type, "sequence", msg.Sequence)

	s.AddAttribute("messaging.message.id", msg.Sequence)

	return nil
}

func (t *Set) Consumer(name string) *Consumer {
	stopCh := make(chan struct{})
	stopOnce := sync.OnceFunc(func() {
		close(stopCh)
	})

	return &Consumer{
		stopCh:   stopCh,
		stopOnce: stopOnce,

		name:       name,
		streamName: t.streamName,
		telemetry:  t.telemetry,
		log:        t.log.WithName(fmt.Sprintf("consumer(%s)", name)),
		t:          t,
	}
}

type Consumer struct {
	stopCh   chan struct{}
	stopOnce func()

	name       string
	streamName string
	telemetry  *unitel.Telemetry
	log        logr.Logger
	t          *Set
}

func (c *Consumer) Close() {
	c.stopOnce()
}

func (c *Consumer) Start(ctx context.Context) error {
	c.log.Info("Starting consumer")

	sub, err := c.t.c.js.CreateConsumer(ctx, c.streamName, jetstream.ConsumerConfig{
		Durable:       c.name,
		DeliverPolicy: jetstream.DeliverAllPolicy,
		ReplayPolicy:  jetstream.ReplayInstantPolicy,
		AckPolicy:     jetstream.AckExplicitPolicy,
		MaxWaiting:    1,
		MaxAckPending: 1,
		HeadersOnly:   false,
		MemoryStorage: false,
	})
	if err != nil {
		return err
	}

	m, err := sub.Messages(jetstream.PullMaxMessages(1))
	if err != nil {
		return err
	}

	go c.handleMessages(m)

	go func() {
		<-ctx.Done()
		c.Close()
	}()

	go func() {
		<-c.stopCh
		m.Drain()
	}()

	return nil
}

func (c *Consumer) handleMessages(m jetstream.MessagesContext) {
	for {
		msg, err := m.Next()
		if err != nil {
			if errors.Is(err, jetstream.ErrMsgIteratorClosed) {
				c.log.Info("Stopping")
				return
			}

			c.log.Error(err, "Failed to get next message")
			break
		}

		if err := c.handleTask(msg); err != nil {
			c.log.Error(err, "Failed to handle task")
			break
		}
	}
}

func (c *Consumer) handleTask(msg jetstream.Msg) error {
	msgMeta, err := msg.Metadata()
	if err != nil {
		return err
	}

	data := msg.Data()

	var w taskWrapper
	if err := json.Unmarshal(data, &w); err != nil {
		if err := msg.Nak(); err != nil {
			c.log.Error(err, "Failed to nak message")
		}

		return err
	}

	s := c.telemetry.StartSpan(
		context.Background(),
		"queue.process",
		"taskqueue/Consumer.handleTask",
		c.telemetry.ContinueFromMap(w.TraceInfo),
	).
		AddAttribute("messaging.destination.name", msg.Subject()).
		AddAttribute("messaging.message.id", msgMeta.Sequence.Stream).
		AddAttribute("messaging.message.retry.count", msgMeta.NumDelivered).
		AddAttribute("messaging.message.body.size", len(data)).
		AddAttribute("messaging.message.receive.latency", time.Since(w.EnqueuedAt).Milliseconds())
	defer s.End()
	ctx := s.Context()

	handlers, ok := c.t.handlers[w.Task.Type]
	if !ok {
		c.log.V(2).Info("No handler for task", "type", w.Task.Type)
		return msg.Nak()
	}

	var errs CombinedError
	for _, handler := range handlers {
		if err := handler(ctx, w.Task); err != nil {
			c.log.Error(err, "Handler failed", "type", w.Task.Type)
			errs.Errors = append(errs.Errors, err)
		}
	}

	if len(errs.Errors) > 0 {
		if err := msg.Nak(); err != nil {
			c.log.Error(err, "Failed to nak message")
			errs.Errors = append(errs.Errors, err)
		}

		return errs
	}

	return msg.Ack()
}
