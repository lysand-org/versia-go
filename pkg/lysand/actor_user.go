package lysand

import (
	"encoding/json"
)

// User represents a user object in the Lysand protocol. For more information, see the [Spec].
//
// [Spec]: https://lysand.org/objects/user
type User struct {
	Entity

	// PublicKey is the public key of the user.
	// https://lysand.org/objects/user#public-key
	PublicKey UserPublicKey `json:"public_key"`

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
