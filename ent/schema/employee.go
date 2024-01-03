package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Comment("创建时间").Default(time.Now).Immutable(),
		field.Int("created_by").Comment("创建者").Positive(),
		field.Int("admin_id").Positive().Comment("管理员id"),
		field.Int("department_id").Positive().Comment("部门id"),
		field.Time("deleted_at").Comment("删除时间").Optional().Nillable(),
		field.Int("updated_by").Comment("最后更新者").Positive(),
		field.Time("updated_at").Comment("最后更新时间").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("employee_creator").Field("created_by").Unique().Required(),
		edge.From("admin", Admin.Type).Ref("employee_admin").Field("admin_id").Unique().Required(),
		edge.From("department", Department.Type).Field("department_id").Ref("employee_department").Unique().Required(),
	}
}
