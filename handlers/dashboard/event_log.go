package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"github.com/gin-gonic/gin"
)

type EventLogHandler struct {
	handlers.DashboardHandler
	Service *services.EventLogService
}

func NewEventLogHandler() *EventLogHandler {
	return &EventLogHandler{}
}
func (h *EventLogHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *EventLogHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *EventLogHandler) GetEntity(c *gin.Context) structs.IEntity {
	return h.Entity
}
func (h *EventLogHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewEventLogService(c)
	h.Service.Ctx = c
	h.Filter = &filters.EventLog{}
	h.Entity = &entities.EventLog{}
	h.DashboardHandler.SetRequestContext(c, h)
}
