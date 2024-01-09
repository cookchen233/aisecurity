package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
)

type RiskCategoryHandler struct {
	handlers.Handler
	Service *services.RiskCategoryService
}

func NewRiskCategoryHandler() *RiskCategoryHandler {
	h := &RiskCategoryHandler{
		Service: services.NewRiskCategoryService(),
	}
	h.Handler.Service = h.Service
	h.Filter = &filters.RiskCategory{}
	h.Handler.Filter = h.Filter
	h.Entity = &entities.RiskCategory{}
	h.Handler.Entity = h.Entity
	return h
}
