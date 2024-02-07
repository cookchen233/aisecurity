package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/eventlog"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go.uber.org/zap"
)

type EventLogService struct {
	Service
	entClient *dao.EventLogClient
}

func NewEventLogService(ctx context.Context) *EventLogService {
	s := &EventLogService{entClient: db.EntClient.EventLog}
	s.Ctx = ctx
	return s
}

func (s *EventLogService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.EventLog)
	c := s.entClient.Create().
		SetDeviceID(e.DeviceID).
		SetEventID(e.EventID).
		SetLogType(e.LogType).
		SetNotes(e.Notes)
	if e.ActorID != 0 {
		c.SetActorID(e.ActorID)
	}
	if e.Actor2ID != 0 {
		c.SetActor2ID(e.Actor2ID)
	}
	saved, err := c.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating EventLogService")
	}
	return structs.ConvertTo[*dao.EventLog, entities.EventLog](saved), nil
}

func (s *EventLogService) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
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

func (s *EventLogService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	// list
	list, err := s.query(fit).
		Limit(fit.GetLimit()).
		Offset(fit.GetOffset()).
		Order(dao.Desc(eventlog.FieldID)).
		All(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		v2 := new(entities.EventLog)
		ents[i] = v
		err := gconv.Struct(v, v2)
		if err != nil {
			utils.Logger.Warn("convert error", zap.Error(err))
		} else {
			ents[i] = v2
		}
	}
	return ents, nil
}

func (s *EventLogService) query(fit structs.IFilter) *dao.EventLogQuery {
	f := fit.(*filters.EventLog)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(eventlog.IDEQ(f.ID))
	}
	if f.DeviceID > 0 {
		q = q.Where(eventlog.DeviceID(f.DeviceID))
	}
	if f.EventID > 0 {
		q = q.Where(eventlog.EventID(f.EventID))
	}
	return q.Clone()
}

func (s *EventLogService) GetTotal(fit structs.IFilter) (int, error) {
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}
