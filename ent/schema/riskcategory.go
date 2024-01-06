package schema

import (
	"aisecurity/ent/mixin"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RiskCategory holds the schema definition for the RiskCategory entity.
type RiskCategory struct {
	ent.Schema
}

func (RiskCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the RiskCategory.
func (RiskCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(255).Comment("标题").NotEmpty().StructTag(`validate:"required"`),
	}
}

// Edges of the RiskCategory.
func (RiskCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("risk_category_creator").Field("created_by").Immutable().Unique().Required(),
		edge.From("updator", Admin.Type).Ref("risk_category_updator").Field("updated_by").Required().Unique(),

		edge.To("risk_risk_category", Risk.Type),
	}
}
