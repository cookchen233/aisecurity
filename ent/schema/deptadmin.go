package schema

import "entgo.io/ent"

// DeptAdmin holds the schema definition for the DeptAdmin entity.
type DeptAdmin struct {
	ent.Schema
}

// Fields of the DeptAdmin.
func (DeptAdmin) Fields() []ent.Field {
	return nil
}

// Edges of the DeptAdmin.
func (DeptAdmin) Edges() []ent.Edge {
	return nil
}
