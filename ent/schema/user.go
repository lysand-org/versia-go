package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"errors"
	"github.com/lysand-org/versia-go/pkg/versia"
	"regexp"
)

var (
	ErrUsernameInvalid = errors.New("username must match ^[a-z0-9_-]+$")
	usernameRegex      = regexp.MustCompile("^[a-z0-9_-]+$")
)

type User struct{ ent.Schema }

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique().
			MaxLen(32).
			Validate(ValidateUsername),
		field.Bytes("passwordHash").
			Optional().
			Nillable(),

		field.String("displayName").
			MaxLen(256).
			Optional().
			Nillable(),
		field.String("biography").
			Optional().
			Nillable(),

		field.Bytes("publicKey"),
		field.String("publicKeyActor"),
		field.String("publicKeyAlgorithm"),
		field.Bytes("privateKey").Optional(),

		field.Bool("indexable").Default(true),
		field.Enum("privacyLevel").
			Values("public", "restricted", "private").
			Default("public"),

		field.JSON("fields", []versia.UserField{}).Default([]versia.UserField{}),

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

		edge.From("servers", InstanceMetadata.Type).Ref("users"),
		edge.From("moderatedServers", InstanceMetadata.Type).Ref("moderators"),
		edge.From("administeredServers", InstanceMetadata.Type).Ref("admins"),
	}
}

func (User) Mixin() []ent.Mixin { return []ent.Mixin{LysandEntityMixin{}} }

func ValidateUsername(username string) error {
	if !usernameRegex.MatchString(username) {
		return ErrUsernameInvalid
	}

	return nil
}
