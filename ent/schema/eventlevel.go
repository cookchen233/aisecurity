package schema

import (
	"aisecurity/ent/mixin"
	"aisecurity/enums"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EventLevel holds the schema definition for the EventLevel entity.
type EventLevel struct {
	ent.Schema
}

func (EventLevel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the EventLevel.
func (EventLevel) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").Optional().MaxLen(255).StructTag(`validate:"required"`),
		field.JSON("event_types", []enums.EventType{}).Comment("包含事件类型").StructTag(`validate:"required"`),
		field.String("description").Comment("描述").Optional(),
		field.Bool("is_report").Comment("是否上报").Optional().Default(false),
	}
}

// Edges of the EventLevel.
func (EventLevel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("event_level_creator").Field("created_by").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("event_level_updater").Field("updated_by").Required().Unique(),
	}
}
