package tasks

import (
	"context"
	"encoding/json"
	"github.com/lysand-org/versia-go/internal/repository"
	"github.com/lysand-org/versia-go/internal/service"

	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/lysand-org/versia-go/pkg/taskqueue"
)

const (
	FederateNote   = "federate_note"
	FederateFollow = "federate_follow"
)

type Handler struct {
	federationService service.FederationService

	repositories repository.Manager

	telemetry *unitel.Telemetry
	log       logr.Logger
}

func NewHandler(federationService service.FederationService, repositories repository.Manager, telemetry *unitel.Telemetry, log logr.Logger) *Handler {
	return &Handler{
		federationService: federationService,

		repositories: repositories,

		telemetry: telemetry,
		log:       log,
	}
}

func (t *Handler) Register(tq *taskqueue.Client) {
	tq.RegisterHandler(FederateNote, parse(t.FederateNote))
	tq.RegisterHandler(FederateFollow, parse(t.FederateFollow))
}

func parse[T any](handler func(context.Context, T) error) func(context.Context, taskqueue.Task) error {
	return func(ctx context.Context, task taskqueue.Task) error {
		var data T
		if err := json.Unmarshal(task.Payload, &data); err != nil {
			return err
		}

		return handler(ctx, data)
	}
}
