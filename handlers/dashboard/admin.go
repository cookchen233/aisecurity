package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"github.com/gin-gonic/gin"
)

// AdminHandler Toycar handles
type AdminHandler struct {
	handlers.DashboardHandler
	Service *services.AdminService
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}
func (h *AdminHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *AdminHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *AdminHandler) GetEntity(c *gin.Context) structs.IEntity {
	return h.Entity
}
func (h *AdminHandler) SetRequestContext(c *gin.Context, h2 handlers.IHandler) {
	h.Service = services.NewAdminService()
	h.Service.Ctx = c
	h.Filter = &filters.Admin{}
	h.Entity = &entities.Admin{}
	h.DashboardHandler.SetRequestContext(c, h)
}
