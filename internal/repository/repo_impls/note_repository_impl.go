package repo_impls

import (
	"context"
	"github.com/versia-pub/versia-go/internal/repository"
	"github.com/versia-pub/versia-go/pkg/versia"

	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/versia-pub/versia-go/ent"
	"github.com/versia-pub/versia-go/ent/note"
	"github.com/versia-pub/versia-go/internal/entity"
	"github.com/versia-pub/versia-go/internal/utils"
)

var _ repository.NoteRepository = (*NoteRepositoryImpl)(nil)

type NoteRepositoryImpl struct {
	db        *ent.Client
	log       logr.Logger
	telemetry *unitel.Telemetry
}

func NewNoteRepositoryImpl(db *ent.Client, log logr.Logger, telemetry *unitel.Telemetry) repository.NoteRepository {
	return &NoteRepositoryImpl{
		db:        db,
		log:       log,
		telemetry: telemetry,
	}
}

func (i *NoteRepositoryImpl) NewNote(ctx context.Context, author *entity.User, content string, mentions []*entity.User) (*entity.Note, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/NoteRepositoryImpl.NewNote")
	defer s.End()
	ctx = s.Context()

	nid := uuid.New()

	n, err := i.db.Note.Create().
		SetID(nid).
		SetIsRemote(false).
		SetURI(utils.NoteAPIURL(nid).String()).
		SetAuthor(author.User).
		SetContent(content).
		AddMentions(utils.MapSlice(mentions, func(m *entity.User) *ent.User { return m.User })...).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	n, err = i.db.Note.Query().
		Where(note.ID(nid)).
		WithAuthor().
		WithMentions().
		Only(ctx)
	if err != nil {
		i.log.Error(err, "Failed to query author", "id", nid)
		return nil, err
	}

	return entity.NewNote(n)
}

func (i *NoteRepositoryImpl) ImportVersiaNote(ctx context.Context, lNote *versia.Note) (*entity.Note, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/NoteRepositoryImpl.ImportVersiaNote")
	defer s.End()
	ctx = s.Context()

	id, err := i.db.Note.Create().
		SetID(uuid.New()).
		SetIsRemote(true).
		SetURI(lNote.URI.String()).
		OnConflict().
		UpdateNewValues().
		ID(ctx)
	if err != nil {
		i.log.Error(err, "Failed to import note into database", "uri", lNote.URI)
		return nil, err
	}

	n, err := i.db.Note.Get(ctx, id)
	if err != nil {
		i.log.Error(err, "Failed to get imported note", "id", id, "uri", lNote.URI)
		return nil, err
	}

	i.log.V(2).Info("Imported note into database", "id", id, "uri", lNote.URI)

	return entity.NewNote(n)
}

func (i *NoteRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*entity.Note, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/NoteRepositoryImpl.LookupByIDOrUsername")
	defer s.End()
	ctx = s.Context()

	n, err := i.db.Note.Query().
		Where(note.ID(id)).
		WithAuthor().
		WithMentions().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			i.log.Error(err, "Failed to query user", "id", id)
			return nil, err
		}

		i.log.V(2).Info("User not found in DB", "id", id)

		return nil, nil
	}

	i.log.V(2).Info("User found in DB", "id", id)

	return entity.NewNote(n)
}
