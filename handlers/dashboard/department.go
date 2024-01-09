package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
)

type DepartmentHandler struct {
	handlers.Handler
	Service *services.DepartmentService
}

func NewDepartmentHandler() *DepartmentHandler {
	h := &DepartmentHandler{
		Service: services.NewDepartmentService(),
	}
	h.Handler.Service = h.Service
	h.Filter = &filters.Department{}
	h.Handler.Filter = h.Filter
	h.Entity = &entities.Department{}
	h.Handler.Entity = h.Entity
	return h
}
