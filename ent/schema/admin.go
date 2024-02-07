package schema

import (
	"aisecurity/ent/mixin"
	"aisecurity/enums"
	"aisecurity/structs/types"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

func (Admin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Comment("用户名").NotEmpty().MaxLen(64).StructTag(`validate:"required"`),
		field.String("password").Comment("密码").NotEmpty().MaxLen(255).Sensitive(),
		field.String("nickname").Comment("昵称").NotEmpty().MaxLen(255).StructTag(`validate:"required"`),
		field.String("real_name").Comment("姓名").MaxLen(255).StructTag(`validate:"required"`),
		field.String("mobile").Comment("手机号").MaxLen(255).Optional(),
		field.JSON("avatar", types.UploadedImage{}).Optional().Comment("头像"),
		field.Int("admin_status").Comment("账户状态").NonNegative().Default(int(enums.AS1)).GoType(enums.AdminStatus(0)).StructTag(`validate:"required"`),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("admin_creator").Field("creator_id").Required().Unique().Immutable(),
		edge.From("updater", Admin.Type).Ref("admin_updater").Field("updater_id").Required().Unique(),
		edge.From("permissions", Permission.Type).Ref("admin"),

		edge.To("admin_creator", Admin.Type),
		edge.To("admin_updater", Admin.Type),

		edge.To("permission_creator", Permission.Type),
		edge.To("permission_updater", Permission.Type),

		edge.To("risk_creator", Risk.Type),
		edge.To("risk_updater", Risk.Type),
		edge.To("risk_maintainer", Risk.Type),

		edge.To("risk_location_creator", RiskLocation.Type),
		edge.To("risk_location_updater", RiskLocation.Type),

		edge.To("risk_category_creator", RiskCategory.Type),
		edge.To("risk_category_updater", RiskCategory.Type),

		edge.To("department_creator", Department.Type),
		edge.To("department_updater", Department.Type),

		edge.To("employee_creator", Employee.Type),
		edge.To("employee_updater", Employee.Type),
		edge.To("employee", Employee.Type).Unique(),

		edge.To("occupation_creator", Occupation.Type),
		edge.To("occupation_updater", Occupation.Type),

		edge.To("event_creator", Event.Type),
		edge.To("event_updater", Event.Type),

		edge.To("video_creator", Video.Type),
		edge.To("video_updater", Video.Type),

		edge.To("area_creator", Area.Type),
		edge.To("area_updater", Area.Type),

		edge.To("device_creator", Device.Type),
		edge.To("device_updater", Device.Type),

		edge.To("device_installation_creator", DeviceInstallation.Type),
		edge.To("device_installation_updater", DeviceInstallation.Type),

		edge.To("event_level_creator", EventLevel.Type),
		edge.To("event_level_updater", EventLevel.Type),

		edge.To("fixing_creator", Fixing.Type),
		edge.To("fixing_updater", Fixing.Type),
		edge.To("fixer", Fixing.Type),

		edge.To("event_log_creator", EventLog.Type),
		edge.To("event_log_updater", EventLog.Type),
		edge.To("event_log_actor", EventLog.Type),
		edge.To("event_log_actor2", EventLog.Type),

		edge.To("sweep_creator", Sweep.Type),
		edge.To("sweep_updater", Sweep.Type),

		edge.To("sweep_schedule_creator", SweepSchedule.Type),
		edge.To("sweep_schedule_updater", SweepSchedule.Type),
		edge.To("sweep_schedule", SweepSchedule.Type),

		edge.To("sweep_result_creator", SweepResult.Type),
		edge.To("sweep_result_updater", SweepResult.Type),

		edge.To("sweep_result_details_creator", SweepResultDetails.Type),
		edge.To("sweep_result_details_updater", SweepResultDetails.Type),
	}
}
