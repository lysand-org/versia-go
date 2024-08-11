package webfinger

import (
	"errors"
	"net/url"
	"strings"
)

var (
	ErrInvalidSyntax = errors.New("must follow the format \"acct:<ID|Username>@<DOMAIN>\"")
)

func ParseResource(res string) (*UserID, error) {
	if !strings.HasPrefix(res, "acct:") {
		return nil, ErrInvalidSyntax
	}

	if !strings.Contains(res, "@") {
		return nil, ErrInvalidSyntax
	}

	spl := strings.Split(res, "@")
	if len(spl) != 2 {
		return nil, ErrInvalidSyntax
	}

	userID := strings.TrimPrefix(spl[0], "acct:")
	domain := spl[1]

	return &UserID{userID, domain}, nil
}

type UserID struct {
	ID     string
	Domain string
}

func (u UserID) String() string {
	return u.ID + "@" + u.Domain
}

type Response struct {
	Subject string `json:"subject,omitempty"`
	Links   []Link `json:"links,omitempty"`

	Error *string `json:"error,omitempty"`
}

type Link struct {
	Relation string `json:"rel"`
	Type     any    `json:"type"`
	Link     string `json:"href"`
}

type User struct {
	UserID

	URI *url.URL

	Avatar         *url.URL
	AvatarMIMEType string
}

func (u User) WebFingerResource() Response {
	return Response{
		Subject: "acct:" + u.String(),
		Links: []Link{
			{"self", "application/json", u.URI.String()},
			{"avatar", u.AvatarMIMEType, u.Avatar.String()},
		},
	}
}
