package schema

import (
	"aisecurity/ent/mixin"
	"aisecurity/properties"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Risk holds the schema definition for the Risk entity.
type Risk struct {
	ent.Schema
}

func (Risk) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Fields of the Risk.
func (Risk) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(255).Comment("标题"),
		field.String("content").Comment("内容"),
		field.JSON("images", []struct {
			Title string `json:"title"`
			URL   string `json:"url"`
		}{}).Optional().Comment("图片"),
		field.Int("risk_category_id").Positive().Comment("风险类别"),
		field.Int("risk_location_id").Positive().Comment("地点"),
		field.Int("maintainer").Positive().Comment("整改人"),
		field.String("measures").Optional().Comment("整改措施"),
		field.Int("maintain_status").Positive().Default(0).GoType(properties.MaintainStatus(0)).Comment("整改状态"),
		field.Time("due_time").Comment("计划完成日期"),
	}
}

// Edges of the Risk.
func (Risk) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("risk_creator").Field("created_by").Immutable().Unique().Required(),
		edge.From("maintainer_admin", Admin.Type).Ref("risk_maintainer").Unique().Field("maintainer").Required(),
		edge.From("category", RiskCategory.Type).Ref("risk_category").Unique().Field("risk_category_id").Required(),
		edge.From("location", RiskLocation.Type).Ref("risk_location").Unique().Field("risk_location_id").Required(),
	}
}

// Annotations of the Risk.
func (Risk) Annotations() []schema.Annotation {
	return nil
	//return []schema.Annotation{
	//	entsql.Annotation{Table: "risk"},
	//}
}
