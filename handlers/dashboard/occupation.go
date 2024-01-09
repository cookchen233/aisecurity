package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
)

type OccupationHandler struct {
	handlers.Handler
	Service *services.OccupationService
}

func NewOccupationHandler() *OccupationHandler {
	h := &OccupationHandler{
		Service: services.NewOccupationService(),
	}
	h.Handler.Service = h.Service
	h.Filter = &filters.Occupation{}
	h.Handler.Filter = h.Filter
	h.Entity = &entities.Occupation{}
	h.Handler.Entity = h.Entity
	return h
}
