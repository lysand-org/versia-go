package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type InstanceMetadata struct{ ent.Schema }

func (InstanceMetadata) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description").
			Optional().
			Nillable(),
		field.String("host").
			NotEmpty().
			Unique(),

		field.Bytes("publicKey"),
		field.String("publicKeyAlgorithm"),
		field.Bytes("privateKey").Optional(),

		field.String("softwareName").NotEmpty(),
		field.String("softwareVersion").NotEmpty(),

		field.String("sharedInboxURI").Validate(ValidateURI),
		field.String("moderatorsURI").
			Validate(ValidateURI).
			Optional().
			Nillable(),
		field.String("adminsURI").
			Validate(ValidateURI).
			Optional().
			Nillable(),

		field.String("logoEndpoint").
			Validate(ValidateURI).
			Optional().
			Nillable(),
		field.String("logoMimeType").
			Validate(ValidateURI).
			Optional().
			Nillable(),
		field.String("bannerEndpoint").
			Validate(ValidateURI).
			Optional().
			Nillable(),
		field.String("bannerMimeType").
			Validate(ValidateURI).
			Optional().
			Nillable(),

		field.Strings("supportedVersions").Default([]string{}),
		field.Strings("supportedExtensions").Default([]string{}),
	}
}

func (InstanceMetadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),

		edge.To("moderators", User.Type),
		edge.To("admins", User.Type),
	}
}

func (InstanceMetadata) Mixin() []ent.Mixin {
	return []ent.Mixin{VersiaEntityMixin{}}
}
