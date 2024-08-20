package svc_impls

import (
	"context"
	"github.com/lysand-org/versia-go/internal/repository"
	"github.com/lysand-org/versia-go/internal/service"

	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/internal/entity"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

var _ service.FollowService = (*FollowServiceImpl)(nil)

type FollowServiceImpl struct {
	federationService service.FederationService

	repositories repository.Manager

	telemetry *unitel.Telemetry
	log       logr.Logger
}

func NewFollowServiceImpl(federationService service.FederationService, repositories repository.Manager, telemetry *unitel.Telemetry, log logr.Logger) *FollowServiceImpl {
	return &FollowServiceImpl{
		federationService: federationService,
		repositories:      repositories,
		telemetry:         telemetry,
		log:               log,
	}
}

func (i FollowServiceImpl) NewFollow(ctx context.Context, follower, followee *entity.User) (*entity.Follow, error) {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/FollowServiceImpl.NewFollow").
		AddAttribute("follower", follower.URI).
		AddAttribute("followee", followee.URI)
	defer s.End()
	ctx = s.Context()

	f, err := i.repositories.Follows().Follow(ctx, follower, followee)
	if err != nil {
		i.log.Error(err, "Failed to create follow", "follower", follower.ID, "followee", followee.ID)
		return nil, err
	}

	s.AddAttribute("followID", f.URI).
		AddAttribute("followURI", f.URI)

	i.log.V(2).Info("Created follow", "follower", follower.ID, "followee", followee.ID)

	return f, nil
}

func (i FollowServiceImpl) GetFollow(ctx context.Context, id uuid.UUID) (*entity.Follow, error) {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/FollowServiceImpl.GetFollow").
		AddAttribute("followID", id)
	defer s.End()
	ctx = s.Context()

	f, err := i.repositories.Follows().GetByID(ctx, id)
	if err != nil {
		return nil, err
	} else if f != nil {
		s.AddAttribute("followURI", f.URI)
	}

	return f, nil
}

func (i FollowServiceImpl) ImportLysandFollow(ctx context.Context, lFollow *lysand.Follow) (*entity.Follow, error) {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/FollowServiceImpl.ImportLysandFollow").
		AddAttribute("uri", lFollow.URI.String())
	defer s.End()
	ctx = s.Context()

	var f *entity.Follow
	if err := i.repositories.Atomic(ctx, func(ctx context.Context, tx repository.Manager) error {
		follower, err := i.repositories.Users().Resolve(ctx, lFollow.Author)
		if err != nil {
			return err
		}
		s.AddAttribute("follower", follower.URI)

		followee, err := i.repositories.Users().Resolve(ctx, lFollow.Followee)
		if err != nil {
			return err
		}
		s.AddAttribute("followee", followee.URI)

		f, err = i.repositories.Follows().Follow(ctx, follower, followee)
		return err
	}); err != nil {
		return nil, err
	}

	s.AddAttribute("followID", f.ID).
		AddAttribute("followURI", f.URI)

	return f, nil
}
