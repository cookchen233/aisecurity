package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils/http"
	"github.com/gin-gonic/gin"
)

// AdminRoleHandler Toycar handles
type AdminRoleHandler struct {
	handlers.DashboardHandler
	Service *services.AdminRoleService
}

func NewAdminRoleHandler() *AdminRoleHandler {
	return &AdminRoleHandler{}
}
func (h *AdminRoleHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *AdminRoleHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *AdminRoleHandler) GetEntity(c *gin.Context) structs.IEntity {
	return h.Entity
}
func (h *AdminRoleHandler) SetRequestContext(c *gin.Context, h2 handlers.IHandler) {
	h.Service = services.NewAdminRoleService()
	h.Service.Ctx = c
	h.Filter = &filters.AdminRole{}
	h.Entity = &entities.AdminRole{}
	h.DashboardHandler.SetRequestContext(c, h)
}
func (h *AdminRoleHandler) GetModules(c *gin.Context) {
	modules, err := h.Service.GetModules()
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, modules)
}
