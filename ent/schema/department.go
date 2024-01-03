package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Comment("创建时间").Default(time.Now).Immutable(),
		field.Int("created_by").Comment("创建者").Positive(),
		field.String("title").MaxLen(255).Comment("标题"),
		field.Int("parent_id").Optional().Positive().Comment("上级部门id"),
		field.Time("deleted_at").Comment("删除时间").Optional().Nillable(),
		field.Int("updated_by").Comment("最后更新者").Positive(),
		field.Time("updated_at").Comment("最后更新时间").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("department_creator").Unique().Field("created_by").Required(),
		edge.From("parent", Department.Type).Ref("children").Unique().Field("parent_id"),

		edge.To("employee_department", Employee.Type),
		edge.To("children", Department.Type),
	}
}
