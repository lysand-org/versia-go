package api_schema

import (
	"github.com/google/uuid"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

type User struct {
	ID       uuid.UUID `json:"id,string"`
	Username string    `json:"username"`
}

type LysandUser lysand.User

type FetchUserResponse = APIResponse[User]

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,username_regex,min=1,max=32"`
	Password string `json:"password" validate:"required,min=8,max=256"`
}

type SearchUserRequest struct {
	Username string  `query:"username" validate:"required,username_regex,min=1,max=32"`
	Domain   *string `query:"domain" validate:"domain_regex"`
}
