package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/structs/types"
	"aisecurity/utils/http"
	"github.com/gin-gonic/gin"
)

type RiskHandler struct {
	handlers.DashboardHandler
	Service *services.RiskService
}

func NewRiskHandler() *RiskHandler {
	return &RiskHandler{}
}
func (h *RiskHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *RiskHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *RiskHandler) GetEntity(c *gin.Context) structs.IEntity {
	return h.Entity
}
func (h *RiskHandler) SetRequestContext(c *gin.Context, h2 handlers.IHandler) {
	h.Service = services.NewRiskService()
	h.Service.Ctx = c
	h.Filter = &filters.Risk{}
	h.Entity = &entities.Risk{}
	h.DashboardHandler.SetRequestContext(c, h)
}

func (h *RiskHandler) GetList(c *gin.Context) {
	var err error
	if err = c.ShouldBindQuery(h.Filter); err != nil {
		http.Error(c, err, 900)
		return
	}
	var resp = struct {
		Total                int                           `json:"total"`
		List                 []structs.IEntity             `json:"list"`
		MaintainStatusCounts []*types.MaintainStatusCounts `json:"maintain_status_counts"`
	}{}

	var f2 *filters.Risk // copy a filter
	c.ShouldBindQuery(&f2)
	f2.MaintainStatus = 0
	resp.MaintainStatusCounts, err = h.Service.GetMaintainStatusCounts(f2)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}

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
