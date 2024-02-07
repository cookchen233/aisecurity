package schema

import (
	"aisecurity/ent/mixin"
	"aisecurity/enums"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

func (Device) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.Int("brand").Comment("设备品牌").Optional().NonNegative().Default(int(enums.DBUnknown)).GoType(enums.DeviceBrand(0)).StructTag(`validate:"required"`),
		field.Int("model").Comment("设备型号").Optional().NonNegative().Default(int(enums.DMUnknown)).GoType(enums.DeviceModel(0)).StructTag(`validate:"required"`),
		field.String("name").Comment("设备名称").Optional().Default("未命名设备").MaxLen(255).StructTag(`validate:"required"`),
		field.String("sn").Comment("设备序列号").Optional().MaxLen(255).StructTag(`validate:"required"`),
		field.Int("device_type").Comment("设备类型").NonNegative().Default(int(enums.DTUnknown)).GoType(enums.DeviceType(0)).StructTag(`validate:"required"`),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("device_creator").Field("creator_id").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("device_updater").Field("updater_id").Required().Unique(),

		edge.To("event", Event.Type),
		edge.To("device_installation", DeviceInstallation.Type),
		edge.To("event_log", EventLog.Type),
		edge.To("fixing", Fixing.Type),
	}
}
