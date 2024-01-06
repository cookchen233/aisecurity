package admin

import (
	"aisecurity/handlers"
	"aisecurity/handlers/admin/security"
	"github.com/gin-gonic/gin"
)

type SecurityRoutes struct {
	RouterGroup *gin.RouterGroup
}

func NewSecurityRoutes(router *gin.RouterGroup) *SecurityRoutes {
	return &SecurityRoutes{
		RouterGroup: router,
	}
}

// Setup Register security handlers
func (routes *SecurityRoutes) Setup() {
	riskHandler := security.NewRiskHandler()
	{
		routes.RouterGroup.GET("/get-enums", handlers.Convert(riskHandler, riskHandler.GetEnums))
		routes.RouterGroup.POST("/risk/create", handlers.Convert(riskHandler, riskHandler.Create))
		routes.RouterGroup.GET("/risk/get-list", handlers.Convert(riskHandler, riskHandler.GetList))
		routes.RouterGroup.GET("/risk/details/:id", handlers.Convert(riskHandler, riskHandler.GetByID))
		routes.RouterGroup.POST("/risk-category/create", handlers.Convert(riskHandler, riskHandler.CreateRiskCategory))
		routes.RouterGroup.POST("/risk-location/create", handlers.Convert(riskHandler, riskHandler.CreateRiskLocation))
	}
}
