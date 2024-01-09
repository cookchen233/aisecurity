package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"context"
)

type RiskLocationHandler struct {
	handlers.Handler
	Service *services.RiskLocationService
}

func NewRiskLocationHandler() *RiskLocationHandler {
	h := &RiskLocationHandler{}
	h.Service = services.NewRiskLocationService()
	h.Handler.Service = h.Service
	return h
}
func (h *RiskLocationHandler) ResetRequest(ctx context.Context) {
	h.Filter = &filters.RiskLocation{}
	h.Handler.Filter = h.Filter
	h.Entity = &entities.RiskLocation{}
	h.Handler.Entity = h.Entity
}
