package schema

import "entgo.io/ent"

// Dept holds the schema definition for the Dept entity.
type Dept struct {
	ent.Schema
}

// Fields of the Dept.
func (Dept) Fields() []ent.Field {
	return nil
}

// Edges of the Dept.
func (Dept) Edges() []ent.Edge {
	return nil
}
