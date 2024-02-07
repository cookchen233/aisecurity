package schema

import (
	"aisecurity/ent/mixin"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Area holds the schema definition for the Area entity.
type Area struct {
	ent.Schema
}

func (Area) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the Area.
func (Area) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").MaxLen(255).Optional().Default("未命名区域"),
		field.String("description").Comment("描述").Optional(),
	}
}

// Edges of the Area.
func (Area) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("area_creator").Field("creator_id").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("area_updater").Field("updater_id").Required().Unique(),

		edge.To("device_installation", DeviceInstallation.Type),
	}
}
