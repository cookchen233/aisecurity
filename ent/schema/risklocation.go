package schema

import (
	"aisecurity/ent/mixin"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RiskLocation holds the schema definition for the RiskLocation entity.
type RiskLocation struct {
	ent.Schema
}

func (RiskLocation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the RiskLocation.
func (RiskLocation) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("名称").NotEmpty().MaxLen(255).StructTag(`validate:"required"`),
	}
}

// Edges of the RiskLocation.
func (RiskLocation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("risk_location_creator").Field("creator_id").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("risk_location_updater").Field("updater_id").Required().Unique(),

		edge.To("risk", Risk.Type),
		edge.To("sweep", Sweep.Type),
	}
}
