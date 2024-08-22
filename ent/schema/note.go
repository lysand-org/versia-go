package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Note struct{ ent.Schema }

func (Note) Fields() []ent.Field {
	return []ent.Field{
		field.String("subject").MaxLen(384).Optional().Nillable(),
		field.String("content"),

		field.Bool("isSensitive").Default(false),
		field.Enum("visibility").Values("public", "unlisted", "followers", "direct").Default("public"),
	}
}

func (Note) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", User.Type).Unique().Required(),
		edge.To("mentions", User.Type),

		edge.To("attachments", Attachment.Type),
	}
}

func (Note) Mixin() []ent.Mixin {
	return []ent.Mixin{VersiaEntityMixin{}}
}
