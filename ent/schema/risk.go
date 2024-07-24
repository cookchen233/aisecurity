package schema

import (
	"aisecurity/ent/mixin"
	"aisecurity/properties/maintain_status"
	"aisecurity/structs/types"
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
		field.String("title").Comment("标题").NotEmpty().MaxLen(255).StructTag(`validate:"required"`),
		field.String("content").Comment("内容").Optional(),
		field.JSON("images", []types.UploadedImage{}).Optional().Default([]types.UploadedImage{}).Comment("图片"),
		field.JSON("maintained_images", []types.UploadedImage{}).Optional().Default([]types.UploadedImage{}).Comment("整改后的图片"),
		field.Int("risk_category_id").Comment("隐患类别").Positive(),
		field.Int("risk_location_id").Comment("隐患地点").Positive(),
		field.Int("maintainer_id").Comment("整改人").Positive(),
		field.String("measures").Comment("整改措施").Optional(),
		field.Int("maintain_status").Comment("整改状态").NonNegative().Default(int(maintain_status.Unknown)).GoType(maintain_status.MaintainStatus(0)),
		field.Time("due_time").Comment("计划完成日期"),
	}
}

// Edges of the Risk.
func (Risk) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", Admin.Type).Ref("risk_creator").Field("creator_id").Immutable().Unique().Required(),
		edge.From("updater", Admin.Type).Ref("risk_updater").Field("updater_id").Required().Unique(),
		edge.From("risk_category", RiskCategory.Type).Ref("risk").Unique().Field("risk_category_id").Required(),
		edge.From("risk_location", RiskLocation.Type).Ref("risk").Unique().Field("risk_location_id").Required(),
		edge.From("maintainer", Admin.Type).Ref("risk_maintainer").Field("maintainer_id").Unique().Required(),
	}
}

// Annotations of the Risk.
func (Risk) Annotations() []schema.Annotation {
	return nil
	//return []schema.Annotation{
	//	entsql.Annotation{Table: "risk"},
	//}
}
