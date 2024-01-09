package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/structs/types"
	"aisecurity/utils/http"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RiskHandler struct {
	handlers.Handler
	Service *services.RiskService
}

func NewRiskHandler() *RiskHandler {
	h := &RiskHandler{}
	h.Service = services.NewRiskService()
	h.Handler.Service = h.Service
	return h
}
func (h *RiskHandler) ResetRequest(ctx context.Context) {
	h.Filter = &filters.Risk{}
	h.Handler.Filter = h.Filter
	h.Entity = &entities.Risk{}
	h.Handler.Entity = h.Entity
}

func (h *RiskHandler) GetList(c *gin.Context) {
	var err error
	if err = c.ShouldBindQuery(h.Filter); err != nil {
		http.Error(c, err, 900)
		return
	}
	var resp = struct {
		Total                int                          `json:"total"`
		List                 []structs.IEntity            `json:"list"`
		MaintainStatusCounts []types.MaintainStatusCounts `json:"maintain_status_counts"`
	}{}

	var f2 *filters.Risk // copy a filter
	c.ShouldBindQuery(&f2)
	f2.MaintainStatus = 0
	resp.MaintainStatusCounts, err = h.Service.GetMaintainStatusCounts(f2)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}

	fmt.Println("total handler filter", h.Filter)
	resp.Total, err = h.Service.GetTotal(h.Filter)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	if resp.Total == 0 {
		http.Success(c, resp)
		return
	}

	resp.List, err = h.Service.GetList(h.Filter)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, resp)
}
