package svc_impls

import (
	"context"
	"github.com/lysand-org/versia-go/internal/repository"
	"github.com/lysand-org/versia-go/internal/service"
	"slices"

	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/internal/api_schema"
	"github.com/lysand-org/versia-go/internal/entity"
	"github.com/lysand-org/versia-go/internal/tasks"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

var _ service.NoteService = (*NoteServiceImpl)(nil)

type NoteServiceImpl struct {
	federationService service.FederationService
	taskService       service.TaskService

	repositories repository.Manager

	telemetry *unitel.Telemetry

	log logr.Logger
}

func NewNoteServiceImpl(federationService service.FederationService, taskService service.TaskService, repositories repository.Manager, telemetry *unitel.Telemetry, log logr.Logger) *NoteServiceImpl {
	return &NoteServiceImpl{
		federationService: federationService,
		taskService:       taskService,
		repositories:      repositories,
		telemetry:         telemetry,
		log:               log,
	}
}

func (i NoteServiceImpl) CreateNote(ctx context.Context, req api_schema.CreateNoteRequest) (*entity.Note, error) {
	s := i.telemetry.StartSpan(ctx, "function", "service/svc_impls.NoteServiceImpl.CreateNote")
	defer s.End()
	ctx = s.Context()

	var n *entity.Note

	if err := i.repositories.Atomic(ctx, func(ctx context.Context, tx repository.Manager) error {
		// FIXME: Use the user that created the note
		author, err := tx.Users().GetLocalByID(ctx, uuid.MustParse("b6f4bcb5-ac5a-4a87-880a-c7f88f58a172"))
		if err != nil {
			return err
		}
		if author == nil {
			return api_schema.ErrBadRequest(map[string]any{"reason": "author not found"})
		}

		mentionedUsers, err := i.repositories.Users().ResolveMultiple(ctx, req.Mentions)
		if err != nil {
			return err
		}

		if slices.ContainsFunc(mentionedUsers, func(u *entity.User) bool { return u.ID == author.ID }) {
			return api_schema.ErrBadRequest(map[string]any{"reason": "cannot mention self"})
		}

		n, err = tx.Notes().NewNote(ctx, author, req.Content, mentionedUsers)
		if err != nil {
			return err
		}

		if err := i.taskService.ScheduleTask(ctx, tasks.FederateNote, tasks.FederateNoteData{NoteID: n.ID}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return n, nil
}

func (i NoteServiceImpl) GetNote(ctx context.Context, id uuid.UUID) (*entity.Note, error) {
	s := i.telemetry.StartSpan(ctx, "function", "service/svc_impls.NoteServiceImpl.GetUserByID")
	defer s.End()
	ctx = s.Context()

	return i.repositories.Notes().GetByID(ctx, id)
}

func (i NoteServiceImpl) ImportLysandNote(ctx context.Context, lNote *lysand.Note) (*entity.Note, error) {
	s := i.telemetry.StartSpan(ctx, "function", "service/svc_impls.NoteServiceImpl.ImportLysandNote")
	defer s.End()
	ctx = s.Context()

	return i.repositories.Notes().ImportLysandNote(ctx, lNote)
}
