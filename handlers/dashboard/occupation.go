package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"github.com/gin-gonic/gin"
)

type OccupationHandler struct {
	handlers.DashboardHandler
	Service *services.OccupationService
}

func NewOccupationHandler() *OccupationHandler {
	return &OccupationHandler{}
}
func (h *OccupationHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *OccupationHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *OccupationHandler) GetEntity(c *gin.Context) structs.IEntity { return h.Entity }
func (h *OccupationHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewOccupationService(c)
	h.Filter = &filters.Occupation{}
	h.Entity = &entities.Occupation{}
	h.DashboardHandler.SetRequestContext(c, h)
}
