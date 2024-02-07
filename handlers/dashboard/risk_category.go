package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"github.com/gin-gonic/gin"
)

type RiskCategoryHandler struct {
	handlers.DashboardHandler
	Service *services.RiskCategoryService
}

func NewRiskCategoryHandler() *RiskCategoryHandler {
	return &RiskCategoryHandler{}
}
func (h *RiskCategoryHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *RiskCategoryHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *RiskCategoryHandler) GetEntity(c *gin.Context) structs.IEntity { return h.Entity }
func (h *RiskCategoryHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewRiskCategoryService()
	h.Service.Ctx = c
	h.Filter = &filters.RiskCategory{}
	h.Entity = &entities.RiskCategory{}
	h.DashboardHandler.SetRequestContext(c, h)
}
