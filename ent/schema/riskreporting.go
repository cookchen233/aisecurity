package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// RiskReporting holds the schema definition for the RiskReporting entity.
type RiskReporting struct {
	ent.Schema
}

// Fields of the RiskReporting.
func (RiskReporting) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable().Comment("创建时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("最后更新时间"),
		field.Time("deleted_at").Optional().Comment("删除时间"),
		field.Int("created_by").Positive().Comment("创建人").Optional(),
		field.String("title").MaxLen(255).Comment("标题"),
		field.String("content").Comment("内容"),
		field.JSON("images", []struct {
			Title string `json:"title"`
			URL   string `json:"url"`
		}{}).Comment("图片"),
		field.Int("risk_category_id").Positive().Comment("风险类别"),
		field.Int("risk_location_id").Positive().Comment("地点"),
		field.Int("maintainer").Positive().Comment("整改人"),
		field.String("measures").Optional().Nillable().Comment("整改措施"),
		field.Time("due_time").Comment("计划完成日期"),
	}
}

// Edges of the RiskReporting.
func (RiskReporting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("AdminRiskReporting", Admin.Type).Ref("RiskReportingAdmin").Unique().Field("created_by"),
	}
}

// Annotations of the RiskReporting.
func (RiskReporting) Annotations() []schema.Annotation {
	return nil
	//return []schema.Annotation{
	//	entsql.Annotation{Table: "risk_reporting"},
	//}
}
