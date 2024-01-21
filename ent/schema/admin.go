package schema

import (
	"aisecurity/ent/mixin"
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
		field.String("avatar").Comment("头像").Optional().MaxLen(255),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("admin_creator").Field("created_by").Required().Unique().Immutable(),
		edge.From("updater", Admin.Type).Ref("admin_updater").Field("updated_by").Required().Unique(),
		edge.From("admin_roles", AdminRole.Type).Ref("admins"),

		edge.To("admin_creator", Admin.Type),
		edge.To("admin_updater", Admin.Type),

		edge.To("admin_role_creator", AdminRole.Type),
		edge.To("admin_role_updater", AdminRole.Type),

		edge.To("risk_creator", Risk.Type),
		edge.To("risk_updater", Risk.Type),

		edge.To("risk_location_creator", RiskLocation.Type),
		edge.To("risk_location_updater", RiskLocation.Type),

		edge.To("risk_category_creator", RiskCategory.Type),
		edge.To("risk_category_updater", RiskCategory.Type),

		edge.To("department_creator", Department.Type),
		edge.To("department_updater", Department.Type),

		edge.To("employee_creator", Employee.Type),
		edge.To("employee_updater", Employee.Type),
		edge.To("employee", Employee.Type),

		edge.To("occupation_creator", Occupation.Type),
		edge.To("occupation_updater", Occupation.Type),

		edge.To("ipc_event_creator", IPCEvent.Type),
		edge.To("ipc_event_updater", IPCEvent.Type),

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
	}
}
