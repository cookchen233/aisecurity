package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"context"
)

type OccupationHandler struct {
	handlers.Handler
	Service *services.OccupationService
}

func NewOccupationHandler() *OccupationHandler {
	h := &OccupationHandler{}
	h.Service = services.NewOccupationService()
	h.Handler.Service = h.Service
	return h
}
func (h *OccupationHandler) ResetRequest(ctx context.Context) {
	h.Filter = &filters.Occupation{}
	h.Handler.Filter = h.Filter
	h.Entity = &entities.Occupation{}
	h.Handler.Entity = h.Entity
}
