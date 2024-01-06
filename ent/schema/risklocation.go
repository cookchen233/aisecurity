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
		field.String("title").MaxLen(255).Comment("标题").NotEmpty().StructTag(`validate:"required"`),
	}
}

// Edges of the RiskLocation.
func (RiskLocation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("risk_location_creator").Field("created_by").Immutable().Unique().Required(),
		edge.From("updator", Admin.Type).Ref("risk_location_updator").Field("updated_by").Required().Unique(),

		edge.To("risk_risk_location", Risk.Type),
	}
}
