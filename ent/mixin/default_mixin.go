package mixin

import (
	"aisecurity/utils"
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

// DefaultMixin implements the ent.Mixin for sharing
// audit-log capabilities with package schemas.
type DefaultMixin struct {
	mixin.Schema
}

// Fields of the Admin.
func (DefaultMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").Comment("创建时间").Default(time.Now).Immutable(),
		field.Time("delete_time").Comment("删除时间").Optional().Nillable(),
		field.Time("update_time").Comment("最后更新时间").Default(time.Now).UpdateDefault(time.Now),
		field.Int("updater_id").Comment("最后更新者").Positive(),
	}
}

// Hooks of the DefaultMixin.
func (DefaultMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		DefaultHook,
	}
}

func DefaultHook(next ent.Mutator) ent.Mutator {
	// DefaultLogger wraps the methods that are shared between all mutations of
	// schemas that embed the DefaultLog mixin. The variable "exists" is true, if
	// the field already exists in the mutation (e.g. was set by a different hook).
	type DefaultLogger interface {
		SetUpdaterID(int)
		UpdaterID() (id int, exists bool)
	}

	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		ml, ok := m.(DefaultLogger)
		if !ok {
			return nil, errors.WithStack(fmt.Errorf("unexpected audit-log call from mutation type %T", m))
		}
		// Extracting Gin context and ensuring its presence
		ginCtx, ok := ctx.(*gin.Context)
		if !ok {
			return nil, errors.WithStack(fmt.Errorf("unexpected context type %T", ctx))
		}

		// set the admin id
		adminID := max(1, ginCtx.GetInt("admin_id"))
		if adminID > 0 {
			switch op := m.Op(); {
			case op.Is(ent.OpCreate):
				if _, exists := ml.UpdaterID(); !exists {
					ml.SetUpdaterID(adminID)
				}
			case op.Is(ent.OpUpdateOne | ent.OpUpdate):
				if _, exists := ml.UpdaterID(); !exists {
					ml.SetUpdaterID(adminID)
				}
			}
		}

		// record data to audit log
		type fieldData struct {
			Field string
			Value ent.Value
		}
		var fields []fieldData
		for _, k := range m.Fields() {
			v, _ := m.Field(k)
			fields = append(fields, fieldData{
				k, v,
			})
		}
		if ginCtx.GetBool("is_audit") {
			utils.Logger.Info("audit log",
				zap.Int("admin_id", adminID),
				zap.String("action", m.Op().String()),
				zap.String("entity", m.Type()),
				zap.String("path", ginCtx.FullPath()),
				zap.String("ip", ginCtx.ClientIP()),
				zap.Any("fields", fields),
			)
		}

		return next.Mutate(ctx, m)
	})
}
