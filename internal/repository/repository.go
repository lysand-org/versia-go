package repository

import (
	"context"
	"crypto/ed25519"
	"github.com/versia-pub/versia-go/pkg/versia"
	versiautils "github.com/versia-pub/versia-go/pkg/versia/utils"

	"github.com/google/uuid"
	"github.com/versia-pub/versia-go/internal/entity"
)

type UserRepository interface {
	NewUser(ctx context.Context, username, password string, privateKey ed25519.PrivateKey, publicKey ed25519.PublicKey) (*entity.User, error)
	ImportVersiaUserByURI(ctx context.Context, uri *versiautils.URL) (*entity.User, error)

	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetLocalByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetLocalByUsername(ctx context.Context, username string) (*entity.User, error)

	Discover(ctx context.Context, host, username string) (*entity.User, error)

	Resolve(ctx context.Context, uri *versiautils.URL) (*entity.User, error)
	ResolveMultiple(ctx context.Context, uris []versiautils.URL) ([]*entity.User, error)

	LookupByURI(ctx context.Context, uri *versiautils.URL) (*entity.User, error)
	LookupByURIs(ctx context.Context, uris []versiautils.URL) ([]*entity.User, error)
	LookupLocalByIDOrUsername(ctx context.Context, idOrUsername string) (*entity.User, error)
}

type FollowRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Follow, error)

	Follow(ctx context.Context, follower, followee *entity.User) (*entity.Follow, error)
	Unfollow(ctx context.Context, follower, followee *entity.User) error
	AcceptFollow(ctx context.Context, follower, followee *entity.User) error
	RejectFollow(ctx context.Context, follower, followee *entity.User) error
}

type NoteRepository interface {
	NewNote(ctx context.Context, author *entity.User, content string, mentions []*entity.User) (*entity.Note, error)
	ImportVersiaNote(ctx context.Context, lNote *versia.Note) (*entity.Note, error)

	GetByID(ctx context.Context, idOrUsername uuid.UUID) (*entity.Note, error)
}

type InstanceMetadataRepository interface {
	GetByHost(ctx context.Context, host string) (*entity.InstanceMetadata, error)

	Resolve(ctx context.Context, host string) (*entity.InstanceMetadata, error)
}

type Manager interface {
	Atomic(ctx context.Context, fn func(ctx context.Context, tx Manager) error) error
	Ping() error

	Users() UserRepository
	Notes() NoteRepository
	Follows() FollowRepository
	InstanceMetadata() InstanceMetadataRepository
}
