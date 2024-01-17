package schema

import (
	"aisecurity/ent/mixin"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Video holds the schema definition for the Video entity.
type Video struct {
	ent.Schema
}

func (Video) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the Video.
func (Video) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").Optional(),
		field.String("url").Comment("文件地址").Optional(),
		field.Int64("size").Comment("文件大小").Default(0),
		field.String("duration").Comment("视频时长").Optional(),
		field.Time("uploaded_at").Comment("上传时间").Optional().Nillable(),
	}
}

// Edges of the Video.
func (Video) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("video_creator").Field("created_by").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("video_updater").Field("updated_by").Required().Unique(),

		edge.To("ipc_report_event_video", IPCReportEvent.Type),
	}
}
