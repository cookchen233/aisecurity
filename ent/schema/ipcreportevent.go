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

// IPCReportEvent holds the schema definition for the IPCReportEvent entity.
type IPCReportEvent struct {
	ent.Schema
}

func (IPCReportEvent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the IPCReportEvent.
func (IPCReportEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("device_id").Comment("设备ID").MaxLen(255).Immutable().StructTag(`validate:"required"`),
		field.String("event_id").Comment("事件ID").MaxLen(255).Immutable().StructTag(`validate:"required"`),
		field.Time("event_time").Comment("事件发生时间").Default(time.Now).Immutable().StructTag(`validate:"required"`),
		field.Int("event_type").Comment("事件类型").NonNegative().Default(int(enums.IRETUnknown)).GoType(enums.IPCReportEventType(0)).StructTag(`validate:"required"`),
		field.Int("event_status").Comment("事件状态").NonNegative().Default(int(enums.IRESPending)).GoType(enums.IPCReportEventStatus(0)).StructTag(`validate:"required"`),
		field.JSON("images", []types.UploadedImage{}).Optional().Default([]types.UploadedImage{}).Comment("图片").StructTag(`validate:"required"`),
		field.JSON("labeled_images", []types.UploadedImage{}).Optional().Default([]types.UploadedImage{}).Comment("标记的图片").StructTag(`validate:"required"`),
		field.JSON("videos", []types.UploadedVideo{}).Optional().Default([]types.UploadedVideo{}).Comment("视频").StructTag(`validate:"required"`),
		field.String("description").Comment("描述").Optional(),
		field.String("raw_data").Comment("设备商原始上报数据").Optional(),
	}
}

// Edges of the IPCReportEvent.
func (IPCReportEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("ipc_report_event_creator").Field("created_by").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("ipc_report_event_updater").Field("updated_by").Required().Unique(),
	}
}
