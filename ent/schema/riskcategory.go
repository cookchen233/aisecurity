package schema

import "entgo.io/ent"

// RiskCategory holds the schema definition for the RiskCategory entity.
type RiskCategory struct {
	ent.Schema
}

// Fields of the RiskCategory.
func (RiskCategory) Fields() []ent.Field {
	return nil
}

// Edges of the RiskCategory.
func (RiskCategory) Edges() []ent.Edge {
	return nil
}
