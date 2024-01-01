package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// RiskLocation holds the schema definition for the RiskLocation entity.
type RiskLocation struct {
	ent.Schema
}

// Fields of the RiskLocation.
func (RiskLocation) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable().Comment("创建时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("最后更新时间"),
		field.Time("deleted_at").Optional().Comment("删除时间"),
		field.String("title").MaxLen(255).Comment("标题"),
	}
}

// Edges of the RiskLocation.
func (RiskLocation) Edges() []ent.Edge {
	return nil
}
