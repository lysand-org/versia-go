package api_schema

import (
	"github.com/google/uuid"
	"github.com/versia-pub/versia-go/pkg/versia"
)

type User struct {
	ID       uuid.UUID `json:"id,string"`
	Username string    `json:"username"`
}

type VersiaUser versia.User

type FetchUserResponse = APIResponse[User]

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,username_regex,min=1,max=32"`
	Password string `json:"password" validate:"required,min=8,max=256"`
}

type SearchUserRequest struct {
	Username string  `query:"username" validate:"required,username_regex,min=1,max=32"`
	Domain   *string `query:"domain" validate:"domain_regex"`
}

//var ErrInvalidUserMention = errors.New("invalid user mention")
//func (r *SearchUserRequest) UnmarshalJSON(raw []byte) error {
//	var s string
//	if err := json.Unmarshal(raw, &s); err != nil {
//		return err
//	}
//
//	s = strings.TrimPrefix(s, "@")
//	spl := strings.Split(s, "@")
//
//	if len(spl) > 2 {
//		return ErrInvalidUserMention
//	}
//
//	username := spl[0]
//
//	var domain *string
//	if len(spl) > 1 {
//		domain = &spl[1]
//	}
//
//	*r = SearchUserRequest{
//		Username: username,
//		Domain:   domain,
//	}
//
//	return nil
//}
