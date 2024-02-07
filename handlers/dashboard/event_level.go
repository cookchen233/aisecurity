package dashboard

import (
	"aisecurity/enums"
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/http"
	"github.com/gin-gonic/gin"
	"slices"
)

type EventLevelHandler struct {
	handlers.DashboardHandler
	Service *services.EventLevelService
}

func NewEventLevelHandler() *EventLevelHandler {
	return &EventLevelHandler{}
}
func (h *EventLevelHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *EventLevelHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *EventLevelHandler) GetEntity(c *gin.Context) structs.IEntity {
	return h.Entity
}
func (h *EventLevelHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewEventLevelService(c)
	h.Service.Ctx = c
	h.Filter = &filters.EventLevel{}
	h.Entity = &entities.EventLevel{}
	h.DashboardHandler.SetRequestContext(c, h)
}

func (h *EventLevelHandler) GetListForEachEventType(c *gin.Context) {
	list, err := h.Service.GetList(&filters.EventLevel{StandardFilter: structs.StandardFilter{Limit: 10}})
	if err != nil {
		http.Error(c, utils.ErrorWithStack(err), 1000)
		return
	}
	type level struct {
		EventType  enums.EventType      `json:"event_type"`
		EventLevel *entities.EventLevel `json:"event_level"`
	}
	types := enums.EventType(0).GetAll()
	var levels = make([]level, len(types))
	for i, t := range types {
		for _, v := range list {
			v2 := v.(*entities.EventLevel)
			if slices.Contains(v2.EventTypes, t) {
				levels[i] = level{
					EventType:  t,
					EventLevel: v2,
				}
			}
		}
	}
	http.Success(c, levels)
}
