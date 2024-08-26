package repo_impls

import (
	"context"
	"crypto/ed25519"
	"errors"
	"github.com/lysand-org/versia-go/config"
	"github.com/lysand-org/versia-go/internal/repository"
	"github.com/lysand-org/versia-go/internal/service"
	versiautils "github.com/lysand-org/versia-go/pkg/versia/utils"
	"golang.org/x/crypto/bcrypt"

	"git.devminer.xyz/devminer/unitel"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/ent"
	"github.com/lysand-org/versia-go/ent/predicate"
	"github.com/lysand-org/versia-go/ent/user"
	"github.com/lysand-org/versia-go/internal/entity"
	"github.com/lysand-org/versia-go/internal/utils"
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
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/UserRepositoryImpl.NewUser")
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
		SetPrivateKey(priv).
		SetPublicKey(pub).
		SetPublicKeyAlgorithm("ed25519").
		SetPublicKeyActor(utils.UserAPIURL(uid).String()).
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

func (i *UserRepositoryImpl) ImportLysandUserByURI(ctx context.Context, uri *versiautils.URL) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/UserRepositoryImpl.ImportLysandUserByURI")
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
		SetPublicKey(lUser.PublicKey.RawKey).
		SetPublicKeyAlgorithm(lUser.PublicKey.Algorithm).
		SetPublicKeyActor(lUser.PublicKey.Actor.String()).
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

func (i *UserRepositoryImpl) Discover(ctx context.Context, domain, username string) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "svc_impls/UserServiceImpl.Search").
		AddAttribute("username", username).
		AddAttribute("domain", domain)
	defer s.End()
	ctx = s.Context()

	l := i.log.WithValues("domain", domain, "username", username)

	// FIXME: This *could* go wrong
	if domain != config.C.Host {
		l.V(2).Info("Discovering instance")

		im, err := i.federationService.DiscoverInstance(ctx, domain)
		if err != nil {
			l.Error(err, "Failed to discover instance")
			return nil, err
		}

		l = l.WithValues("host", im.Host)

		l.V(2).Info("Discovering user")

		wf, err := i.federationService.DiscoverUser(ctx, im.Host, username)
		if err != nil {
			l.Error(err, "Failed to discover user")
			return nil, err
		}

		l.V(2).Info("Found remote user", "userURI", wf.URI)

		u, err := i.Resolve(ctx, versiautils.URLFromStd(wf.URI))
		if err != nil {
			l.Error(err, "Failed to resolve user")
			return nil, err
		}

		return u, nil
	}

	l.V(2).Info("Finding local user")

	u, err := i.GetLocalByUsername(ctx, username)
	if err != nil {
		l.Error(err, "Failed to find local user", "username", username)
		return nil, err
	}

	l.V(2).Info("Found local user", "userURI", u.URI)

	return u, nil
}

func (i *UserRepositoryImpl) Resolve(ctx context.Context, uri *versiautils.URL) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/UserRepositoryImpl.Resolve")
	defer s.End()
	ctx = s.Context()

	l := i.log.WithValues("uri", uri)

	u, err := i.LookupByURI(ctx, uri)
	if err != nil {
		return nil, err
	}

	// check if the user is already imported
	if u == nil {
		l.V(2).Info("User not found in DB")

		u, err := i.ImportLysandUserByURI(ctx, uri)
		if err != nil {
			l.Error(err, "Failed to import user")
			return nil, err
		}

		return u, nil
	}

	l.V(2).Info("User found in DB")

	return u, nil
}

func (i *UserRepositoryImpl) ResolveMultiple(ctx context.Context, uris []versiautils.URL) ([]*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/UserRepositoryImpl.ResolveMultiple")
	defer s.End()
	ctx = s.Context()

	us, err := i.LookupByURIs(ctx, uris)
	if err != nil {
		return nil, err
	}

	// TODO: Refactor to use async imports using a work queue
outer:
	for _, uri := range uris {
		l := i.log.WithValues("uri", uri)

		// check if the user is already imported
		for _, u := range us {
			if uri.String() == u.URI.String() {
				l.V(2).Info("User found in DB")

				continue outer
			}
		}

		l.V(2).Info("User not found in DB")

		importedUser, err := i.ImportLysandUserByURI(ctx, &uri)
		if err != nil {
			l.Error(err, "Failed to import user")

			continue
		}

		l.V(2).Info("Imported user")

		us = append(us, importedUser)
	}

	return us, nil
}

func (i *UserRepositoryImpl) GetByID(ctx context.Context, uid uuid.UUID) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/UserRepositoryImpl.GetByID")
	defer s.End()
	ctx = s.Context()

	l := i.log.WithValues("id", uid)

	u, err := i.db.User.Query().
		Where(user.IDEQ(uid)).
		WithAvatarImage().
		WithHeaderImage().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			l.Error(err, "Failed to query user")
			return nil, err
		}

		l.V(2).Info("User not found in DB")

		return nil, nil
	}

	l.V(2).Info("User found in DB")

	return entity.NewUser(u)
}

func (i *UserRepositoryImpl) GetLocalByID(ctx context.Context, uid uuid.UUID) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/UserRepositoryImpl.GetLocalByID")
	defer s.End()
	ctx = s.Context()

	l := i.log.WithValues("id", uid)

	u, err := i.db.User.Query().
		Where(user.And(user.ID(uid), user.IsRemote(false))).
		WithAvatarImage().
		WithHeaderImage().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			l.Error(err, "Failed to query local user")
			return nil, err
		}

		l.V(2).Info("Local user not found in DB")

		return nil, nil
	}

	l.V(2).Info("Local user found in DB", "uri", u.URI)

	return entity.NewUser(u)
}

func (i *UserRepositoryImpl) GetLocalByUsername(ctx context.Context, username string) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/UserRepositoryImpl.GetLocalByUsername")
	defer s.End()
	ctx = s.Context()

	l := i.log.WithValues("username", username)

	u, err := i.db.User.Query().
		Where(user.And(user.Username(username), user.IsRemote(false))).
		WithAvatarImage().
		WithHeaderImage().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			l.Error(err, "Failed to query local user")
			return nil, err
		}

		l.V(2).Info("Local user not found in DB")

		return nil, nil
	}

	l.V(2).Info("Local user found in DB", "uri", u.URI)

	return entity.NewUser(u)
}

func (i *UserRepositoryImpl) LookupByURI(ctx context.Context, uri *versiautils.URL) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/UserRepositoryImpl.LookupByURI")
	defer s.End()
	ctx = s.Context()

	l := i.log.WithValues("uri", uri)

	// check if the user is already imported
	u, err := i.db.User.Query().
		Where(user.URI(uri.String())).
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			l.Error(err, "Failed to query user")
			return nil, err
		}

		l.V(2).Info("User not found in DB")

		return nil, nil
	}

	l.V(2).Info("User found in DB")

	return entity.NewUser(u)
}

func (i *UserRepositoryImpl) LookupByURIs(ctx context.Context, uris []versiautils.URL) ([]*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/UserRepositoryImpl.LookupByURIs")
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

func (i *UserRepositoryImpl) LookupLocalByIDOrUsername(ctx context.Context, idOrUsername string) (*entity.User, error) {
	s := i.telemetry.StartSpan(ctx, "function", "repo_impls/UserRepositoryImpl.LookupLocalByIDOrUsername")
	defer s.End()
	ctx = s.Context()

	preds := []predicate.User{user.IsRemote(false)}
	if u, err := uuid.Parse(idOrUsername); err == nil {
		preds = append(preds, user.IDEQ(u))
	} else {
		preds = append(preds, user.UsernameEQ(idOrUsername))
	}

	l := i.log.WithValues("idOrUsername", idOrUsername)

	u, err := i.db.User.Query().
		Where(preds...).
		WithAvatarImage().
		WithHeaderImage().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			l.Error(err, "Failed to query user")
			return nil, err
		}

		l.V(2).Info("User not found in DB")

		return nil, nil
	}

	l.V(2).Info("User found in DB", "id", u.ID)

	return entity.NewUser(u)
}
