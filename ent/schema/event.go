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

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

func (Event) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.Int("device_id").Comment("设备ID").NonNegative().Immutable().StructTag(`validate:"required"`),
		field.Int("video_id").Comment("视频ID").Optional().NonNegative().Immutable(),
		field.Time("event_time").Comment("事件发生时间").Default(time.Now).Immutable().StructTag(`validate:"required"`),
		field.Int("event_type").Comment("事件类型").NonNegative().Default(int(enums.ETUnknown)).GoType(enums.EventType(0)).Immutable().StructTag(`validate:"required"`),
		field.Int("event_status").Comment("事件状态").NonNegative().Default(int(enums.ES1)).GoType(enums.EventStatus(0)),
		field.JSON("images", []*types.UploadedImage{}).Optional().Default([]*types.UploadedImage{}).Comment("图片").Immutable().StructTag(`validate:"required"`),
		field.JSON("labeled_images", []*types.UploadedImage{}).Optional().Default([]*types.UploadedImage{}).Comment("标记的图片").Immutable(),
		field.String("data_id").Comment("请求数据ID").MaxLen(255).Immutable().StructTag(`validate:"required"`),
		field.String("description").Comment("描述").Optional().Immutable(),
		field.String("raw_data").Comment("设备商原始上报数据").Optional().Immutable().StructTag(`json:"-"`),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("event_creator").Field("creator_id").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("event_updater").Field("updater_id").Required().Unique(),
		edge.From("video", Video.Type).Ref("event").Field("video_id").Immutable().Unique(),
		edge.From("device", Device.Type).Ref("event").Field("device_id").Immutable().Required().Unique(),

		edge.To("fixing", Fixing.Type).Unique(),
		edge.To("event_log", EventLog.Type),
	}
}
