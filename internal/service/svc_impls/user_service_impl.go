package svc_impls

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"github.com/lysand-org/versia-go/internal/api_schema"
	"github.com/lysand-org/versia-go/internal/repository"
	"github.com/lysand-org/versia-go/internal/service"
	"net/url"

	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/config"
	"github.com/lysand-org/versia-go/ent/schema"
	"github.com/lysand-org/versia-go/internal/entity"
	"github.com/lysand-org/versia-go/internal/utils"
	"github.com/lysand-org/versia-go/pkg/webfinger"
)

var _ service.UserService = (*UserServiceImpl)(nil)

type UserServiceImpl struct {
	repositories repository.Manager

	federationService service.FederationService

	telemetry *unitel.Telemetry

	log logr.Logger
}

func NewUserServiceImpl(repositories repository.Manager, federationService service.FederationService, telemetry *unitel.Telemetry, log logr.Logger) *UserServiceImpl {
	return &UserServiceImpl{
		repositories:      repositories,
		federationService: federationService,
		telemetry:         telemetry,
		log:               log,
	}
}

func (i UserServiceImpl) WithRepositories(repositories repository.Manager) service.UserService {
	return NewUserServiceImpl(repositories, i.federationService, i.telemetry, i.log)
}

func (i UserServiceImpl) NewUser(ctx context.Context, username, password string) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/UserServiceImpl.NewUser")
	defer s.End()
	ctx = s.Context()

	if err := schema.ValidateUsername(username); err != nil {
		return nil, err
	}

	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		i.log.Error(err, "Failed to generate ed25519 key pair")

		return nil, err
	}

	user, err := i.repositories.Users().NewUser(ctx, username, password, priv, pub)
	if err != nil {
		i.log.Error(err, "Failed to create user", "username", username)

		return nil, err
	}

	i.log.V(2).Info("Create user", "id", user.ID, "uri", user.URI)

	return user, nil
}

func (i UserServiceImpl) GetLocalUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/UserServiceImpl.GetLocalUserByID")
	defer s.End()
	ctx = s.Context()

	return i.repositories.Users().GetLocalByID(ctx, id)
}

func (i UserServiceImpl) GetWebfingerForUser(ctx context.Context, userID string) (*webfinger.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/UserServiceImpl.GetWebfingerForUser")
	defer s.End()
	ctx = s.Context()

	u, err := i.repositories.Users().LookupLocalByIDOrUsername(ctx, userID)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, api_schema.ErrUserNotFound
	}

	wf := &webfinger.User{
		UserID: webfinger.UserID{
			ID: u.ID.String(),
			// FIXME: Move this away into a service or sth
			Domain: config.C.PublicAddress.Host,
		},
		URI: utils.UserAPIURL(u.ID).ToStd(),
	}

	if u.Edges.AvatarImage != nil {
		avatarURL, err := url.Parse(u.Edges.AvatarImage.URL)
		if err != nil {
			i.log.Error(err, "Failed to parse avatar URL")

			wf.Avatar = utils.DefaultAvatarURL(u.ID).ToStd()
			wf.AvatarMIMEType = "image/svg+xml"
		} else {
			wf.Avatar = avatarURL
			wf.AvatarMIMEType = u.Edges.AvatarImage.MimeType
		}
	} else {
		wf.Avatar = utils.DefaultAvatarURL(u.ID).ToStd()
		wf.AvatarMIMEType = "image/svg+xml"
	}

	return wf, nil
}

func (i UserServiceImpl) Search(ctx context.Context, req api_schema.SearchUserRequest) (u *entity.User, err error) {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/UserServiceImpl.Search").
		AddAttribute("username", req.Username)
	defer s.End()
	ctx = s.Context()

	domain := ""
	if req.Domain != nil {
		domain = *req.Domain
	}

	err = i.repositories.Atomic(ctx, func(ctx context.Context, tx repository.Manager) error {
		var err error
		if u, err = i.repositories.Users().Discover(ctx, domain, req.Username); err != nil {
			return err
		}

		return nil
	})
	return
}
