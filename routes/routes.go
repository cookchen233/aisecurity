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
	router.POST("/api/dashboard/login", dashboard2.Convert(indexHandler, indexHandler.Login))
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
