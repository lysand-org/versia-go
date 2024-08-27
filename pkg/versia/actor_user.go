package versia

import (
	"encoding/json"
	versiacrypto "github.com/lysand-org/versia-go/pkg/versia/crypto"
	versiautils "github.com/lysand-org/versia-go/pkg/versia/utils"
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
	Avatar versiautils.ImageContentTypeMap `json:"avatar,omitempty"`

	// Header is the header image of the user in different image content types.
	// https://lysand.org/objects/user#header
	Header versiautils.ImageContentTypeMap `json:"header,omitempty"`

	// Bio is the biography of the user in different text content types.
	// https://lysand.org/objects/user#bio
	Bio versiautils.TextContentTypeMap `json:"bio"`

	// Fields is a list of fields that the user has filled out.
	// https://lysand.org/objects/user#fields
	Fields []UserField `json:"fields,omitempty"`

	// Inbox is the inbox of the user.
	// https://lysand.org/objects/user#posts
	Inbox *versiautils.URL `json:"inbox"`

	// Collections is a map of all collections for a user
	Collections UserCollections `json:"collections"`
}

func (u User) MarshalJSON() ([]byte, error) {
	type user User
	u2 := user(u)
	u2.Type = "User"
	return json.Marshal(u2)
}

type UserField struct {
	Key   versiautils.TextContentTypeMap `json:"key"`
	Value versiautils.TextContentTypeMap `json:"value"`
}

type UserCollectionType string

const (
	// UserCollectionFeatured is a URL to a collection of the user's featured posts.
	UserCollectionFeatured UserCollectionType = "featured"
	// UserCollectionFollowers is a URL to a collection of the user's followers.
	UserCollectionFollowers UserCollectionType = "followers"
	// UserCollectionFollowing is a URL to a collection of the users that the user is following.
	UserCollectionFollowing UserCollectionType = "following"
	// UserCollectionOutbox is a URL to a collection of the user's posts.
	UserCollectionOutbox UserCollectionType = "outbox"
)

// UserCollections is a map of all collections for a user
type UserCollections map[UserCollectionType]*versiautils.URL

// UserPublicKey represents a public key for a user. For more information, see the [Spec].
//
// [Spec]: https://lysand.org/security/keys#public-key-cryptography
type UserPublicKey struct {
	Actor *versiautils.URL `json:"actor"`

	// Algorithm can only be `ed25519` for now
	Algorithm string `json:"algorithm"`

	Key    *versiacrypto.SPKIPublicKey `json:"-"`
	RawKey json.RawMessage             `json:"key"`
}

func (k *UserPublicKey) UnmarshalJSON(raw []byte) error {
	type t UserPublicKey
	k2 := (*t)(k)

	if err := json.Unmarshal(raw, k2); err != nil {
		return err
	}

	var err error
	if k2.Key, err = versiacrypto.UnmarshalSPKIPubKey(k2.Algorithm, k2.RawKey); err != nil {
		return err
	}

	*k = UserPublicKey(*k2)

	return nil
}

func (k UserPublicKey) MarshalJSON() ([]byte, error) {
	type t UserPublicKey
	k2 := t(k)

	var err error
	if k2.RawKey, err = k2.Key.MarshalJSON(); err != nil {
		return nil, err
	}

	return json.Marshal(k2)
}
