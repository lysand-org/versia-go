package taskqueue

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

type Client struct {
	name    string
	subject string

	nc *nats.Conn
	js jetstream.JetStream
	s  jetstream.Stream

	stopCh    chan struct{}
	closeOnce func()

	telemetry *unitel.Telemetry
	log       logr.Logger
}

func NewClient(streamName string, natsClient *nats.Conn, telemetry *unitel.Telemetry, log logr.Logger) (*Client, error) {
	js, err := jetstream.New(natsClient)
	if err != nil {
		return nil, err
	}

	return &Client{
		name: streamName,

		js: js,

		telemetry: telemetry,
		log:       log,
	}, nil
}

func (c *Client) Set(ctx context.Context, name string) (*Set, error) {
	streamName := fmt.Sprintf("%s_%s", c.name, name)

	s, err := c.js.CreateStream(ctx, jetstream.StreamConfig{
		Name:              streamName,
		MaxConsumers:      -1,
		MaxMsgs:           -1,
		Discard:           jetstream.DiscardOld,
		MaxMsgsPerSubject: -1,
		Storage:           jetstream.FileStorage,
		Compression:       jetstream.S2Compression,
		AllowDirect:       true,
	})
	if errors.Is(err, nats.ErrStreamNameAlreadyInUse) {
		s, err = c.js.Stream(ctx, streamName)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	return &Set{
		handlers: make(map[string][]TaskHandler),

		streamName: streamName,
		c:          c,
		s:          s,
		log:        c.log.WithName(fmt.Sprintf("taskset(%s)", name)),
		telemetry:  c.telemetry,
	}, nil
}
