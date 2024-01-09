package schema

import (
	"aisecurity/ent/mixin"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

func (Employee) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.Int("admin_id").Comment("管理员id").Positive().Immutable(),
		field.Int("department_id").Comment("部门id").Positive(),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("employee_creator").Field("created_by").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("employee_updater").Field("updated_by").Required().Unique(),

		edge.From("admin", Admin.Type).Ref("employee").Field("admin_id").Immutable().Unique().Required(),
		edge.From("department", Department.Type).Field("department_id").Ref("employees").Unique().Required(),
		edge.From("occupations", Occupation.Type).Ref("employees"),

		edge.To("risk_reporter", Risk.Type),
		edge.To("risk_maintainer", Risk.Type),
	}
}
