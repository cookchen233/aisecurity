package schema

import (
	"aisecurity/ent/mixin"
	"aisecurity/structs/types"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Sweep holds the schema definition for the Sweep entity.
type Sweep struct {
	ent.Schema
}

func (Sweep) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the Sweep.
func (Sweep) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").NotEmpty().MaxLen(255).StructTag(`validate:"required"`),
		field.Int("risk_category_id").Comment("隐患类别").Positive(),
		field.Int("risk_location_id").Comment("隐患地点").Positive(),
		field.JSON("sweep_jobs", []*types.SweepJob{}).Optional().Comment("排查内容"),
	}
}

// Edges of the Sweep.
func (Sweep) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("sweep_creator").Field("creator_id").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("sweep_updater").Field("updater_id").Required().Unique(),

		edge.From("risk_category", RiskCategory.Type).Ref("sweep").Unique().Field("risk_category_id").Required(),
		edge.From("risk_location", RiskLocation.Type).Ref("sweep").Unique().Field("risk_location_id").Required(),

		edge.To("sweep_schedule", SweepSchedule.Type),
		edge.To("sweep_result", SweepResult.Type),
		edge.To("sweep_result_details", SweepResultDetails.Type),
	}
}
