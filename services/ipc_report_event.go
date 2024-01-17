package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/ipcreportevent"
	"aisecurity/enums"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/structs/types"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"go.uber.org/zap"
)

type IPCReportEventService struct {
	Service
	entClient *dao.IPCReportEventClient
}

func NewIPCReportEventService() *IPCReportEventService {
	return &IPCReportEventService{
		entClient: db.EntClient.IPCReportEvent,
	}
}

func (s *IPCReportEventService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.IPCReportEvent)
	c := s.entClient.Create().
		SetDeviceID(e.DeviceID).
		SetEventID(e.EventID).
		SetEventTime(e.EventTime).
		SetEventType(e.EventType).
		SetDescription(e.Description).
		SetRawData(e.RawData)
	if e.Images != nil {
		c.SetImages(e.Images)
	}
	if e.Videos != nil {
		c.SetVideos(e.Videos)
	}
	save, err := c.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating IPCReportEvent")
	}
	return save, nil
}

func (s *IPCReportEventService) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
	fit.SetPage(1)
	fit.SetLimit(1)
	list, err := s.GetList(fit)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	if len(list) == 0 {
		return nil, utils.ErrorWithStack(expects.NewDataNotFound())
	}
	return list[0], nil
}

func (s *IPCReportEventService) query(fit structs.IFilter) *dao.IPCReportEventQuery {
	f := fit.(*filters.IPCReportEvent)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(ipcreportevent.IDEQ(f.ID))
	}
	if f.Keywords != "" {
		q = q.Where(ipcreportevent.DescriptionContainsFold(f.Keywords))
	}
	eventTypes := utils.FilterZerosFromArray(f.EventTypes)
	if len(eventTypes) > 0 {
		q = q.Where(ipcreportevent.EventTypeIn(eventTypes...))
	}
	if f.EventStatus != 0 {
		q = q.Where(ipcreportevent.EventStatusIn(f.EventStatus))
	}
	return q.Clone()
}

func (s *IPCReportEventService) GetTotal(fit structs.IFilter) (int, error) {
	// total
	fmt.Println("risk get total", fit)
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *IPCReportEventService) GetIREEventStatusCounts(fit structs.IFilter) ([]*types.IREEventStatusCounts, error) {
	// status counts
	var counts []*types.IREEventStatusCounts
	err := s.query(fit).GroupBy(ipcreportevent.FieldEventStatus).
		Aggregate(dao.Count()).
		Scan(s.Ctx, &counts)
	if err != nil {
		return counts, utils.ErrorWithStack(err)
	}
	for _, status := range enums.IPCReportEventStatus(0).GetAll() {
		if status == enums.IRESUnknown {
			continue
		}
		var ex bool
		for _, count := range counts {
			if count.EventStatus == status {
				count.Label = status.Label()
				ex = true
				break
			}
		}
		if !ex {
			counts = append(counts, &types.IREEventStatusCounts{
				EventStatus: status,
				Count:       0,
				Label:       status.Label(),
			})
		}
	}
	return counts, nil
}

func (s *IPCReportEventService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	// list
	list, err := s.query(fit).
		Limit(fit.GetLimit()).
		Offset(fit.GetOffset()).
		All(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		v2 := new(entities.IPCReportEvent)
		ents[i] = v
		err := gconv.Struct(v, v2)
		if err != nil {
			utils.Logger.Warn("convert error", zap.Error(err))
		} else {
			v2.EventStatusLabel = v2.EventStatus.Label()
			ents[i] = v2
		}
	}
	return ents, nil
}

func (s *IPCReportEventService) GetByVideoName(name string) (*dao.IPCReportEvent, error) {
	result, err := s.entClient.QueryContext(s.Ctx, "SELECT id,videos FROM ipc_report_events WHERE videos->0->>'name' = $1", name)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	for result.Next() {
		var row dao.IPCReportEvent
		err = result.Scan(&row.ID, &row.Videos)
		if err != nil {
			return nil, utils.ErrorWithStack(err)
		}
		return &row, nil
	}
	return nil, &dao.NotFoundError{}
}

func (s *IPCReportEventService) UpdateVideos(ent structs.IEntity) (*dao.IPCReportEvent, error) {
	e := ent.(*entities.IPCReportEvent)
	save, err := s.entClient.UpdateOneID(e.ID).
		SetVideos(e.Videos).
		Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed updating IPCReportEvent")
	}
	return save, nil
}
