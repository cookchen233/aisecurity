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
		field.String("title").Comment("标题").MaxLen(255),
		field.Int("parent_id").Comment("上级部门id").NonNegative(),
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("department_creator").Field("created_by").Immutable().Unique().Required(),
		edge.From("parent", Department.Type).Ref("children").Unique().Field("parent_id").Required(),

		edge.To("employee_department", Employee.Type),
		edge.To("children", Department.Type),
	}
}
