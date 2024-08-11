package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Image struct{ ent.Schema }

func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.String("url"),
		field.String("mimeType"),
	}
}

func (Image) Edges() []ent.Edge {
	return nil
}
