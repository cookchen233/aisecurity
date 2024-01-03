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
	riskReportingHandler := security.NewRiskHandler()
	routes.RouterGroup.POST("/risk/create", handlers.Convert(riskReportingHandler, riskReportingHandler.Create))
	routes.RouterGroup.GET("/risk/get-list", handlers.Convert(riskReportingHandler, riskReportingHandler.GetList))
	routes.RouterGroup.POST("/risk/create-risk-category", handlers.Convert(riskReportingHandler, riskReportingHandler.CreateRiskCategory))
	routes.RouterGroup.POST("/risk/create-risk-location", handlers.Convert(riskReportingHandler, riskReportingHandler.CreateRiskLocation))
}
