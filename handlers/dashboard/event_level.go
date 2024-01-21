package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"github.com/gin-gonic/gin"
)

type IEventLevelHandler struct {
	handlers.DashboardHandler
	Service *services.EventLevelService
}

func NewIEventLevelHandler() *IEventLevelHandler {
	return &IEventLevelHandler{}
}
func (h *IEventLevelHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *IEventLevelHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *IEventLevelHandler) GetEntity(c *gin.Context) structs.IEntity {
	return h.Entity
}
func (h *IEventLevelHandler) SetRequestContext(c *gin.Context, h2 handlers.IHandler) {
	h.Service = services.NewEventLevelService(c)
	h.Service.Ctx = c
	h.Filter = &filters.EventLevel{}
	h.Entity = &entities.EventLevel{}
	h.DashboardHandler.SetRequestContext(c, h)
}
