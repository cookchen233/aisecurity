package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
)

// AdminHandler Toycar handles
type AdminHandler struct {
	handlers.Handler
	Service *services.AdminService
}

func NewAdminHandler() *AdminHandler {
	h := &AdminHandler{
		Service: services.NewAdminService(),
	}
	h.Handler.Service = h.Service
	h.Filter = &filters.Admin{}
	h.Handler.Filter = h.Filter
	h.Entity = &entities.Admin{}
	h.Handler.Entity = h.Entity
	return h
}
