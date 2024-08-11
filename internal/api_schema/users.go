package api_schema

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id,string"`
	Username string    `json:"username"`
}

type FetchUserResponse = APIResponse[User]

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,username_regex,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=256"`
}
