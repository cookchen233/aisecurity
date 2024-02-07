package dashboard

import (
	"aisecurity/enums"
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

type EventHandler struct {
	handlers.DashboardHandler
	Service *services.EventService
}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}
func (h *EventHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *EventHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *EventHandler) GetEntity(c *gin.Context) structs.IEntity {
	return h.Entity
}
func (h *EventHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewEventService(c)
	h.Service.Ctx = c
	h.Filter = &filters.Event{}
	h.Entity = &entities.Event{}
	h.DashboardHandler.SetRequestContext(c, h)
}

func (h *EventHandler) GetStatusCounts(c *gin.Context) {
	var f = filters.Event{
		EventStatus: enums.EventStatus(0),
	}
	list, err := h.Service.GetStatusCounts(&f)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, list)
}

func (h *EventHandler) GetList(c *gin.Context) {
	var err error
	utils.Logger.Info("called default GetList")
	if err := c.ShouldBindQuery(h.Handler.Filter); err != nil {
		http.Error(c, err, 900)
		return
	}
	utils.Logger.Info("bound filter", zap.Any("filter", h.Handler.Filter))
	var resp = struct {
		Total        int                  `json:"total"`
		List         []structs.IEntity    `json:"list"`
		StatusCounts []*types.StatusCount `json:"status_counts"`
	}{}

	var f2 filters.Event // copy a filter
	c.ShouldBindQuery(&f2)
	f2.EventStatus = 0
	resp.StatusCounts, err = h.Service.GetStatusCounts(&f2)
	if err != nil {
		var expectErr *expects.NotImplementedMethod
		unwrappedErr := errors.Unwrap(err)
		if unwrappedErr != nil && errors.As(unwrappedErr, &expectErr) {
			resp.StatusCounts = []*types.StatusCount{}
		} else {
			http.Error(c, err, 1000)
			return
		}
	}

	resp.Total, err = h.Handler.Service.GetTotal(h.Handler.Filter)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	utils.Logger.Info("called total", zap.Int("total", resp.Total))
	if resp.Total == 0 {
		http.Success(c, resp)
		return
	}

	resp.List, err = h.Handler.Service.GetList(h.Handler.Filter)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	utils.Logger.Info("called service.GetList", zap.Int("length", len(resp.List)))
	http.Success(c, resp)
}
