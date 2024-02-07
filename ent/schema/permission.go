package schema

import (
	"aisecurity/ent/mixin"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

func (Permission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").NotEmpty().MaxLen(255).StructTag(`validate:"required"`),
		field.JSON("access_ids", []string{}).Comment("权限点").StructTag(`validate:"required"`),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("permission_creator").Unique().Field("creator_id").Required().Immutable(),
		edge.From("updater", Admin.Type).Ref("permission_updater").Unique().Field("updater_id").Required(),

		edge.To("admin", Admin.Type),
		edge.To("department", Department.Type),
	}
}
