package entity

import (
	"github.com/lysand-org/versia-go/internal/helpers"
	"github.com/lysand-org/versia-go/pkg/versia"
	versiacrypto "github.com/lysand-org/versia-go/pkg/versia/crypto"
	versiautils "github.com/lysand-org/versia-go/pkg/versia/utils"
	"net/url"

	"github.com/lysand-org/versia-go/ent"
	"github.com/lysand-org/versia-go/internal/utils"
)

type User struct {
	*ent.User

	URI        *versiautils.URL
	PKActorURI *versiautils.URL
	PublicKey  *versiacrypto.SPKIPublicKey
	Inbox      *versiautils.URL
	Outbox     *versiautils.URL
	Featured   *versiautils.URL
	Followers  *versiautils.URL
	Following  *versiautils.URL

	DisplayName     string
	LysandAvatar    versiautils.ImageContentTypeMap
	LysandBiography versiautils.TextContentTypeMap
	Signer          versiacrypto.Signer
}

func NewUser(dbData *ent.User) (*User, error) {
	u := &User{
		User: dbData,
		PublicKey: &versiacrypto.SPKIPublicKey{
			Key:       nil,
			Algorithm: dbData.PublicKeyAlgorithm,
		},
		DisplayName: dbData.Username,

		LysandAvatar:    lysandAvatar(dbData),
		LysandBiography: lysandBiography(dbData),
	}

	if dbData.DisplayName != nil {
		u.DisplayName = *dbData.DisplayName
	}

	var err error
	if u.PublicKey.Key, err = versiacrypto.ToTypedKey(dbData.PublicKeyAlgorithm, dbData.PublicKey); err != nil {
		return nil, err
	}

	if u.URI, err = versiautils.ParseURL(dbData.URI); err != nil {
		return nil, err
	}
	if u.PKActorURI, err = versiautils.ParseURL(dbData.PublicKeyActor); err != nil {
		return nil, err
	}
	if u.Inbox, err = versiautils.ParseURL(dbData.Inbox); err != nil {
		return nil, err
	}
	if u.Outbox, err = versiautils.ParseURL(dbData.Outbox); err != nil {
		return nil, err
	}
	if u.Featured, err = versiautils.ParseURL(dbData.Featured); err != nil {
		return nil, err
	}
	if u.Followers, err = versiautils.ParseURL(dbData.Followers); err != nil {
		return nil, err
	}
	if u.Following, err = versiautils.ParseURL(dbData.Following); err != nil {
		return nil, err
	}

	u.Signer = versiacrypto.Signer{
		PrivateKey: dbData.PrivateKey,
		UserURL:    u.URI.ToStd(),
	}

	return u, nil
}

func (u User) ToLysand() *versia.User {
	return &versia.User{
		Entity: versia.Entity{
			ID:         u.ID,
			URI:        u.URI,
			CreatedAt:  versiautils.Time(u.CreatedAt),
			Extensions: u.Extensions,
		},
		DisplayName: helpers.StringPtr(u.DisplayName),
		Username:    u.Username,
		Avatar:      u.LysandAvatar,
		Header:      imageMap(u.Edges.HeaderImage),
		Indexable:   u.Indexable,
		PublicKey: versia.UserPublicKey{
			Actor:     u.PKActorURI,
			Algorithm: u.PublicKeyAlgorithm,
			Key:       u.PublicKey,
		},
		Bio:    u.LysandBiography,
		Fields: u.Fields,

		Inbox:     u.Inbox,
		Outbox:    u.Outbox,
		Featured:  u.Featured,
		Followers: u.Followers,
		Following: u.Following,
	}
}

func lysandAvatar(u *ent.User) versiautils.ImageContentTypeMap {
	if avatar := imageMap(u.Edges.AvatarImage); avatar != nil {
		return avatar
	}

	return versiautils.ImageContentTypeMap{
		"image/svg+xml": versiautils.ImageContent{
			Content: utils.DefaultAvatarURL(u.ID),
		},
	}
}

func lysandBiography(u *ent.User) versiautils.TextContentTypeMap {
	if u.Biography == nil {
		return nil
	}

	// TODO: Render HTML

	return versiautils.TextContentTypeMap{
		"text/html": versiautils.TextContent{
			Content: *u.Biography,
		},
	}
}

func imageMap(i *ent.Image) versiautils.ImageContentTypeMap {
	if i == nil {
		return nil
	}

	u, err := url.Parse(i.URL)
	if err != nil {
		return nil
	}

	return versiautils.ImageContentTypeMap{
		i.MimeType: {
			Content: (*versiautils.URL)(u),
		},
	}
}
