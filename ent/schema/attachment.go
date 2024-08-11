package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Attachment struct{ ent.Schema }

func (Attachment) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").MaxLen(384),
		field.Bytes("sha256"),
		field.Int("size"),

		field.String("blurhash").Optional().Nillable(),
		field.Int("height").Optional().Nillable(),
		field.Int("width").Optional().Nillable(),
		field.Int("fps").Optional().Nillable(),

		field.String("mimeType"),
	}
}

func (Attachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", User.Type).Unique().Required(),
	}
}

func (Attachment) Mixin() []ent.Mixin {
	return []ent.Mixin{LysandEntityMixin{}}
}
