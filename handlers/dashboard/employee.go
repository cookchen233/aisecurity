package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
)

type EmployeeHandler struct {
	handlers.Handler
	Service *services.EmployeeService
}

func NewEmployeeHandler() *EmployeeHandler {
	h := &EmployeeHandler{
		Service: services.NewEmployeeService(),
	}
	h.Handler.Service = h.Service
	h.Filter = &filters.Employee{}
	h.Handler.Filter = h.Filter
	h.Entity = &entities.Employee{}
	h.Handler.Entity = h.Entity
	return h
}
