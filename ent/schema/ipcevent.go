package schema

import (
	"aisecurity/ent/mixin"
	"aisecurity/enums"
	"aisecurity/structs/types"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// IPCEvent holds the schema definition for the IPCEvent entity.
type IPCEvent struct {
	ent.Schema
}

func (IPCEvent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the IPCEvent.
func (IPCEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Int("device_id").Comment("设备ID").NonNegative().Immutable().StructTag(`validate:"required"`),
		field.Int("video_id").Comment("视频ID").Optional().NonNegative().Immutable(),
		field.Time("event_time").Comment("事件发生时间").Default(time.Now).Immutable().StructTag(`validate:"required"`),
		field.Int("event_type").Comment("事件类型").NonNegative().Default(int(enums.ETUnknown)).GoType(enums.EventType(0)).Immutable().StructTag(`validate:"required"`),
		field.Int("event_status").Comment("事件状态").NonNegative().Default(int(enums.ESPending)).GoType(enums.EventStatus(0)).Immutable().StructTag(`validate:"required"`),
		field.JSON("images", []*types.UploadedImage{}).Optional().Default([]*types.UploadedImage{}).Comment("图片").StructTag(`validate:"required"`),
		field.JSON("labeled_images", []*types.UploadedImage{}).Optional().Default([]*types.UploadedImage{}).Comment("标记的图片").StructTag(`validate:"required"`),
		field.String("event_id").Comment("事件ID").MaxLen(255).Immutable().StructTag(`validate:"required"`),
		field.String("description").Comment("描述").Optional(),
		field.String("raw_data").Comment("设备商原始上报数据").Optional().StructTag(`json:"-"`),
	}
}

// Edges of the IPCEvent.
func (IPCEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("ipc_event_creator").Field("created_by").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("ipc_event_updater").Field("updated_by").Required().Unique(),
		edge.From("video", Video.Type).Ref("ipc_event_video").Field("video_id").Immutable().Unique(),
		edge.From("device", Device.Type).Ref("ipc_event_device").Field("device_id").Immutable().Required().Unique(),

		edge.To("fixers", Employee.Type),
	}
}
