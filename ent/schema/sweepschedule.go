package schema

import (
	"aisecurity/ent/mixin"
	"aisecurity/enums"
	"aisecurity/structs/types"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SweepSchedule holds the schema definition for the SweepSchedule entity.
type SweepSchedule struct {
	ent.Schema
}

func (SweepSchedule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the SweepSchedule.
func (SweepSchedule) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").NotEmpty().MaxLen(255).StructTag(`validate:"required"`),
		field.Int("sweep_id").Comment("排查ID").Positive(),
		field.Int("schedule_status").Comment("任务状态").NonNegative().Optional().Default(int(enums.SS1)).GoType(enums.AdminStatus(0)),
		field.Time("action_time").Comment("任务开始时间"),
		field.JSON("remind", types.ScheduleRemind{}).Comment("提醒"),
		field.JSON("repeat", types.ScheduleRepeat{}).Comment("重复").StructTag(`validate:"required"`),
	}
}

// Edges of the SweepSchedule.
func (SweepSchedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("sweep_schedule_creator").Field("creator_id").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("sweep_schedule_updater").Field("updater_id").Required().Unique(),
		edge.From("sweep", Sweep.Type).Ref("sweep_schedule").Unique().Field("sweep_id").Required(),
		edge.From("workers", Admin.Type).Ref("sweep_schedule"),

		edge.To("sweep_result", SweepResult.Type),
		edge.To("sweep_result_details", SweepResultDetails.Type),
	}
}
