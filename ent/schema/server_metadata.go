package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type ServerMetadata struct{ ent.Schema }

func (ServerMetadata) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("description").
			Optional().
			Nillable(),

		field.String("version").
			NotEmpty(),

		field.Strings("supportedExtensions").
			Default([]string{}),
	}
}

func (ServerMetadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("follower", User.Type).Unique().Required(),
		edge.To("followee", User.Type).Unique().Required(),
	}
}

func (ServerMetadata) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("follower", "followee").Unique(),
	}
}

func (ServerMetadata) Mixin() []ent.Mixin {
	return []ent.Mixin{LysandEntityMixin{}}
}
