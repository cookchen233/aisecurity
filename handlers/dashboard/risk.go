package dashboard

import (
	"aisecurity/expects"
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/structs/types"
	"aisecurity/utils"
	"aisecurity/utils/http"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
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
func (h *RiskHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewRiskService(c)
	h.Service.Ctx = c
	h.Filter = &filters.Risk{}
	h.Entity = &entities.Risk{}
	h.DashboardHandler.SetRequestContext(c, h)
}

func (h *RiskHandler) GetStatusCounts(c *gin.Context) {
	list, err := h.Service.GetStatusCounts(&filters.Risk{})
	if err != nil {
		http.Error(c, err, 2000)
		return
	}
	http.Success(c, list)
}

func (h *RiskHandler) GetList(c *gin.Context) {
	var err error
	utils.Logger.Info("called default GetList")
	if err := c.ShouldBindQuery(h.Filter); err != nil {
		http.Error(c, err, 2000)
		return
	}
	utils.Logger.Info("bound filter", zap.Any("filter", h.Filter))
	var resp = struct {
		Total        int                 `json:"total"`
		List         []structs.IEntity   `json:"list"`
		StatusCounts []*types.GroupCount `json:"status_counts"`
	}{}

	var f2 filters.Risk // copy a filter
	c.ShouldBindQuery(&f2)
	f2.MaintainStatus = 0
	f2.Status = 0
	resp.StatusCounts, err = h.Service.GetStatusCounts(&f2)
	if err != nil {
		var expectErr *expects.NotImplementedMethod
		unwrappedErr := errors.Unwrap(err)
		if unwrappedErr != nil && errors.As(unwrappedErr, &expectErr) {
			resp.StatusCounts = []*types.GroupCount{}
		} else {
			http.Error(c, err, 2000)
			return
		}
	}

	resp.Total, err = h.Service.GetTotal(h.Filter)
	if err != nil {
		http.Error(c, err, 2000)
		return
	}
	utils.Logger.Info("called total", zap.Int("total", resp.Total))
	if resp.Total == 0 {
		http.Success(c, resp)
		return
	}

	resp.List, err = h.Service.GetList(h.Filter)
	if err != nil {
		http.Error(c, err, 2000)
		return
	}
	utils.Logger.Info("called service.GetList", zap.Int("length", len(resp.List)))
	http.Success(c, resp)
}

func (h *RiskHandler) GetRiskCategoryCounts(c *gin.Context) {
	if err := c.ShouldBindQuery(h.Handler.Filter); err != nil {
		http.Error(c, err, 1000)
		return
	}
	list, err := h.Service.GetRiskCategoryCounts(h.Handler.Filter)
	if err != nil {
		http.Error(c, err, 2000)
		return
	}
	http.Success(c, list)
}
