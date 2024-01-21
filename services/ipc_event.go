package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/employee"
	"aisecurity/ent/dao/ipcevent"
	"aisecurity/ent/dao/video"
	"aisecurity/enums"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/structs/types"
	"aisecurity/utils"
	"aisecurity/utils/db"
	stdsql "database/sql"
	"encoding/json"
	"go.uber.org/zap"
)

type IPCEventService struct {
	Service
	entClient *dao.IPCEventClient
}

func NewIPCEventService() *IPCEventService {
	return &IPCEventService{
		entClient: db.EntClient.IPCEvent,
	}
}

func (s *IPCEventService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.IPCEvent)
	c := s.entClient.Create().
		SetDeviceID(e.DeviceID).
		SetEventID(e.EventID).
		SetEventTime(e.EventTime).
		SetEventType(e.EventType).
		SetDescription(e.Description).
		SetRawData(e.RawData).AddFixerIDs(1)
	if e.Images != nil {
		c.SetImages(e.Images)
	}
	if e.LabeledImages != nil {
		c.SetLabeledImages(e.LabeledImages)
	}
	if e.VideoID != 0 {
		c.SetVideoID(e.VideoID)
	}
	saved, err := c.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating IPCEvent")
	}
	return structs.ConvertTo[*dao.IPCEvent, entities.IPCEvent](saved), nil
}

func (s *IPCEventService) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
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

func (s *IPCEventService) query(fit structs.IFilter) *dao.IPCEventQuery {
	f := fit.(*filters.IPCEvent)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(ipcevent.IDEQ(f.ID))
	}
	if f.Keywords != "" {
		q = q.Where(ipcevent.DescriptionContainsFold(f.Keywords))
	}
	eventTypes := utils.FilterZerosFromArray(f.EventTypes)
	if len(eventTypes) > 0 {
		q = q.Where(ipcevent.EventTypeIn(eventTypes...))
	}
	fixerIDs := utils.FilterZerosFromArray(f.FixerIDs)
	if len(fixerIDs) > 0 {
		q = q.Where(ipcevent.HasFixersWith(employee.IDIn(fixerIDs...)))
	}
	if f.EventStatus != 0 {
		q = q.Where(ipcevent.EventStatusIn(f.EventStatus))
	}
	return q.Clone()
}

func (s *IPCEventService) GetTotal(fit structs.IFilter) (int, error) {
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *IPCEventService) GetIEEventStatusCounts(fit structs.IFilter) ([]*types.IEEventStatusCounts, error) {
	// status counts
	var counts []*types.IEEventStatusCounts
	err := s.query(fit).GroupBy(ipcevent.FieldEventStatus).
		Aggregate(dao.Count()).
		Scan(s.Ctx, &counts)
	if err != nil {
		return counts, utils.ErrorWithStack(err)
	}
	for _, status := range enums.EventStatus(0).GetAll() {
		if status == enums.ESUnknown {
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
			counts = append(counts, &types.IEEventStatusCounts{
				EventStatus: status,
				Count:       0,
				Label:       status.Label(),
			})
		}
	}
	return counts, nil
}

func (s *IPCEventService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	// list
	list, err := s.query(fit).
		WithFixers(func(q *dao.EmployeeQuery) {
			q.WithDepartment().WithAdmin()
		}).
		Limit(fit.GetLimit()).
		Offset(fit.GetOffset()).
		All(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	ents := make([]structs.IEntity, len(list))
	for i, _ := range list {
		ents[i] = structs.ConvertTo[*dao.IPCEvent, entities.IPCEvent](list[i])
	}
	return ents, nil
}

func (s *IPCEventService) GetListByImageName(name string) ([]structs.IEntity, error) {
	result, err := s.entClient.QueryContext(s.Ctx, "SELECT id, images FROM ipc_events WHERE images->0->>'name' = $1", name)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	defer func(result *stdsql.Rows) {
		err := result.Close()
		if err != nil {
			utils.Logger.Error("failed closing rows", zap.Error(err))
		}
	}(result)

	var rows []structs.IEntity
	for result.Next() {
		var id int
		var imagesData []byte
		err = result.Scan(&id, &imagesData)
		if err != nil {
			return nil, utils.ErrorWithStack(err)
		}

		var images []*types.UploadedImage
		if len(imagesData) > 0 {
			err = json.Unmarshal(imagesData, &images)
			if err != nil {
				return nil, utils.ErrorWithStack(err)
			}
		}
		rows = append(rows, &entities.IPCEvent{
			IPCEvent: dao.IPCEvent{
				ID:     id,
				Images: images,
			},
		})
	}

	return rows, nil
}

func (s *IPCEventService) GetByVideoName(name string) (structs.IEntity, error) {
	first, err := s.entClient.Query().
		Where(ipcevent.HasVideoWith(video.NameEQ(name))).
		WithVideo().
		First(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	return structs.ConvertTo[*dao.IPCEvent, entities.IPCEvent](first), nil
}
