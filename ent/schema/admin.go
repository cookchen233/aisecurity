package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable().Comment("创建时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("最后更新时间"),
		field.Time("deleted_at").Optional().Comment("删除时间"),
		field.Int("created_by").Positive().Comment("创建人").Optional(),
		field.String("username").MaxLen(64).Comment("用户名"),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("RiskReportingAdmin", RiskReporting.Type),
	}
}
