package routes

import (
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
	adminRoutes := router.Group("/admin")
	{
		// Register auth routes
		admin.NewAuthRoutes(adminRoutes.Group("/auth")).Setup()
		admin.NewSecurityRoutes(adminRoutes.Group("/security")).Setup()
	}
}
