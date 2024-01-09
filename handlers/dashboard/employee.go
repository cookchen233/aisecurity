package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"context"
)

type EmployeeHandler struct {
	handlers.Handler
	Service *services.EmployeeService
}

func NewEmployeeHandler() *EmployeeHandler {
	h := &EmployeeHandler{}
	h.Service = services.NewEmployeeService()
	h.Handler.Service = h.Service
	return h
}
func (h *EmployeeHandler) ResetRequest(ctx context.Context) {
	h.Filter = &filters.Employee{}
	h.Handler.Filter = h.Filter
	h.Entity = &entities.Employee{}
	h.Handler.Entity = h.Entity
}
