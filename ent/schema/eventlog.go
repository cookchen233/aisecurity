package schema

import (
	"aisecurity/ent/mixin"
	"aisecurity/enums"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// EventLog holds the schema definition for the EventLog entity.
type EventLog struct {
	ent.Schema
}

func (EventLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the EventLog.
func (EventLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("device_id").Comment("设备ID").NonNegative().Immutable().StructTag(`validate:"required"`),
		field.Int("event_id").Comment("事件ID").NonNegative().Immutable().StructTag(`validate:"required"`),
		field.Int("actor_id").Comment("当事人").Optional().NonNegative().Immutable(),
		field.Int("actor2_id").Comment("当事人2").Optional().NonNegative().Immutable(),
		field.Int("log_type").Comment("日志类型").NonNegative().Default(int(enums.ELTUnknown)).GoType(enums.EventLogType(0)).Immutable().StructTag(`validate:"required"`),
		field.String("notes").Comment("备注").Optional().Immutable(),
	}
}

// Edges of the EventLog.
func (EventLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("event_log_creator").Field("creator_id").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("event_log_updater").Field("updater_id").Required().Unique(),
		edge.From("event", Event.Type).Ref("event_log").Field("event_id").Immutable().Required().Unique(),
		edge.From("device", Device.Type).Ref("event_log").Field("device_id").Immutable().Required().Unique(),
		edge.From("actor", Admin.Type).Ref("event_log_actor").Field("actor_id").Immutable().Unique(),
		edge.From("actor2", Admin.Type).Ref("event_log_actor2").Field("actor2_id").Immutable().Unique(),
	}
}
