package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/http"
	"github.com/gin-gonic/gin"
)

type SweepResultHandler struct {
	handlers.DashboardHandler
	Service *services.SweepResultService
}

func NewSweepResultHandler() *SweepResultHandler {
	return &SweepResultHandler{}
}
func (h *SweepResultHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *SweepResultHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *SweepResultHandler) GetEntity(c *gin.Context) structs.IEntity { return h.Entity }
func (h *SweepResultHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewSweepResultService(c)
	h.Filter = &filters.SweepResult{}
	h.Entity = &entities.SweepResult{}
	h.DashboardHandler.SetRequestContext(c, h)
}

func (h *SweepResultHandler) CheckIn(c *gin.Context) {
	if err := c.ShouldBindJSON(h.Entity); err != nil {
		http.Error(c, utils.ErrorWithStack(err), 1000)
		return
	}
	if err := h.Validate(h.Entity); err != nil {
		http.Error(c, utils.ErrorWithStack(err), 1000)
		return
	}
	saved, err := h.Service.CheckIn(h.Entity)
	if err != nil {
		http.Error(c, utils.ErrorWithStack(err), 2000)
		return
	}
	http.Success(c, saved)
}

func (h *SweepResultHandler) StartWork(c *gin.Context) {
	if err := c.ShouldBindJSON(h.Entity); err != nil {
		http.Error(c, utils.ErrorWithStack(err), 1000)
		return
	}
	if err := h.Validate(h.Entity); err != nil {
		http.Error(c, utils.ErrorWithStack(err), 1000)
		return
	}
	saved, err := h.Service.StartWork(h.Entity)
	if err != nil {
		http.Error(c, utils.ErrorWithStack(err), 2000)
		return
	}
	http.Success(c, saved)
}
