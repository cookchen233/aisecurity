package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"github.com/gin-gonic/gin"
)

type RiskLocationHandler struct {
	handlers.DashboardHandler
	Service *services.RiskLocationService
}

func NewRiskLocationHandler() *RiskLocationHandler {
	return &RiskLocationHandler{}
}
func (h *RiskLocationHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *RiskLocationHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *RiskLocationHandler) GetEntity(c *gin.Context) structs.IEntity { return h.Entity }
func (h *RiskLocationHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewRiskLocationService()
	h.Service.Ctx = c
	h.Filter = &filters.RiskLocation{}
	h.Entity = &entities.RiskLocation{}
	h.DashboardHandler.SetRequestContext(c, h)
}
