package repo_impls

import (
	"context"
	"errors"
	"fmt"
	"github.com/versia-pub/versia-go/internal/repository"

	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/versia-pub/versia-go/ent"
	"github.com/versia-pub/versia-go/ent/follow"
	"github.com/versia-pub/versia-go/ent/predicate"
	"github.com/versia-pub/versia-go/ent/user"
	"github.com/versia-pub/versia-go/internal/entity"
	"github.com/versia-pub/versia-go/internal/utils"
)

var ErrFollowAlreadyExists = errors.New("follow already exists")

var _ repository.FollowRepository = (*FollowRepositoryImpl)(nil)

type FollowRepositoryImpl struct {
	db        *ent.Client
	log       logr.Logger
	telemetry *unitel.Telemetry
}

func NewFollowRepositoryImpl(db *ent.Client, log logr.Logger, telemetry *unitel.Telemetry) repository.FollowRepository {
	return &FollowRepositoryImpl{
		db:        db,
		log:       log,
		telemetry: telemetry,
	}
}

func (i FollowRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*entity.Follow, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/FollowRepositoryImpl.GetByID").
		AddAttribute("followID", id)
	defer s.End()
	ctx = s.Context()

	f, err := i.db.Follow.Query().
		Where(follow.ID(id)).
		WithFollowee().
		WithFollower().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	s.AddAttribute("follower", f.Edges.Follower.URI).
		AddAttribute("followee", f.Edges.Followee.URI).
		AddAttribute("followURI", f.URI)

	return entity.NewFollow(f)
}

func (i FollowRepositoryImpl) Follow(ctx context.Context, follower, followee *entity.User) (*entity.Follow, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/FollowRepositoryImpl.Follow").
		AddAttribute("follower", follower.URI).
		AddAttribute("followee", followee.URI)
	defer s.End()
	ctx = s.Context()

	fid := uuid.New()

	fid, err := i.db.Follow.Create().
		SetID(fid).
		SetIsRemote(false).
		SetURI(utils.UserAPIURL(fid).String()).
		SetStatus(follow.StatusPending).
		SetFollower(follower.User).
		SetFollowee(followee.User).
		OnConflictColumns(follow.FollowerColumn, follow.FolloweeColumn).
		UpdateStatus().
		ID(ctx)
	if err != nil {
		if !ent.IsConstraintError(err) {
			return nil, err
		}

		return nil, ErrFollowAlreadyExists
	}

	s.AddAttribute("followID", fid)

	f, err := i.db.Follow.Query().
		Where(follow.ID(fid)).
		WithFollowee().
		WithFollower().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	s.AddAttribute("followURI", f.URI)

	return entity.NewFollow(f)
}

func (i FollowRepositoryImpl) Unfollow(ctx context.Context, follower, followee *entity.User) error {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/FollowRepositoryImpl.Unfollow").
		AddAttribute("follower", follower.URI).
		AddAttribute("followee", followee.URI)
	defer s.End()
	ctx = s.Context()

	n, err := i.db.Follow.Delete().
		Where(matchFollowUsers(follower, followee)).
		Exec(ctx)
	if err != nil {
		s.CaptureError(err)
	} else {
		s.AddAttribute("deleted", n).
			CaptureMessage(fmt.Sprintf("Deleted %d follow(s)", n))
	}

	return nil
}

func (i FollowRepositoryImpl) AcceptFollow(ctx context.Context, follower, followee *entity.User) error {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/FollowRepositoryImpl.AcceptFollow").
		AddAttribute("follower", follower.URI).
		AddAttribute("followee", followee.URI)
	defer s.End()
	ctx = s.Context()

	n, err := i.db.Follow.Update().
		Where(matchFollowUsers(follower, followee), follow.StatusEQ(follow.StatusPending)).
		SetStatus(follow.StatusAccepted).
		Save(ctx)
	if err != nil {
		s.CaptureError(err)
	} else {
		s.CaptureMessage(fmt.Sprintf("Accepted %d follow(s)", n))
	}

	return err
}

func (i FollowRepositoryImpl) RejectFollow(ctx context.Context, follower, followee *entity.User) error {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/FollowRepositoryImpl.RejectFollow").
		AddAttribute("follower", follower.URI).
		AddAttribute("followee", followee.URI)
	defer s.End()
	ctx = s.Context()

	n, err := i.db.Follow.Delete().
		Where(follow.And(matchFollowUsers(follower, followee), follow.StatusEQ(follow.StatusPending))).
		Exec(ctx)
	if err != nil {
		s.CaptureError(err)
	} else {
		s.CaptureMessage(fmt.Sprintf("Deleted %d follow(s)", n))
	}

	return err
}

func matchFollowUsers(follower, followee *entity.User) predicate.Follow {
	return follow.And(
		follow.HasFollowerWith(
			user.ID(follower.ID), user.ID(followee.ID),
		),
		follow.HasFolloweeWith(
			user.ID(follower.ID), user.ID(followee.ID),
		),
	)
}
