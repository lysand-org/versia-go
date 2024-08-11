package lysand

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// User represents a user object in the Lysand protocol. For more information, see the [Spec].
//
// [Spec]: https://lysand.org/objects/user
type User struct {
	Entity

	// PublicKey is the public key of the user.
	// https://lysand.org/objects/user#public-key
	PublicKey PublicKey `json:"public_key"`

	// DisplayName is the display name of the user.
	// https://lysand.org/objects/user#display-name
	DisplayName *string `json:"display_name,omitempty"`

	// Username is the username of the user. Must be unique on the instance and match the following regex: ^[a-z0-9_-]+$
	// https://lysand.org/objects/user#username
	Username string `json:"username"`

	// Indexable is a boolean that indicates whether the user is indexable by search engines.
	// https://lysand.org/objects/user#indexable
	Indexable bool `json:"indexable"`

	// ManuallyApprovesFollowers is a boolean that indicates whether the user manually approves followers.
	// https://lysand.org/objects/user#manually-approves-followers
	ManuallyApprovesFollowers bool `json:"manually_approves_followers"`

	// Avatar is the avatar of the user in different image content types.
	// https://lysand.org/objects/user#avatar
	Avatar ImageContentTypeMap `json:"avatar,omitempty"`

	// Header is the header image of the user in different image content types.
	// https://lysand.org/objects/user#header
	Header ImageContentTypeMap `json:"header,omitempty"`

	// Bio is the biography of the user in different text content types.
	// https://lysand.org/objects/user#bio
	Bio TextContentTypeMap `json:"bio"`

	// Fields is a list of fields that the user has filled out.
	// https://lysand.org/objects/user#fields
	Fields []Field `json:"fields,omitempty"`

	// Featured is the featured posts of the user.
	// https://lysand.org/objects/user#featured
	Featured *URL `json:"featured"`

	// Followers is the followers of the user.
	// https://lysand.org/objects/user#followers
	Followers *URL `json:"followers"`

	// Following is the users that the user is following.
	// https://lysand.org/objects/user#following
	Following *URL `json:"following"`

	// Likes is the likes of the user.
	// https://lysand.org/objects/user#likes
	Likes *URL `json:"likes"`

	// Dislikes is the dislikes of the user.
	// https://lysand.org/objects/user#dislikes
	Dislikes *URL `json:"dislikes"`

	// Inbox is the inbox of the user.
	// https://lysand.org/objects/user#posts
	Inbox *URL `json:"inbox"`

	// Outbox is the outbox of the user.
	// https://lysand.org/objects/user#outbox
	Outbox *URL `json:"outbox"`
}

func (u User) MarshalJSON() ([]byte, error) {
	type user User
	u2 := user(u)
	u2.Type = "User"
	return json.Marshal(u2)
}

type Field struct {
	Key   TextContentTypeMap `json:"key"`
	Value TextContentTypeMap `json:"value"`
}

func (c *FederationClient) GetUser(ctx context.Context, uri *url.URL) (*User, error) {
	resp, body, err := c.rawGET(ctx, uri)
	if err != nil {
		return nil, err
	}

	user := &User{}
	if err := json.Unmarshal(body, user); err != nil {
		return nil, err
	}

	date, sigHeader, err := ExtractFederationHeaders(resp.Header)
	if err != nil {
		return nil, err
	}

	v := Verifier{ed25519.PublicKey(user.PublicKey.PublicKey)}
	if !v.Verify("GET", date, uri.Host, uri.Path, body, sigHeader) {
		c.log.V(2).Info("signature verification failed", "user", user.URI.String())
		return nil, fmt.Errorf("signature verification failed")
	}
	c.log.V(2).Info("signature verification succeeded", "user", user.URI.String())

	return user, nil
}

func (c *FederationClient) SendToInbox(ctx context.Context, signer Signer, user *User, object any) ([]byte, error) {
	uri := user.Inbox.ToStd()

	body, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}

	date := time.Now()

	sigData := NewSignatureData("POST", date, uri.Host, uri.Path, hashSHA256(body))
	sig := signer.Sign(*sigData)

	req, err := http.NewRequestWithContext(ctx, "POST", uri.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Date", TimeFromStd(date).String())
	req.Header.Set("Signature", sig.String())

	_, respBody, err := c.doReq(req)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
