package schema

import (
	"aisecurity/ent/mixin"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DeviceInstallation holds the schema definition for the DeviceInstallation entity.
type DeviceInstallation struct {
	ent.Schema
}

func (DeviceInstallation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the DeviceInstallation.
func (DeviceInstallation) Fields() []ent.Field {
	return []ent.Field{
		field.Int("device_id").Comment("设备ID").NonNegative().StructTag(`validate:"required"`),
		field.Int("area_id").Comment("区域ID").NonNegative().StructTag(`validate:"required"`),
		field.String("alias_name").Comment("设备别名").MaxLen(255).Optional(),
		field.Float("longitude").Comment("经度").StructTag(`validate:"required"`),
		field.Float("latitude").Comment("纬度").StructTag(`validate:"required"`),
		field.String("location_data").Comment("其他位置数据").Optional(), // could store a JSON string representing a geo object
		field.String("location").Comment("位置描述").Optional(),
		field.String("installer").Comment("安装人").Optional(),
		field.Time("install_time").Comment("安装时间").Optional(),
	}
}

// Edges of the DeviceInstallation.
func (DeviceInstallation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("device_installation_creator").Field("creator_id").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("device_installation_updater").Field("updater_id").Required().Unique(),
		edge.From("area", Area.Type).Ref("device_installation").Field("area_id").Required().Unique(),
		edge.From("device", Device.Type).Ref("device_installation").Field("device_id").Required().Unique(),
	}
}
