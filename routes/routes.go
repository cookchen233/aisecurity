package routes

import (
	dashboard2 "aisecurity/handlers"
	"aisecurity/handlers/dashboard"
	"aisecurity/middlewares"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// Setup Register all routes
func Setup(router *gin.Engine) {

	// 静态文件
	router.Use(static.Serve("/", static.LocalFile("./home", false)))

	// Register Toycar handlers
	NewToycarRoutes(router.Group("/toycar")).Setup()

	// Register dashboard handlers
	indexHandler := dashboard.NewIndexHandler()
	// Without checking session
	router.POST("/api/dashboard/create-super-admin", dashboard2.Convert(indexHandler, indexHandler.CreateSuperAdmin))
	router.POST("/api/dashboard/login", dashboard2.Convert(indexHandler, indexHandler.Login))
	router.GET("/api/dashboard/current-admin",
		middlewares.IsAdminAuthorized(),
		dashboard2.Convert(indexHandler, indexHandler.GetCurrentAdmin),
	)
	// Need to check session
	NewDashboardRoutes(router.Group("/api/dashboard",
		middlewares.IsAdminAuthorized(),
		middlewares.DatabaseAudit(),
		middlewares.JoyRequestLog(),
	)).Setup()

	// IPC
	NewIPCRoutes(router.Group("/api/ipc",
		middlewares.JoyRequestLog(),
	)).Setup()
}
