package repository

import (
	"context"
	"crypto/ed25519"

	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/internal/entity"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

type UserRepository interface {
	NewUser(ctx context.Context, username, password string, privateKey ed25519.PrivateKey, publicKey ed25519.PublicKey) (*entity.User, error)
	ImportLysandUserByURI(ctx context.Context, uri *lysand.URL) (*entity.User, error)

	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetLocalByID(ctx context.Context, id uuid.UUID) (*entity.User, error)

	Resolve(ctx context.Context, uri *lysand.URL) (*entity.User, error)
	ResolveMultiple(ctx context.Context, uris []lysand.URL) ([]*entity.User, error)

	LookupByURI(ctx context.Context, uri *lysand.URL) (*entity.User, error)
	LookupByURIs(ctx context.Context, uris []lysand.URL) ([]*entity.User, error)
	LookupByIDOrUsername(ctx context.Context, idOrUsername string) (*entity.User, error)
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
	ImportLysandNote(ctx context.Context, lNote *lysand.Note) (*entity.Note, error)

	GetByID(ctx context.Context, idOrUsername uuid.UUID) (*entity.Note, error)
}

type Manager interface {
	Atomic(ctx context.Context, fn func(ctx context.Context, tx Manager) error) error

	Users() UserRepository
	Notes() NoteRepository
	Follows() FollowRepository
}
