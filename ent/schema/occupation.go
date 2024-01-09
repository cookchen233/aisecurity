package schema

import (
	"aisecurity/ent/mixin"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Occupation holds the schema definition for the Occupation entity.
type Occupation struct {
	ent.Schema
}

func (Occupation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the Occupation.
func (Occupation) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").NotEmpty().MaxLen(255).StructTag(`validate:"required"`),
		field.String("description").Comment("描述").Optional(),
	}
}

// Edges of the Occupation.
func (Occupation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("occupation_creator").Field("created_by").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("occupation_updater").Field("updated_by").Required().Unique(),

		edge.To("employees", Employee.Type),
	}
}
