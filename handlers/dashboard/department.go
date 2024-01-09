package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"context"
)

type DepartmentHandler struct {
	handlers.Handler
	Service *services.DepartmentService
}

func NewDepartmentHandler() *DepartmentHandler {
	h := &DepartmentHandler{}
	h.Service = services.NewDepartmentService()
	h.Handler.Service = h.Service
	return h
}
func (h *DepartmentHandler) ResetRequest(ctx context.Context) {
	h.Filter = &filters.Department{}
	h.Handler.Filter = h.Filter
	h.Entity = &entities.Department{}
	h.Handler.Entity = h.Entity
}
