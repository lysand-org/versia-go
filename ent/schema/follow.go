package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Follow struct{ ent.Schema }

func (Follow) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").
			Values("pending", "accepted").
			Default("pending"),
	}
}

func (Follow) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("follower", User.Type).Unique().Required(),
		edge.To("followee", User.Type).Unique().Required(),
	}
}

func (Follow) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("follower", "followee").Unique(),
	}
}

func (Follow) Mixin() []ent.Mixin {
	return []ent.Mixin{VersiaEntityMixin{}}
}
