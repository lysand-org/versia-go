package webfinger

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lysand-org/versia-go/pkg/protoretry"
	"net/http"
	"net/url"
	"strings"
)

var (
	ErrInvalidSyntax = errors.New("must follow the format \"acct:<ID|Username>@<DOMAIN>\"")
	ErrUserNotFound  = errors.New("user could not be found")
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
	Type     string `json:"type"`
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

func Discover(c *protoretry.Client, ctx context.Context, baseURI, username string) (*User, error) {
	u := &User{UserID: UserID{ID: username, Domain: baseURI}}

	body, resp, err := c.GET(ctx, &url.URL{
		Scheme:   "https",
		Host:     u.UserID.Domain,
		Path:     "/.well-known/webfinger",
		RawQuery: url.Values{"resource": []string{"acct:" + u.String()}}.Encode(),
	})
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrUserNotFound
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var respBody Response
	if err := json.Unmarshal(body, &respBody); err != nil {
		return nil, err
	}

	if respBody.Error != nil {
		return nil, fmt.Errorf("webfinger error: %s", *respBody.Error)
	}

	for _, link := range respBody.Links {
		if link.Relation == "self" {
			if u.URI, err = url.Parse(link.Link); err != nil {
				return nil, err
			}
		} else if link.Relation == "avatar" {
			u.AvatarMIMEType = link.Type
			if u.Avatar, err = url.Parse(link.Link); err != nil {
				return nil, err
			}
		}
	}

	return u, nil
}
