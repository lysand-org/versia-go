package repo_impls

import (
	"context"
	"crypto/ed25519"
	"errors"
	"github.com/lysand-org/versia-go/internal/repository"
	"github.com/lysand-org/versia-go/internal/service"
	"golang.org/x/crypto/bcrypt"

	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/ent"
	"github.com/lysand-org/versia-go/ent/predicate"
	"github.com/lysand-org/versia-go/ent/user"
	"github.com/lysand-org/versia-go/internal/entity"
	"github.com/lysand-org/versia-go/internal/utils"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

const bcryptCost = 12

var (
	ErrUsernameTaken = errors.New("username taken")

	_ repository.UserRepository = (*UserRepositoryImpl)(nil)
)

type UserRepositoryImpl struct {
	federationService service.FederationService

	db        *ent.Client
	log       logr.Logger
	telemetry *unitel.Telemetry
}

func NewUserRepositoryImpl(federationService service.FederationService, db *ent.Client, log logr.Logger, telemetry *unitel.Telemetry) repository.UserRepository {
	return &UserRepositoryImpl{
		federationService: federationService,
		db:                db,
		log:               log,
		telemetry:         telemetry,
	}
}

func (i *UserRepositoryImpl) NewUser(ctx context.Context, username, password string, priv ed25519.PrivateKey, pub ed25519.PublicKey) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repository/repo_impls.UserRepositoryImpl.NewUser")
	defer s.End()
	ctx = s.Context()

	pwHash, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return nil, err
	}

	uid := uuid.New()

	u, err := i.db.User.Create().
		SetID(uid).
		SetIsRemote(false).
		SetURI(utils.UserAPIURL(uid).String()).
		SetUsername(username).
		SetPasswordHash(pwHash).
		SetPublicKey(pub).
		SetPrivateKey(priv).
		SetInbox(utils.UserInboxAPIURL(uid).String()).
		SetOutbox(utils.UserOutboxAPIURL(uid).String()).
		SetFeatured(utils.UserFeaturedAPIURL(uid).String()).
		SetFollowers(utils.UserFollowersAPIURL(uid).String()).
		SetFollowing(utils.UserFollowingAPIURL(uid).String()).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, ErrUsernameTaken
		}

		return nil, err
	}

	return entity.NewUser(u)
}

func (i *UserRepositoryImpl) ImportLysandUserByURI(ctx context.Context, uri *lysand.URL) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repository/repo_impls.UserRepositoryImpl.ImportLysandUserByURI")
	defer s.End()
	ctx = s.Context()

	lUser, err := i.federationService.GetUser(ctx, uri)
	if err != nil {
		i.log.Error(err, "Failed to fetch remote user", "uri", uri)
		return nil, err
	}

	id, err := i.db.User.Create().
		SetID(uuid.New()).
		SetIsRemote(true).
		SetURI(lUser.URI.String()).
		SetUsername(lUser.Username).
		SetNillableDisplayName(lUser.DisplayName).
		SetBiography(lUser.Bio.String()).
		SetPublicKey(lUser.PublicKey.PublicKey.ToStd()).
		SetIndexable(lUser.Indexable).
		SetFields(lUser.Fields).
		SetExtensions(lUser.Extensions).
		SetInbox(lUser.Inbox.String()).
		SetOutbox(lUser.Outbox.String()).
		SetFeatured(lUser.Featured.String()).
		SetFollowers(lUser.Followers.String()).
		SetFollowing(lUser.Following.String()).
		OnConflict().
		UpdateNewValues().
		ID(ctx)
	if err != nil {
		i.log.Error(err, "Failed to import user into database", "uri", lUser.URI)
		return nil, err
	}

	u, err := i.db.User.Get(ctx, id)
	if err != nil {
		i.log.Error(err, "Failed to get imported user", "id", id, "uri", lUser.URI)
		return nil, err
	}

	i.log.V(2).Info("Imported user into database", "id", id, "uri", lUser.URI)

	return entity.NewUser(u)
}

func (i *UserRepositoryImpl) Resolve(ctx context.Context, uri *lysand.URL) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repository/repo_impls.UserRepositoryImpl.Resolve")
	defer s.End()
	ctx = s.Context()

	u, err := i.LookupByURI(ctx, uri)
	if err != nil {
		return nil, err
	}

	// check if the user is already imported
	if u == nil {
		i.log.V(2).Info("User not found in DB", "uri", uri)

		u, err := i.ImportLysandUserByURI(ctx, uri)
		if err != nil {
			i.log.Error(err, "Failed to import user", "uri", uri)
			return nil, err
		}

		return u, nil
	}

	i.log.V(2).Info("User found in DB", "uri", uri)

	return u, nil
}

func (i *UserRepositoryImpl) ResolveMultiple(ctx context.Context, uris []lysand.URL) ([]*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repository/repo_impls.UserRepositoryImpl.ResolveMultiple")
	defer s.End()
	ctx = s.Context()

	us, err := i.LookupByURIs(ctx, uris)
	if err != nil {
		return nil, err
	}

	// TODO: Refactor to use async imports using a work queue
outer:
	for _, uri := range uris {
		// check if the user is already imported
		for _, u := range us {
			if uri.String() == u.URI.String() {
				i.log.V(2).Info("User found in DB", "uri", uri)

				continue outer
			}
		}

		i.log.V(2).Info("User not found in DB", "uri", uri)

		importedUser, err := i.ImportLysandUserByURI(ctx, &uri)
		if err != nil {
			i.log.Error(err, "Failed to import user", "uri", uri)

			continue
		}

		i.log.V(2).Info("Imported user", "uri", uri)

		us = append(us, importedUser)
	}

	return us, nil
}

func (i *UserRepositoryImpl) GetByID(ctx context.Context, uid uuid.UUID) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repository/repo_impls.UserRepositoryImpl.GetByID")
	defer s.End()
	ctx = s.Context()

	u, err := i.db.User.Query().
		Where(user.IDEQ(uid)).
		WithAvatarImage().
		WithHeaderImage().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			i.log.Error(err, "Failed to query user", "id", uid)
			return nil, err
		}

		i.log.V(2).Info("User not found in DB", "id", uid)

		return nil, nil
	}

	i.log.V(2).Info("User found in DB", "id", uid)

	return entity.NewUser(u)
}

func (i *UserRepositoryImpl) GetLocalByID(ctx context.Context, uid uuid.UUID) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repository/repo_impls.UserRepositoryImpl.GetLocalByID")
	defer s.End()
	ctx = s.Context()

	u, err := i.db.User.Query().
		Where(user.And(user.ID(uid), user.IsRemote(false))).
		WithAvatarImage().
		WithHeaderImage().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			i.log.Error(err, "Failed to query local user", "id", uid)
			return nil, err
		}

		i.log.V(2).Info("Local user not found in DB", "id", uid)

		return nil, nil
	}

	i.log.V(2).Info("Local user found in DB", "id", uid)

	return entity.NewUser(u)
}

func (i *UserRepositoryImpl) LookupByURI(ctx context.Context, uri *lysand.URL) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repository/repo_impls.UserRepositoryImpl.LookupByURI")
	defer s.End()
	ctx = s.Context()

	// check if the user is already imported
	u, err := i.db.User.Query().
		Where(user.URI(uri.String())).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			i.log.Error(err, "Failed to query user", "uri", uri)
			return nil, err
		}

		i.log.V(2).Info("User not found in DB", "uri", uri)

		return nil, nil
	}

	i.log.V(2).Info("User found in DB", "uri", uri)

	return entity.NewUser(u)
}

func (i *UserRepositoryImpl) LookupByURIs(ctx context.Context, uris []lysand.URL) ([]*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repository/repo_impls.UserRepositoryImpl.LookupByURIs")
	defer s.End()
	ctx = s.Context()

	urisStrs := make([]string, 0, len(uris))
	for _, u := range uris {
		urisStrs = append(urisStrs, u.String())
	}

	us, err := i.db.User.Query().
		Where(user.URIIn(urisStrs...)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return utils.MapErrorSlice(us, entity.NewUser)
}

func (i *UserRepositoryImpl) LookupByIDOrUsername(ctx context.Context, idOrUsername string) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repository/repo_impls.UserRepositoryImpl.LookupByIDOrUsername")
	defer s.End()
	ctx = s.Context()

	var preds []predicate.User
	if u, err := uuid.Parse(idOrUsername); err == nil {
		preds = append(preds, user.IDEQ(u))
	} else {
		preds = append(preds, user.UsernameEQ(idOrUsername))
	}

	u, err := i.db.User.Query().
		Where(preds...).
		WithAvatarImage().
		WithHeaderImage().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			i.log.Error(err, "Failed to query user", "idOrUsername", idOrUsername)
			return nil, err
		}

		i.log.V(2).Info("User not found in DB", "idOrUsername", idOrUsername)

		return nil, nil
	}

	i.log.V(2).Info("User found in DB", "idOrUsername", idOrUsername, "id", u.ID)

	return entity.NewUser(u)
}
