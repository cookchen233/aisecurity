package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"context"
)

// AdminHandler Toycar handles
type AdminHandler struct {
	handlers.Handler
	Service *services.AdminService
}

func NewAdminHandler() *AdminHandler {
	h := &AdminHandler{}
	h.Service = services.NewAdminService()
	h.Handler.Service = h.Service
	return h
}
func (h *AdminHandler) ResetRequest(ctx context.Context) {
	h.Filter = &filters.Admin{}
	h.Handler.Filter = h.Filter
	h.Entity = &entities.Admin{}
	h.Handler.Entity = h.Entity
}
