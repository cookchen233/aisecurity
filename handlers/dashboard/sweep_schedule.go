package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"github.com/gin-gonic/gin"
)

type SweepScheduleHandler struct {
	handlers.DashboardHandler
	Service *services.SweepScheduleService
}

func NewSweepScheduleHandler() *SweepScheduleHandler {
	return &SweepScheduleHandler{}
}
func (h *SweepScheduleHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *SweepScheduleHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *SweepScheduleHandler) GetEntity(c *gin.Context) structs.IEntity { return h.Entity }
func (h *SweepScheduleHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewSweepScheduleService(c)
	h.Filter = &filters.SweepSchedule{}
	h.Entity = &entities.SweepSchedule{}
	h.DashboardHandler.SetRequestContext(c, h)
}
