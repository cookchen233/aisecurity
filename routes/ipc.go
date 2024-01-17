package routes

import (
	"aisecurity/handlers"
	"aisecurity/handlers/ipc"
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
}
