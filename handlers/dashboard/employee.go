package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	handlers.DashboardHandler
	Service *services.EmployeeService
}

func NewEmployeeHandler() *EmployeeHandler {
	return &EmployeeHandler{}
}
func (h *EmployeeHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *EmployeeHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *EmployeeHandler) GetEntity(c *gin.Context) structs.IEntity { return h.Entity }
func (h *EmployeeHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewEmployeeService(c)
	h.Service.Ctx = c
	h.Filter = &filters.Employee{}
	h.Entity = &entities.Employee{}
	h.DashboardHandler.SetRequestContext(c, h)
}
