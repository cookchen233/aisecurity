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

type FixingHandler struct {
	handlers.DashboardHandler
	Service *services.FixingService
}

func NewFixingHandler() *FixingHandler {
	return &FixingHandler{}
}
func (h *FixingHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *FixingHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *FixingHandler) GetEntity(c *gin.Context) structs.IEntity {
	return h.Entity
}
func (h *FixingHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewFixingService(c)
	h.Service.Ctx = c
	h.Filter = &filters.Fixing{}
	h.Entity = &entities.Fixing{}
	h.DashboardHandler.SetRequestContext(c, h)
}

func (h *FixingHandler) Flow(c *gin.Context) {
	if err := c.ShouldBindJSON(h.Entity); err != nil {
		http.Error(c, utils.ErrorWithStack(err), 1000)
		return
	}
	if err := h.Validate(h.Entity); err != nil {
		http.Error(c, utils.ErrorWithStack(err), 1000)
		return
	}
	saved, err := h.Service.Flow(h.Entity)
	if err != nil {
		http.Error(c, utils.ErrorWithStack(err), 2000)
		return
	}
	http.Success(c, saved)
}
