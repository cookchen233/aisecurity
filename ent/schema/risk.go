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
		field.Int("risk_category_id").Comment("风险类别").Positive(),
		field.Int("risk_location_id").Comment("地点").Positive(),
		field.Int("maintainer_id").Comment("整改人").Positive(),
		field.String("measures").Comment("整改措施").Optional(),
		field.Int("maintain_status").Comment("整改状态").Positive().Default(int(properties.Unknown)).GoType(properties.MaintainStatus(0)),
		field.Time("due_time").Comment("计划完成日期"),
	}
}

// Edges of the Risk.
func (Risk) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Employee.Type).Ref("risk_creator").Field("created_by").Immutable().Unique().Required(),
		edge.From("updator", Admin.Type).Ref("risk_updator").Field("updated_by").Required().Unique(),
		edge.From("maintainer", Employee.Type).Ref("risk_maintainer").Field("maintainer_id").Unique().Required(),
		edge.From("risk_category", RiskCategory.Type).Ref("risk_risk_category").Unique().Field("risk_category_id").Required(),
		edge.From("risk_location", RiskLocation.Type).Ref("risk_risk_location").Unique().Field("risk_location_id").Required(),
	}
}

// Annotations of the Risk.
func (Risk) Annotations() []schema.Annotation {
	return nil
	//return []schema.Annotation{
	//	entsql.Annotation{Table: "risk"},
	//}
}
