package schema

import (
	"aisecurity/ent/mixin"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SweepResultDetails holds the schema definition for the SweepResultDetails entity.
type SweepResultDetails struct {
	ent.Schema
}

func (SweepResultDetails) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the SweepResultDetails.
func (SweepResultDetails) Fields() []ent.Field {
	return []ent.Field{
		field.Int("sweep_id").Comment("排查模板ID").Positive(),
		field.Int("sweep_schedule_id").Comment("排查任务ID").Positive(),
		field.Int("sweep_result_id").Comment("排查结果ID").Positive(),
		field.String("title").Comment("排查项标题").NotEmpty().MaxLen(255).StructTag(`validate:"required"`),
		field.Int("result").Comment("排查结果").Positive(),
	}
}

// Edges of the SweepResultDetails.
func (SweepResultDetails) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("sweep_result_details_creator").Field("creator_id").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("sweep_result_details_updater").Field("updater_id").Required().Unique(),
		edge.From("sweep", Sweep.Type).Ref("sweep_result_details").Unique().Field("sweep_id").Required(),
		edge.From("sweep_schedule", SweepSchedule.Type).Ref("sweep_result_details").Unique().Field("sweep_schedule_id").Required(),
		edge.From("sweep_result", SweepResult.Type).Ref("sweep_result_details").Unique().Field("sweep_result_id").Required(),
	}
}
