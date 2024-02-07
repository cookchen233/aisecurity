package routes

import (
	"aisecurity/handlers"
	"aisecurity/handlers/ipc"
	"aisecurity/middlewares"
	"github.com/gin-gonic/gin"
)

type IPCRoutes struct {
	RouterGroup *gin.RouterGroup
}

func NewIPCRoutes(router *gin.RouterGroup) *IPCRoutes {
	return &IPCRoutes{
		RouterGroup: router,
	}
}

// Setup Register admin handlers
func (routes *IPCRoutes) Setup() {
	twoHandler := ipc.NewTwoHandler()
	{
		routes.RouterGroup.POST("/two/report-event", handlers.Convert(twoHandler, twoHandler.ReportEvent))
		routes.RouterGroup.POST("/two/upload-videos", handlers.Convert(twoHandler, twoHandler.UploadVideos))
		//routes.RouterGroup.POST("/two/extra-messages", handlers.Convert(twoHandler, twoHandler.ExtraMessages))
	}

	fuyouHandler := ipc.NewFuyouHandler()
	{
		routes.RouterGroup.POST("/fuyou/report-event",
			middlewares.FuyouResponse(),
			handlers.Convert(fuyouHandler, fuyouHandler.ReportEvent),
		)
		routes.RouterGroup.POST("/fuyou/upload-videos",
			middlewares.FuyouResponse(),
			handlers.Convert(fuyouHandler, fuyouHandler.UploadVideos),
		)
		//routes.RouterGroup.POST("/two/extra-messages", handlers.Convert(twoHandler, twoHandler.ExtraMessages))
	}
}
