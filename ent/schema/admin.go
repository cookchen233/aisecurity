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
		field.String("username").Comment("用户名").MaxLen(64).StructTag(`validate:"required"`),
		field.String("password").Comment("密码").MaxLen(255).Sensitive(),
		field.String("name").Comment("名字").Optional().MaxLen(64).StructTag(`validate:"required"`),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("admin_creator").Field("created_by").Required().Unique().Immutable(),
		edge.From("updator", Admin.Type).Ref("admin_updator").Field("updated_by").Required().Unique(),
		edge.From("admin_roles", AdminRole.Type).Ref("admins"),

		edge.To("admin_creator", Admin.Type),
		edge.To("admin_updator", Admin.Type),
		edge.To("admin_role_creator", AdminRole.Type),
		edge.To("admin_role_updator", AdminRole.Type),
		edge.To("risk_creator", Risk.Type),
		edge.To("risk_updator", Risk.Type),
		edge.To("risk_maintainer", Risk.Type),
		edge.To("risk_location_creator", RiskLocation.Type),
		edge.To("risk_location_updator", RiskLocation.Type),
		edge.To("risk_category_creator", RiskCategory.Type),
		edge.To("risk_category_updator", RiskCategory.Type),
		edge.To("department_creator", Department.Type),
		edge.To("department_updator", Department.Type),
		edge.To("employee_creator", Employee.Type),
		edge.To("employee_updator", Employee.Type),
		edge.To("employee_admin", Employee.Type),
	}
}
