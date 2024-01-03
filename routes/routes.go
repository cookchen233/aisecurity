package routes

import (
	"aisecurity/handlers"
	"aisecurity/handlers/admin/auth"
	"aisecurity/middlewares"
	admin "aisecurity/routes/admin"
	"github.com/gin-gonic/gin"
)

// Setup Register all routes
func Setup(router *gin.Engine) {

	// 静态文件
	router.Static("/public", "./public/")

	// Register Toycar handlers
	NewToycarRoutes(router.Group("/toycar")).Setup()

	// Register admin routes
	adminHandler := auth.NewAdminHandler()
	router.POST("/admin/login", handlers.Convert(adminHandler, adminHandler.Login)) // without checking session
	routes := router.Group("/admin", middlewares.IsAdminAuthorized())
	{
		admin.NewAuthRoutes(routes.Group("/auth")).Setup()
		admin.NewSecurityRoutes(routes.Group("/security")).Setup()
		admin.NewOrganizationRoutes(routes.Group("/organization")).Setup()
	}
}
