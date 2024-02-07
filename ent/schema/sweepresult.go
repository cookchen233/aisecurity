package schema

import (
	"aisecurity/ent/mixin"
	"aisecurity/structs/types"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SweepResult holds the schema definition for the SweepResult entity.
type SweepResult struct {
	ent.Schema
}

func (SweepResult) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the SweepResult.
func (SweepResult) Fields() []ent.Field {
	return []ent.Field{
		field.Int("sweep_id").Comment("排查模板ID").Positive(),
		field.Int("sweep_schedule_id").Comment("排查任务ID").Positive(),
		field.Time("check_in_time").Comment("签到时间").Optional(),
		field.JSON("check_in_image", types.UploadedImage{}).Optional().Comment("签到照片"),
		field.JSON("sweep_jobs", []*types.SweepJob{}).Optional().Comment("排查内容"),
	}
}

// Edges of the SweepResult.
func (SweepResult) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("sweep_result_creator").Field("creator_id").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("sweep_result_updater").Field("updater_id").Required().Unique(),
		edge.From("sweep", Sweep.Type).Ref("sweep_result").Unique().Field("sweep_id").Required(),
		edge.From("sweep_schedule", SweepSchedule.Type).Ref("sweep_result").Unique().Field("sweep_schedule_id").Required(),

		edge.To("sweep_result_details", SweepResultDetails.Type),
	}
}
