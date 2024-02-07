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
		field.Int("admin_id").Comment("管理员id").NonNegative().Immutable(),
		field.Int("department_id").Comment("部门id").NonNegative().Optional(),
		field.Int("occupation_id").Comment("岗位ID").NonNegative().Optional(),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("employee_creator").Field("creator_id").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("employee_updater").Field("updater_id").Required().Unique(),
		edge.From("admin", Admin.Type).Ref("employee").Field("admin_id").Immutable().Unique().Required(),
		edge.From("occupation", Occupation.Type).Ref("employee").Field("occupation_id").Unique(),
		edge.From("department", Department.Type).Ref("employees").Field("department_id").Unique(),
	}
}
