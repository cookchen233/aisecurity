package routes

import (
	"aisecurity/handlers"
	"github.com/gin-gonic/gin"
)

type WechatRoutes struct {
	RouterGroup *gin.RouterGroup
}

func NewWechatRoutes(router *gin.RouterGroup) *WechatRoutes {
	return &WechatRoutes{
		RouterGroup: router,
	}
}

// Setup Register admin handlers
func (routes *WechatRoutes) Setup() {
	wechatHandler := handlers.NewWechatHandler()
	{
		routes.RouterGroup.GET("/official-account", handlers.Convert(wechatHandler, wechatHandler.OfficialAccount))
		routes.RouterGroup.POST("/official-account", handlers.Convert(wechatHandler, wechatHandler.OfficialAccount))
		routes.RouterGroup.GET("/oauth2", handlers.Convert(wechatHandler, wechatHandler.AdminOauth2))
	}
}
