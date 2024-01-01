package routes

import (
	handlers "aisecurity/handlers/toycar"
	"aisecurity/utils"
	"github.com/gin-gonic/gin"
)

type ToycarRoutes struct {
	Router *gin.RouterGroup
}

func NewToycarRoutes(router *gin.RouterGroup) *ToycarRoutes {
	return &ToycarRoutes{
		Router: router,
	}
}

// Setup Register toycar routes
func (toycar *ToycarRoutes) Setup() {
	handler := handlers.NewToycarHandler()
	pool := utils.NewConnPool()

	toycar.Router.GET("/accept-control", utils.HandleWebSocket(handler.AcceptControl, pool, "accept-control"))
	toycar.Router.GET("/control-panel", utils.HandleWebSocket(handler.ControlPanel, pool, "control-panel"))
}
