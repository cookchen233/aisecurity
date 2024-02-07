package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"github.com/gin-gonic/gin"
)

type SweepHandler struct {
	handlers.DashboardHandler
	Service *services.SweepService
}

func NewSweepHandler() *SweepHandler {
	return &SweepHandler{}
}
func (h *SweepHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *SweepHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *SweepHandler) GetEntity(c *gin.Context) structs.IEntity { return h.Entity }
func (h *SweepHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewSweepService(c)
	h.Filter = &filters.Sweep{}
	h.Entity = &entities.Sweep{}
	h.DashboardHandler.SetRequestContext(c, h)
}
