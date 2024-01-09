package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
)

type RiskLocationHandler struct {
	handlers.Handler
	Service *services.RiskLocationService
}

func NewRiskLocationHandler() *RiskLocationHandler {
	h := &RiskLocationHandler{
		Service: services.NewRiskLocationService(),
	}
	h.Handler.Service = h.Service
	h.Filter = &filters.RiskLocation{}
	h.Handler.Filter = h.Filter
	h.Entity = &entities.RiskLocation{}
	h.Handler.Entity = h.Entity
	return h
}
