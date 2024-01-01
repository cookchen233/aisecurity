package admin

import (
	"aisecurity/handlers/admin/auth"
	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	RouterGroup *gin.RouterGroup
}

func NewAuthRoutes(router *gin.RouterGroup) *AuthRoutes {
	return &AuthRoutes{
		RouterGroup: router,
	}
}

// Setup Register auth handlers
func (routes *AuthRoutes) Setup() {
	routes.RouterGroup.GET("/role/index", auth.NewRoleHandler().Index)
	routes.RouterGroup.GET("/admin/create", auth.NewAdminHandler().Create)
}
