package entity

import (
	"github.com/lysand-org/versia-go/internal/helpers"
	"net/url"

	"github.com/lysand-org/versia-go/ent"
	"github.com/lysand-org/versia-go/internal/utils"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

type User struct {
	*ent.User

	URI       *lysand.URL
	Inbox     *lysand.URL
	Outbox    *lysand.URL
	Featured  *lysand.URL
	Followers *lysand.URL
	Following *lysand.URL

	DisplayName     string
	LysandAvatar    lysand.ImageContentTypeMap
	LysandBiography lysand.TextContentTypeMap
	Signer          lysand.Signer
}

func NewUser(dbUser *ent.User) (*User, error) {
	u := &User{User: dbUser}

	u.DisplayName = u.Username
	if dbUser.DisplayName != nil {
		u.DisplayName = *dbUser.DisplayName
	}

	var err error
	if u.URI, err = lysand.ParseURL(dbUser.URI); err != nil {
		return nil, err
	}
	if u.Inbox, err = lysand.ParseURL(dbUser.Inbox); err != nil {
		return nil, err
	}
	if u.Outbox, err = lysand.ParseURL(dbUser.Outbox); err != nil {
		return nil, err
	}
	if u.Featured, err = lysand.ParseURL(dbUser.Featured); err != nil {
		return nil, err
	}
	if u.Followers, err = lysand.ParseURL(dbUser.Followers); err != nil {
		return nil, err
	}
	if u.Following, err = lysand.ParseURL(dbUser.Following); err != nil {
		return nil, err
	}

	u.LysandAvatar = lysandAvatar(dbUser)
	u.LysandBiography = lysandBiography(dbUser)
	u.Signer = lysand.Signer{
		PrivateKey: dbUser.PrivateKey,
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
		PublicKey: lysand.PublicKey{
			Actor:     utils.UserAPIURL(u.ID),
			PublicKey: lysand.SPKIPublicKey(u.PublicKey),
		},
		Bio:    u.LysandBiography,
		Fields: u.Fields,

		Inbox:     u.Inbox,
		Outbox:    u.Outbox,
		Featured:  u.Featured,
		Followers: u.Followers,
		Following: u.Following,

		// TODO: Remove these, they got deprecated and moved into an extension
		Likes:    utils.UserLikesAPIURL(u.ID),
		Dislikes: utils.UserDislikesAPIURL(u.ID),
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
