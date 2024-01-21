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

type IPCEventHandler struct {
	handlers.DashboardHandler
	Service *services.IPCEventService
}

func NewIPCEventHandler() *IPCEventHandler {
	return &IPCEventHandler{}
}
func (h *IPCEventHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *IPCEventHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *IPCEventHandler) GetEntity(c *gin.Context) structs.IEntity {
	return h.Entity
}
func (h *IPCEventHandler) SetRequestContext(c *gin.Context, h2 handlers.IHandler) {
	h.Service = services.NewIPCEventService()
	h.Service.Ctx = c
	h.Filter = &filters.IPCEvent{}
	h.Entity = &entities.IPCEvent{}
	h.DashboardHandler.SetRequestContext(c, h)
}

func (h *IPCEventHandler) GetList(c *gin.Context) {
	var err error
	if err = c.ShouldBindQuery(h.Filter); err != nil {
		http.Error(c, err, 900)
		return
	}
	var resp = struct {
		Total             int                          `json:"total"`
		List              []structs.IEntity            `json:"list"`
		EventStatusCounts []*types.IEEventStatusCounts `json:"event_status_counts"`
	}{}

	var f2 *filters.IPCEvent // copy a filter
	c.ShouldBindQuery(&f2)
	f2.EventStatus = 0
	resp.EventStatusCounts, err = h.Service.GetIEEventStatusCounts(f2)
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
	eventLevelService := services.NewEventLevelService(c)
	installationService := services.NewDeviceInstallation(c)
	for _, v1 := range resp.List {
		v := v1.(*entities.IPCEvent)
		v.EventStatusLabel = v.EventStatus.Label()
		v.EventTypeLabel = v.EventType.Label()
		eventLevels, err := eventLevelService.GetList(&filters.EventLevel{
			EventType: v.EventType,
		})
		if err == nil && len(eventLevels) > 0 {
			v.EventLevel = *eventLevels[0].(*entities.EventLevel)
		}
		ins, err := installationService.GetByDeviceID(v.DeviceID)
		if err == nil {
			v.DeviceInstallation = *ins.(*entities.DeviceInstallation)
		}

	}
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, resp)
}
