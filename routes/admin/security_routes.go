package admin

import (
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
func (admin *SecurityRoutes) Setup() {
	riskReportingHandler := security.NewRiskReportingHandler()
	admin.RouterGroup.POST("/risk-reporting/create", riskReportingHandler.Create)
	admin.RouterGroup.GET("/risk-reporting/get-list", riskReportingHandler.GetList)
	admin.RouterGroup.POST("/risk-reporting/create-risk-location", riskReportingHandler.CreateRiskLocation)
}
