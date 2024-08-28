package task_impls

import (
	"context"

	"github.com/versia-pub/versia-go/internal/entity"
	"github.com/versia-pub/versia-go/internal/repository"
	"github.com/versia-pub/versia-go/internal/service"
	"github.com/versia-pub/versia-go/internal/task"
	task_dtos "github.com/versia-pub/versia-go/internal/task/dtos"
	"github.com/versia-pub/versia-go/internal/utils"

	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/versia-pub/versia-go/pkg/taskqueue"
)

var _ task.Handler = (*NoteHandler)(nil)

type NoteHandler struct {
	federationService service.FederationService

	repositories repository.Manager

	telemetry *unitel.Telemetry
	log       logr.Logger
	set       *taskqueue.Set
}

func NewNoteHandler(federationService service.FederationService, repositories repository.Manager, telemetry *unitel.Telemetry, log logr.Logger) *NoteHandler {
	return &NoteHandler{
		federationService: federationService,

		repositories: repositories,

		telemetry: telemetry,
		log:       log,
	}
}

func (t *NoteHandler) Start(ctx context.Context) error {
	consumer := t.set.Consumer("note-handler")

	return consumer.Start(ctx)
}

func (t *NoteHandler) Register(s *taskqueue.Set) {
	t.set = s
	s.RegisterHandler(task_dtos.FederateNote, utils.ParseTask(t.FederateNote))
}

func (t *NoteHandler) Submit(ctx context.Context, task taskqueue.Task) error {
	s := t.telemetry.StartSpan(ctx, "function", "task_impls/NoteHandler.Submit")
	defer s.End()
	ctx = s.Context()

	return t.set.Submit(ctx, task)
}

func (t *NoteHandler) FederateNote(ctx context.Context, data task_dtos.FederateNoteData) error {
	s := t.telemetry.StartSpan(ctx, "function", "task_impls/NoteHandler.FederateNote")
	defer s.End()
	ctx = s.Context()

	var n *entity.Note
	if err := t.repositories.Atomic(ctx, func(ctx context.Context, tx repository.Manager) error {
		var err error
		n, err = tx.Notes().GetByID(ctx, data.NoteID)
		if err != nil {
			return err
		}
		if n == nil {
			t.log.V(-1).Info("Could not find note", "id", data.NoteID)
			return nil
		}

		for _, uu := range n.Mentions {
			if !uu.IsRemote {
				t.log.V(2).Info("User is not remote", "user", uu.ID)
				continue
			}

			res, err := t.federationService.SendToInbox(ctx, n.Author, &uu, n.ToVersia())
			if err != nil {
				t.log.Error(err, "Failed to send note to remote user", "res", res, "user", uu.ID)
			} else {
				t.log.V(2).Info("Sent note to remote user", "res", res, "user", uu.ID)
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
