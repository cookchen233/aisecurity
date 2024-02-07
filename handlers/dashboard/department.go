package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	handlers.DashboardHandler
	Service *services.DepartmentService
}

func NewDepartmentHandler() *DepartmentHandler {
	return &DepartmentHandler{}
}
func (h *DepartmentHandler) GetService(c *gin.Context) services.IService { return h.Service }
func (h *DepartmentHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *DepartmentHandler) GetEntity(c *gin.Context) structs.IEntity {
	return h.Entity
}
func (h *DepartmentHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewDepartmentService(c)
	h.Filter = &filters.Department{}
	h.Entity = &entities.Department{}
	h.DashboardHandler.SetRequestContext(c, h)
}
