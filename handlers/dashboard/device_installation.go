package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"github.com/gin-gonic/gin"
)

type DeviceInstallationHandler struct {
	handlers.DashboardHandler
	Service *services.DeviceInstallationService
}

func NewIDeviceInstallationHandler() *DeviceInstallationHandler {
	return &DeviceInstallationHandler{}
}
func (h *DeviceInstallationHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *DeviceInstallationHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *DeviceInstallationHandler) GetEntity(c *gin.Context) structs.IEntity {
	return h.Entity
}
func (h *DeviceInstallationHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewDeviceInstallationService(c)
	h.Service.Ctx = c
	h.Filter = &filters.DeviceInstallation{}
	h.Entity = &entities.DeviceInstallation{}
	h.DashboardHandler.SetRequestContext(c, h)
}
