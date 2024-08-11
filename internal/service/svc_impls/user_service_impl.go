package svc_impls

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
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
	s := i.telemetry.StartSpan(ctx, "function", "service/svc_impls.UserServiceImpl.NewUser")
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

func (i UserServiceImpl) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "service/svc_impls.UserServiceImpl.GetUserByID")
	defer s.End()
	ctx = s.Context()

	return i.repositories.Users().LookupByIDOrUsername(ctx, id.String())
}

func (i UserServiceImpl) GetWebfingerForUser(ctx context.Context, userID string) (*webfinger.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "service/svc_impls.UserServiceImpl.GetWebfingerForUser")
	defer s.End()
	ctx = s.Context()

	u, err := i.repositories.Users().LookupByIDOrUsername(ctx, userID)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, fmt.Errorf("user not found")
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
