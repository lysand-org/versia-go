package taskqueue

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"sync"
	"time"

	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type taskWrapper struct {
	Task       Task              `json:"task"`
	EnqueuedAt time.Time         `json:"enqueuedAt"`
	TraceInfo  map[string]string `json:"traceInfo"`
}

func (c *Client) newTaskWrapper(ctx context.Context, task Task) taskWrapper {
	traceInfo := make(map[string]string)
	c.telemetry.InjectIntoMap(ctx, traceInfo)

	return taskWrapper{
		Task:       task,
		EnqueuedAt: time.Now(),
		TraceInfo:  traceInfo,
	}
}

type Task struct {
	ID      string
	Type    string
	Payload json.RawMessage
}

func NewTask(type_ string, payload any) (Task, error) {
	id := uuid.New()

	d, err := json.Marshal(payload)
	if err != nil {
		return Task{}, err
	}

	return Task{
		ID:      id.String(),
		Type:    type_,
		Payload: d,
	}, nil
}

type Handler func(ctx context.Context, task Task) error

type Client struct {
	name     string
	subject  string
	handlers map[string][]Handler

	nc *nats.Conn
	js jetstream.JetStream
	s  jetstream.Stream

	stopCh    chan struct{}
	closeOnce func()

	telemetry *unitel.Telemetry
	log       logr.Logger
}

func NewClient(ctx context.Context, name string, natsClient *nats.Conn, telemetry *unitel.Telemetry, log logr.Logger) (*Client, error) {
	js, err := jetstream.New(natsClient)
	if err != nil {
		return nil, err
	}

	s, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:              name,
		Subjects:          []string{name + ".*"},
		MaxConsumers:      -1,
		MaxMsgs:           -1,
		Discard:           jetstream.DiscardOld,
		MaxMsgsPerSubject: -1,
		Storage:           jetstream.FileStorage,
		Compression:       jetstream.S2Compression,
		AllowDirect:       true,
	})
	if errors.Is(err, nats.ErrStreamNameAlreadyInUse) {
		s, err = js.Stream(ctx, name)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	stopCh := make(chan struct{})

	c := &Client{
		name:    name,
		subject: name + ".tasks",

		handlers: map[string][]Handler{},

		stopCh: stopCh,
		closeOnce: sync.OnceFunc(func() {
			close(stopCh)
		}),

		nc: natsClient,
		js: js,
		s:  s,

		telemetry: telemetry,
		log:       log,
	}

	return c, nil
}

func (c *Client) Close() {
	c.closeOnce()
	c.nc.Close()
}

func (c *Client) Submit(ctx context.Context, task Task) error {
	s := c.telemetry.StartSpan(ctx, "queue.publish", "taskqueue/Client.Submit").
		AddAttribute("messaging.destination.name", c.subject)
	defer s.End()
	ctx = s.Context()

	s.AddAttribute("jobID", task.ID)

	data, err := json.Marshal(c.newTaskWrapper(ctx, task))
	if err != nil {
		return err
	}

	s.AddAttribute("messaging.message.body.size", len(data))

	msg, err := c.js.PublishMsg(ctx, &nats.Msg{Subject: c.subject, Data: data})
	if err != nil {
		return err
	}
	c.log.V(1).Info("submitted task", "id", task.ID, "type", task.Type, "sequence", msg.Sequence)

	s.AddAttribute("messaging.message.id", msg.Sequence)

	return nil
}

func (c *Client) RegisterHandler(type_ string, handler Handler) {
	c.log.V(2).Info("registering handler", "type", type_)

	if _, ok := c.handlers[type_]; !ok {
		c.handlers[type_] = []Handler{}
	}
	c.handlers[type_] = append(c.handlers[type_], handler)
}

func (c *Client) Start(ctx context.Context) error {
	c.log.Info("starting")

	sub, err := c.js.CreateConsumer(ctx, c.name, jetstream.ConsumerConfig{
		// TODO: set name properly
		Name:          "versia-go",
		Durable:       "versia-go",
		DeliverPolicy: jetstream.DeliverAllPolicy,
		ReplayPolicy:  jetstream.ReplayInstantPolicy,
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: c.subject,
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

	go func() {
		for {
			msg, err := m.Next()
			if err != nil {
				if errors.Is(err, jetstream.ErrMsgIteratorClosed) {
					c.log.Info("stopping")
					return
				}

				c.log.Error(err, "failed to get next message")
				break
			}

			if err := c.handleTask(ctx, msg); err != nil {
				c.log.Error(err, "failed to handle task")
				break
			}
		}
	}()
	go func() {
		<-c.stopCh
		m.Drain()
	}()

	return nil
}

func (c *Client) handleTask(ctx context.Context, msg jetstream.Msg) error {
	msgMeta, err := msg.Metadata()
	if err != nil {
		return err
	}

	data := msg.Data()

	var w taskWrapper
	if err := json.Unmarshal(data, &w); err != nil {
		if err := msg.Nak(); err != nil {
			c.log.Error(err, "failed to nak message")
		}

		return err
	}

	s := c.telemetry.StartSpan(
		context.Background(),
		"queue.process",
		"taskqueue/Client.handleTask",
		c.telemetry.ContinueFromMap(w.TraceInfo),
	).
		AddAttribute("messaging.destination.name", c.subject).
		AddAttribute("messaging.message.id", msgMeta.Sequence.Stream).
		AddAttribute("messaging.message.retry.count", msgMeta.NumDelivered).
		AddAttribute("messaging.message.body.size", len(data)).
		AddAttribute("messaging.message.receive.latency", time.Since(w.EnqueuedAt).Milliseconds())
	defer s.End()
	ctx = s.Context()

	handlers, ok := c.handlers[w.Task.Type]
	if !ok {
		c.log.V(1).Info("no handler for task", "type", w.Task.Type)
		return msg.Nak()
	}

	var errs CombinedError
	for _, handler := range handlers {
		if err := handler(ctx, w.Task); err != nil {
			c.log.Error(err, "handler failed", "type", w.Task.Type)
			errs.Errors = append(errs.Errors, err)
		}
	}

	if len(errs.Errors) > 0 {
		if err := msg.Nak(); err != nil {
			c.log.Error(err, "failed to nak message")
			errs.Errors = append(errs.Errors, err)
		}

		return errs
	}

	return msg.Ack()
}

type CombinedError struct {
	Errors []error
}

func (e CombinedError) Error() string {
	sb := strings.Builder{}
	sb.WriteRune('[')
	for i, err := range e.Errors {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(err.Error())
	}
	sb.WriteRune(']')
	return sb.String()
}
