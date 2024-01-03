package admin

import (
	"aisecurity/handlers"
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
	adminRoleHandler := auth.NewAdminRoleHandler()
	routes.RouterGroup.POST("/admin-role/create", adminRoleHandler.Create)
	routes.RouterGroup.GET("/admin-role/get-list", adminRoleHandler.GetList)
	routes.RouterGroup.GET("/admin-role/get-modules", adminRoleHandler.GetModules)

	adminHandler := auth.NewAdminHandler()
	{
		routes.RouterGroup.POST("/admin/create-super-admin", handlers.Convert(adminHandler, adminHandler.CreateSuperAdmin))
		routes.RouterGroup.POST("/admin/create", handlers.Convert(adminHandler, adminHandler.Create))
		routes.RouterGroup.POST("/admin/update", handlers.Convert(adminHandler, adminHandler.Update))
		routes.RouterGroup.GET("/admin/get-list", handlers.Convert(adminHandler, adminHandler.GetList))
	}
}
