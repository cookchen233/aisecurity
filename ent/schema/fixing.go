package schema

import (
	"aisecurity/ent/mixin"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Fixing holds the schema definition for the Fixing entity.
type Fixing struct {
	ent.Schema
}

func (Fixing) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the Fixing.
func (Fixing) Fields() []ent.Field {
	return []ent.Field{
		field.Int("fixer_id").Comment("处理人ID").Positive(),
		field.Int("event_id").Comment("事件ID").Positive().Immutable(),
		field.Int("device_id").Comment("设备ID").Positive().Immutable(),
		field.String("assign_notes").Comment("分派备注").Optional().Immutable(),
		field.Time("fix_time").Comment("处理时间").Optional(),
		field.Time("complete_time").Comment("处理完成时间").Optional(),
		field.String("complete_notes").Comment("处理完成备注").Optional(),
	}
}

// Edges of the Fixing.
func (Fixing) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("fixing_creator").Field("creator_id").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("fixing_updater").Field("updater_id").Required().Unique(),
		edge.From("fixer", Admin.Type).Ref("fixer").Field("fixer_id").Unique().Required(),
		edge.From("event", Event.Type).Ref("fixing").Field("event_id").Unique().Required().Immutable(),
		edge.From("device", Device.Type).Ref("fixing").Field("device_id").Unique().Required().Immutable(),
	}
}
