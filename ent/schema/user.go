package schema

import (
	"crypto/ed25519"
	"errors"
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/lysand-org/versia-go/pkg/lysand"
)

var (
	ErrUsernameInvalid = errors.New("username must match ^[a-z0-9_-]+$")
	usernameRegex      = regexp.MustCompile("^[a-z0-9_-]+$")
)

type User struct{ ent.Schema }

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().MaxLen(32).Validate(ValidateUsername),
		field.Bytes("passwordHash").Optional().Nillable(),

		field.String("displayName").MaxLen(256).Optional().Nillable(),
		field.String("biography").Optional().Nillable(),

		field.Bytes("publicKey").GoType(ed25519.PublicKey([]byte{})),
		field.Bytes("privateKey").GoType(ed25519.PrivateKey([]byte{})).Optional(),

		field.Bool("indexable").Default(true),
		field.Enum("privacyLevel").Values("public", "restricted", "private").Default("public"),

		field.JSON("fields", []lysand.Field{}).Default([]lysand.Field{}),

		field.String("inbox").Validate(ValidateURI),

		// Collections
		field.String("featured").Validate(ValidateURI),
		field.String("followers").Validate(ValidateURI),
		field.String("following").Validate(ValidateURI),
		field.String("outbox").Validate(ValidateURI),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("avatarImage", Image.Type).Unique(),
		edge.To("headerImage", Image.Type).Unique(),

		edge.From("authoredNotes", Note.Type).Ref("author"),
		edge.From("mentionedNotes", Note.Type).Ref("mentions"),
	}
}

func (User) Mixin() []ent.Mixin { return []ent.Mixin{LysandEntityMixin{}} }

func ValidateUsername(username string) error {
	if !usernameRegex.MatchString(username) {
		return ErrUsernameInvalid
	}

	return nil
}
