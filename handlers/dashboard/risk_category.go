package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"context"
)

type RiskCategoryHandler struct {
	handlers.Handler
	Service *services.RiskCategoryService
}

func NewRiskCategoryHandler() *RiskCategoryHandler {
	h := &RiskCategoryHandler{}
	h.Service = services.NewRiskCategoryService()
	h.Handler.Service = h.Service
	return h
}
func (h *RiskCategoryHandler) ResetRequest(ctx context.Context) {
	h.Filter = &filters.RiskCategory{}
	h.Handler.Filter = h.Filter
	h.Entity = &entities.RiskCategory{}
	h.Handler.Entity = h.Entity
}
