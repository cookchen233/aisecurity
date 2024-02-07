package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/device"
	"aisecurity/ent/dao/event"
	"aisecurity/ent/dao/eventlevel"
	"aisecurity/ent/dao/fixing"
	"aisecurity/ent/dao/video"
	"aisecurity/enums"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/structs/types"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
	stdsql "database/sql"
	"encoding/json"
	"entgo.io/ent/dialect/sql"
	"go.uber.org/zap"
)

type EventService struct {
	Service
	entClient         *dao.EventClient
	eventLevelService *EventLevelService
}

func NewEventService(ctx context.Context) *EventService {
	s := &EventService{
		entClient:         db.EntClient.Event,
		eventLevelService: NewEventLevelService(ctx),
	}
	s.Ctx = ctx
	return s
}

func (s *EventService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Event)
	c := s.entClient.Create().
		SetDeviceID(e.DeviceID).
		SetDataID(e.DataID).
		SetEventTime(e.EventTime).
		SetEventType(e.EventType).
		SetDescription(e.Description).
		SetRawData(e.RawData)
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
		return nil, utils.ErrorWrap(err, "failed creating Event")
	}

	return structs.ConvertTo[*dao.Event, entities.Event](saved), nil
}

func (s *EventService) Update(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Event)
	u := s.entClient.UpdateOneID(e.ID).
		SetEventStatus(e.EventStatus)
	saved, err := u.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed updating Event")
	}

	return structs.ConvertTo[*dao.Event, entities.Event](saved), nil
}

func (s *EventService) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
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

func (s *EventService) query(fit structs.IFilter) *dao.EventQuery {
	f := fit.(*filters.Event)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(event.IDEQ(f.ID))
	}
	if f.Keywords != "" {
		q = q.Where(event.DescriptionContainsFold(f.Keywords))
	}
	eventTypes := utils.FilterZerosFromArray(f.EventTypes)
	if len(eventTypes) > 0 {
		q = q.Where(event.EventTypeIn(eventTypes...))
	}
	eventLevelIDs := utils.FilterZerosFromArray(f.EventLevelIDs)
	if len(eventLevelIDs) > 0 {
		eventTypes, err := db.EntClient.EventLevel.Query().Where(eventlevel.IDIn(eventLevelIDs...)).Select(eventlevel.FieldEventTypes).Ints(s.Ctx)
		if err != nil {
			utils.Logger.Error("failed getting event levels", zap.Error(err))
		}
		var es []enums.EventType
		for _, v := range eventTypes {
			es = append(es, enums.EventType(v))
		}
		q = q.Where(event.EventTypeIn(es...))
	}
	fixerIDs := utils.FilterZerosFromArray(f.FixerIDs)
	if len(fixerIDs) > 0 {
		q = q.Where(event.HasFixingWith(fixing.FixerIDIn(fixerIDs...)))
	}
	if !f.EventTimeRange.Start.IsZero() {
		q = q.Where(event.EventTimeGTE(f.EventTimeRange.Start))
	}
	if !f.EventTimeRange.End.IsZero() {
		q = q.Where(event.EventTimeLTE(f.EventTimeRange.End))
	}

	if !f.FixTimeRange.Start.IsZero() {
		q = q.Where(event.HasFixingWith(fixing.FixTimeGTE(f.FixTimeRange.Start)))
	}
	if !f.FixTimeRange.End.IsZero() {
		q = q.Where(event.HasFixingWith(fixing.FixTimeLTE(f.FixTimeRange.End)))
	}

	deviceIDs := utils.FilterZerosFromArray(f.DeviceIDs)
	if len(deviceIDs) > 0 {
		q = q.Where(event.DeviceIDIn(deviceIDs...))
	}
	if f.EventStatus != 0 {
		q = q.Where(event.EventStatusIn(f.EventStatus))
	}
	return q.Clone()
}

func (s *EventService) GetTotal(fit structs.IFilter) (int, error) {
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *EventService) GetStatusCounts(fit structs.IFilter) ([]*types.StatusCount, error) {
	// status counts
	var queryCounts []struct {
		EventStatus enums.EventStatus `json:"event_status"`
		Count       int
	}
	err := s.query(fit).GroupBy(event.FieldEventStatus).
		Aggregate(dao.Count()).
		Scan(s.Ctx, &queryCounts)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	var statusCounts []*types.StatusCount
	for _, s := range enums.EventStatus(0).GetAll() {
		var c int
		for _, q := range queryCounts {
			if q.EventStatus == s {
				c = q.Count
				break
			}
			if s == enums.ESUnknown {
				c += q.Count
			}
		}
		statusCounts = append(statusCounts, &types.StatusCount{
			Value: int(s),
			Label: s.Label(),
			Count: c,
		})
	}
	return statusCounts, nil
}

func (s *EventService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	// list
	list, err := s.query(fit).
		WithCreator(func(q *dao.AdminQuery) {
			q.WithEmployee(func(q2 *dao.EmployeeQuery) {
				q2.WithDepartment().WithOccupation()
			})
		}).
		WithFixing(func(q *dao.FixingQuery) {
			q.WithFixer(func(q2 *dao.AdminQuery) {
				q2.WithEmployee(func(q3 *dao.EmployeeQuery) {
					q3.WithDepartment().WithOccupation()
				})
			}).
				WithCreator(func(q2 *dao.AdminQuery) {
					q2.WithEmployee(func(q3 *dao.EmployeeQuery) {
						q3.WithDepartment().WithOccupation()
					})
				})
		}).
		WithDevice(func(q *dao.DeviceQuery) {
			q.WithDeviceInstallation(func(query *dao.DeviceInstallationQuery) {
				query.WithArea()
			}).Order(device.ByID(sql.OrderDesc()))
		}).
		WithVideo().
		Limit(fit.GetLimit()).
		Offset(fit.GetOffset()).
		Order(dao.Desc(event.FieldID)).
		All(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		v.EventTypeLabel = v.EventType.Label()
		v.EventStatusLabel = v.EventStatus.Label()
		if len(v.Edges.Device.Edges.DeviceInstallation) > 0 {
			d := v.Edges.Device.Edges.DeviceInstallation[0]
			v.Location = d.Location
			v.LocationWithAliasName = d.Location
			if d.AreaName != "" {
				v.LocationWithAliasName = d.Location + " (" + d.AliasName + ")"
			}
		}
		ents[i] = structs.ConvertTo[*dao.Event, entities.Event](v)
	}
	return ents, nil
}

func (s *EventService) GetListByImageName(name string) ([]structs.IEntity, error) {
	result, err := s.entClient.QueryContext(s.Ctx, "SELECT id, images FROM events WHERE images->0->>'name' = $1", name)
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
		rows = append(rows, &entities.Event{
			Event: dao.Event{
				ID:     id,
				Images: images,
			},
		})
	}

	return rows, nil
}

func (s *EventService) GetByVideoName(name string) (structs.IEntity, error) {
	first, err := s.entClient.Query().
		Where(event.HasVideoWith(video.NameEQ(name))).
		WithVideo().
		First(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	return structs.ConvertTo[*dao.Event, entities.Event](first), nil
}
