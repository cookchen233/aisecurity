package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	handlers.DashboardHandler
	Service *services.VideoService
}

func NewVideoHandler() *VideoHandler {
	return &VideoHandler{}
}
func (h *VideoHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *VideoHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *VideoHandler) GetEntity(c *gin.Context) structs.IEntity {
	return h.Entity
}
func (h *VideoHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewVideoService(c)
	h.Service.Ctx = c
	h.Filter = &filters.Video{}
	h.Entity = &entities.Video{}
	h.DashboardHandler.SetRequestContext(c, h)
}
