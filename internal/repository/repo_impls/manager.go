package repo_impls

import (
	"context"
	"github.com/lysand-org/versia-go/internal/repository"

	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/lysand-org/versia-go/ent"
	"github.com/lysand-org/versia-go/internal/database"
)

type Factory[T any] func(db *ent.Client, log logr.Logger, telemetry *unitel.Telemetry) T

var _ repository.Manager = (*ManagerImpl)(nil)

type ManagerImpl struct {
	users   repository.UserRepository
	notes   repository.NoteRepository
	follows repository.FollowRepository

	uRFactory Factory[repository.UserRepository]
	nRFactory Factory[repository.NoteRepository]
	fRFactory Factory[repository.FollowRepository]

	db        *ent.Client
	log       logr.Logger
	telemetry *unitel.Telemetry
}

func NewManagerImpl(db *ent.Client, telemetry *unitel.Telemetry, log logr.Logger, userRepositoryFunc Factory[repository.UserRepository], noteRepositoryFunc Factory[repository.NoteRepository], followRepositoryFunc Factory[repository.FollowRepository]) *ManagerImpl {
	userRepository := userRepositoryFunc(db, log.WithName("users"), telemetry)
	noteRepository := noteRepositoryFunc(db, log.WithName("notes"), telemetry)
	followRepository := followRepositoryFunc(db, log.WithName("follows"), telemetry)

	return &ManagerImpl{
		users:   userRepository,
		notes:   noteRepository,
		follows: followRepository,

		uRFactory: userRepositoryFunc,
		nRFactory: noteRepositoryFunc,
		fRFactory: followRepositoryFunc,

		db:        db,
		log:       log,
		telemetry: telemetry,
	}
}

func (i *ManagerImpl) withDB(db *ent.Client) *ManagerImpl {
	return NewManagerImpl(db, i.telemetry, i.log, i.uRFactory, i.nRFactory, i.fRFactory)
}

func (i *ManagerImpl) Atomic(ctx context.Context, fn func(ctx context.Context, tx repository.Manager) error) error {
	s := i.telemetry.StartSpan(ctx, "function", "repository/repo_impls.ManagerImpl.Atomic")
	defer s.End()
	ctx = s.Context()

	tx, err := database.BeginTx(ctx, i.db, i.telemetry)
	if err != nil {
		return err
	}
	defer func(tx *database.Tx) {
		err := tx.Finish()
		if err != nil {
			i.log.Error(err, "Failed to finish transaction")
		}
	}(tx)

	if err := fn(ctx, i.withDB(tx.Client())); err != nil {
		return err
	}

	tx.MarkForCommit()

	return tx.Finish()
}

func (i *ManagerImpl) Users() repository.UserRepository {
	return i.users
}

func (i *ManagerImpl) Notes() repository.NoteRepository {
	return i.notes
}

func (i *ManagerImpl) Follows() repository.FollowRepository {
	return i.follows
}
