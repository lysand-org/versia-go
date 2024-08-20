package svc_impls

import (
	"context"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/internal/repository"
	"github.com/lysand-org/versia-go/internal/service"

	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/lysand-org/versia-go/ent"
	"github.com/lysand-org/versia-go/ent/user"
	"github.com/lysand-org/versia-go/internal/api_schema"
	"github.com/lysand-org/versia-go/internal/entity"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

var _ service.InboxService = (*InboxServiceImpl)(nil)

type InboxServiceImpl struct {
	federationService service.FederationService

	repositories repository.Manager

	telemetry *unitel.Telemetry
	log       logr.Logger
}

func NewInboxService(federationService service.FederationService, repositories repository.Manager, telemetry *unitel.Telemetry, log logr.Logger) *InboxServiceImpl {
	return &InboxServiceImpl{
		federationService: federationService,

		repositories: repositories,

		telemetry: telemetry,
		log:       log,
	}
}

func (i InboxServiceImpl) WithRepositories(repositories repository.Manager) service.InboxService {
	return NewInboxService(i.federationService, repositories, i.telemetry, i.log)
}

func (i InboxServiceImpl) Handle(ctx context.Context, obj any, userId uuid.UUID) error {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/InboxServiceImpl.Handle")
	defer s.End()
	ctx = s.Context()

	return i.repositories.Atomic(ctx, func(ctx context.Context, tx repository.Manager) error {
		i := i.WithRepositories(tx).(*InboxServiceImpl)

		u, err := i.repositories.Users().GetLocalByID(ctx, userId)
		if err != nil {
			i.log.Error(err, "Failed to get user", "id", userId)

			return api_schema.ErrInternalServerError(nil)
		}
		if u == nil {
			return api_schema.ErrNotFound(map[string]any{
				"id": userId,
			})
		}

		// TODO: Implement more types
		switch o := obj.(type) {
		case lysand.Note:
			i.log.Info("Received note", "note", o)
			if err := i.handleNote(ctx, o, u); err != nil {
				i.log.Error(err, "Failed to handle note", "note", o)
				return err
			}

		case lysand.Patch:
			i.log.Info("Received patch", "patch", o)
		case lysand.Follow:
			if err := i.handleFollow(ctx, o, u); err != nil {
				i.log.Error(err, "Failed to handle follow", "follow", o)
				return err
			}
		case lysand.Undo:
			i.log.Info("Received undo", "undo", o)
		default:
			i.log.Info("Unimplemented object type", "object", obj)
			return api_schema.ErrNotImplemented(nil)
		}

		return nil
	})
}

func (i InboxServiceImpl) handleFollow(ctx context.Context, o lysand.Follow, u *entity.User) error {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/InboxServiceImpl.handleFollow")
	defer s.End()
	ctx = s.Context()

	author, err := i.repositories.Users().Resolve(ctx, o.Author)
	if err != nil {
		i.log.Error(err, "Failed to resolve author", "author", o.Author)
		return err
	}

	f, err := i.repositories.Follows().Follow(ctx, author, u)
	if err != nil {
		// TODO: Handle constraint errors
		if ent.IsConstraintError(err) {
			i.log.Error(err, "Follow already exists", "user", user.ID, "author", author.ID)
			return nil
		}

		i.log.Error(err, "Failed to create follow", "user", user.ID, "author", author.ID)
		return err
	}

	switch u.PrivacyLevel {
	case user.PrivacyLevelPublic:
		if err := i.repositories.Follows().AcceptFollow(ctx, author, u); err != nil {
			i.log.Error(err, "Failed to accept follow", "user", user.ID, "author", author.ID)
			return err
		}

		if _, err := i.federationService.SendToInbox(ctx, u, author, f.ToLysandAccept()); err != nil {
			i.log.Error(err, "Failed to send follow accept to inbox", "user", user.ID, "author", author.ID)
			return err
		}

	case user.PrivacyLevelRestricted:
	case user.PrivacyLevelPrivate:
	}

	return nil
}

func (i InboxServiceImpl) handleNote(ctx context.Context, o lysand.Note, u *entity.User) error {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/InboxServiceImpl.handleNote")
	defer s.End()
	ctx = s.Context()

	author, err := i.repositories.Users().Resolve(ctx, o.Author)
	if err != nil {
		i.log.Error(err, "Failed to resolve author", "author", o.Author)
		return err
	}

	// TODO: Implement

	_ = author

	return nil
}
