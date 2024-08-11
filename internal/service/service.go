package service

import (
	"context"
	"github.com/lysand-org/versia-go/internal/repository"

	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/internal/api_schema"
	"github.com/lysand-org/versia-go/internal/entity"
	"github.com/lysand-org/versia-go/pkg/lysand"
	"github.com/lysand-org/versia-go/pkg/webfinger"
)

type UserService interface {
	WithRepositories(repositories repository.Manager) UserService

	NewUser(ctx context.Context, username, password string) (*entity.User, error)

	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)

	GetWebfingerForUser(ctx context.Context, userID string) (*webfinger.User, error)
}

type FederationService interface {
	SendToInbox(ctx context.Context, author *entity.User, target *entity.User, object any) ([]byte, error)
	GetUser(ctx context.Context, uri *lysand.URL) (*lysand.User, error)
}

type InboxService interface {
	Handle(ctx context.Context, obj any, userId uuid.UUID) error
}

type NoteService interface {
	CreateNote(ctx context.Context, req api_schema.CreateNoteRequest) (*entity.Note, error)
	GetNote(ctx context.Context, id uuid.UUID) (*entity.Note, error)

	ImportLysandNote(ctx context.Context, lNote *lysand.Note) (*entity.Note, error)
}

type FollowService interface {
	NewFollow(ctx context.Context, follower, followee *entity.User) (*entity.Follow, error)
	GetFollow(ctx context.Context, id uuid.UUID) (*entity.Follow, error)

	ImportLysandFollow(ctx context.Context, lFollow *lysand.Follow) (*entity.Follow, error)
}

type TaskService interface {
	ScheduleTask(ctx context.Context, type_ string, data any) error
}
