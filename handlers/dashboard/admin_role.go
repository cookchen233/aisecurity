package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils/http"
	"github.com/gin-gonic/gin"
)

// AdminRoleHandler Toycar handles
type AdminRoleHandler struct {
	handlers.Handler
	Service *services.AdminRoleService
}

func NewAdminRoleHandler() *AdminRoleHandler {
	h := &AdminRoleHandler{
		Service: services.NewAdminRoleService(),
	}
	h.Handler.Service = h.Service
	h.Filter = &filters.AdminRole{}
	h.Handler.Filter = h.Filter
	h.Entity = &entities.AdminRole{}
	h.Handler.Entity = h.Entity
	return h
}

func (handler *AdminRoleHandler) GetModules(c *gin.Context) {
	modules, err := handler.Service.GetModules()
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, modules)
}
