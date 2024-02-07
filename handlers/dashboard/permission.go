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

// PermissionHandler Toycar handles
type PermissionHandler struct {
	handlers.DashboardHandler
	Service *services.PermissionService
}

func NewPermissionHandler() *PermissionHandler {
	return &PermissionHandler{}
}
func (h *PermissionHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *PermissionHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *PermissionHandler) GetEntity(c *gin.Context) structs.IEntity {
	return h.Entity
}
func (h *PermissionHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewPermissionService(c)
	h.Filter = &filters.Permission{}
	h.Entity = &entities.Permission{}
	h.DashboardHandler.SetRequestContext(c, h)
}
func (h *PermissionHandler) GetAccessGroups(c *gin.Context) {
	modules, err := h.Service.GetAccessGroups()
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, modules)
}
