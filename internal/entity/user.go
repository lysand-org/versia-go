package entity

import (
	"github.com/lysand-org/versia-go/internal/helpers"
	versiacrypto "github.com/lysand-org/versia-go/pkg/lysand/crypto"
	"net/url"

	"github.com/lysand-org/versia-go/ent"
	"github.com/lysand-org/versia-go/internal/utils"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

type User struct {
	*ent.User

	URI        *lysand.URL
	PKActorURI *lysand.URL
	PublicKey  *lysand.SPKIPublicKey
	Inbox      *lysand.URL
	Outbox     *lysand.URL
	Featured   *lysand.URL
	Followers  *lysand.URL
	Following  *lysand.URL

	DisplayName     string
	LysandAvatar    lysand.ImageContentTypeMap
	LysandBiography lysand.TextContentTypeMap
	Signer          lysand.Signer
}

func NewUser(dbData *ent.User) (*User, error) {
	u := &User{
		User: dbData,
		PublicKey: &lysand.SPKIPublicKey{
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

	if u.URI, err = lysand.ParseURL(dbData.URI); err != nil {
		return nil, err
	}
	if u.PKActorURI, err = lysand.ParseURL(dbData.PublicKeyActor); err != nil {
		return nil, err
	}
	if u.Inbox, err = lysand.ParseURL(dbData.Inbox); err != nil {
		return nil, err
	}
	if u.Outbox, err = lysand.ParseURL(dbData.Outbox); err != nil {
		return nil, err
	}
	if u.Featured, err = lysand.ParseURL(dbData.Featured); err != nil {
		return nil, err
	}
	if u.Followers, err = lysand.ParseURL(dbData.Followers); err != nil {
		return nil, err
	}
	if u.Following, err = lysand.ParseURL(dbData.Following); err != nil {
		return nil, err
	}

	u.Signer = lysand.Signer{
		PrivateKey: dbData.PrivateKey,
		UserURL:    u.URI.ToStd(),
	}

	return u, nil
}

func (u User) ToLysand() *lysand.User {
	return &lysand.User{
		Entity: lysand.Entity{
			ID:         u.ID,
			URI:        u.URI,
			CreatedAt:  lysand.TimeFromStd(u.CreatedAt),
			Extensions: u.Extensions,
		},
		DisplayName: helpers.StringPtr(u.DisplayName),
		Username:    u.Username,
		Avatar:      u.LysandAvatar,
		Header:      imageMap(u.Edges.HeaderImage),
		Indexable:   u.Indexable,
		PublicKey: lysand.UserPublicKey{
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

func lysandAvatar(u *ent.User) lysand.ImageContentTypeMap {
	if avatar := imageMap(u.Edges.AvatarImage); avatar != nil {
		return avatar
	}

	return lysand.ImageContentTypeMap{
		"image/svg+xml": lysand.ImageContent{
			Content: utils.DefaultAvatarURL(u.ID),
		},
	}
}

func lysandBiography(u *ent.User) lysand.TextContentTypeMap {
	if u.Biography == nil {
		return nil
	}

	// TODO: Render HTML

	return lysand.TextContentTypeMap{
		"text/html": lysand.TextContent{
			Content: *u.Biography,
		},
	}
}

func imageMap(i *ent.Image) lysand.ImageContentTypeMap {
	if i == nil {
		return nil
	}

	u, err := url.Parse(i.URL)
	if err != nil {
		return nil
	}

	return lysand.ImageContentTypeMap{
		i.MimeType: {
			Content: (*lysand.URL)(u),
		},
	}
}
