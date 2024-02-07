package schema

import (
	"aisecurity/ent/mixin"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

func (Department) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").NotEmpty().MaxLen(255).StructTag(`validate:"required"`),
		field.Int("parent_id").Comment("上级部门id").Optional().NonNegative(),
		field.String("notes").Comment("备注").Optional(),
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("department_creator").Field("creator_id").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("department_updater").Field("updater_id").Required().Unique(),

		edge.From("parent", Department.Type).Ref("children").Field("parent_id").Unique(),
		edge.From("permissions", Permission.Type).Ref("department"),

		edge.To("employees", Employee.Type),
		edge.To("children", Department.Type),
	}
}
