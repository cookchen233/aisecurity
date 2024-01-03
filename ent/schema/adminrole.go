package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// AdminRole holds the schema definition for the AdminRole entity.
type AdminRole struct {
	ent.Schema
}

// Fields of the AdminRole.
func (AdminRole) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Comment("创建时间").Default(time.Now).Immutable(),
		field.Int("created_by").Comment("创建者").Optional().Positive(),
		field.String("name").MaxLen(255).Comment("名称"),
		field.Time("deleted_at").Comment("删除时间").Optional().Nillable(),
		field.Int("updated_by").Comment("最后更新者").Optional().Positive(),
		field.Time("updated_at").Comment("最后更新时间").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the AdminRole.
func (AdminRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("admin_role_creator").Unique().Field("created_by"),
		edge.From("updator", Admin.Type).Ref("admin_role_updator").Unique().Field("updated_by"),

		edge.To("admins", Admin.Type),
	}
}
