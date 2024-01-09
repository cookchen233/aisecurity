package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
)

type RiskHandler struct {
	handlers.Handler
	Service *services.RiskService
}

func NewRiskHandler() *RiskHandler {
	h := &RiskHandler{
		Service: services.NewRiskService(),
	}
	h.Handler.Service = h.Service
	h.Filter = &filters.Risk{}
	h.Handler.Filter = h.Filter
	h.Entity = &entities.Risk{}
	h.Handler.Entity = h.Entity
	return h
}
