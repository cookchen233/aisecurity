package schema

import (
	"aisecurity/ent/mixin"
	"aisecurity/enums"
	"aisecurity/structs/types"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.DefaultMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Comment("用户名").NotEmpty().MaxLen(64).StructTag(`validate:"required"`),
		field.String("password").Comment("密码").MaxLen(255).Sensitive(),
		field.String("nickname").Comment("昵称").MaxLen(255).StructTag(`validate:"required"`),
		field.String("real_name").Comment("姓名").MaxLen(255).StructTag(`validate:"required"`),
		field.String("mobile").Comment("手机号").MaxLen(255).Optional(),
		field.String("wechat_openid").Comment("微信openid").MaxLen(255).Optional(),
		field.JSON("avatar", types.UploadedImage{}).Optional().Comment("头像"),
		field.Int("user_status").Comment("账户状态").NonNegative().Default(int(enums.ENS1)).GoType(enums.EnabledStatus(0)).StructTag(`validate:"required"`),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("updater", Admin.Type).Ref("user_updater").Field("updater_id").Required().Unique(),
	}
}
