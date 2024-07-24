package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/fixing"
	"aisecurity/enums"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"go.uber.org/zap"
	"os"
	"time"
)

type FixingService struct {
	Service
	entClient *dao.FixingClient
}

func NewFixingService(ctx context.Context) *FixingService {
	s := &FixingService{entClient: db.EntClient.Fixing}
	s.Ctx = ctx
	return s
}

func (s *FixingService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Fixing)
	c := s.entClient.Create().
		SetFixer(e.Fixer).
		SetEvent(e.Event).
		SetAssignNotes(e.AssignNotes).
		SetDevice(e.Device)
	saved, err := c.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating Fixing")
	}
	return structs.ConvertTo[*dao.Fixing, entities.Fixing](saved), nil
}

func (s *FixingService) Update(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Fixing)
	u := s.entClient.UpdateOneID(e.ID).
		SetCompleteNotes(e.CompleteNotes)
	if e.FixerID != 0 {
		u.SetFixerID(e.FixerID)
	}
	if !e.FixTime.IsZero() {
		u.SetFixTime(e.FixTime)
	}
	if !e.CompleteTime.IsZero() {
		u.SetCompleteTime(e.CompleteTime)
	}
	saved, err := u.Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed updating Fixing")
	}
	return structs.ConvertTo[*dao.Fixing, entities.Fixing](saved), nil
}

func (s *FixingService) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
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

func (s *FixingService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	// list
	list, err := s.query(fit).
		WithCreator(func(q *dao.AdminQuery) {
			q.WithEmployee(func(q2 *dao.EmployeeQuery) {
				q2.WithDepartment().WithOccupation()
			})
		}).
		Limit(fit.GetLimit()).
		Offset(fit.GetOffset()).
		Order(dao.Desc(fixing.FieldID)).
		All(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		v2 := new(entities.Fixing)
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

func (s *FixingService) query(fit structs.IFilter) *dao.FixingQuery {
	f := fit.(*filters.Fixing)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(fixing.IDEQ(f.ID))
	}
	if f.EventID > 0 {
		q = q.Where(fixing.EventID(f.EventID))
	}
	return q.Clone()
}

func (s *FixingService) GetTotal(fit structs.IFilter) (int, error) {
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *FixingService) Flow(ent structs.IEntity) (structs.IEntity, error) {
	var eventStatus enums.EventStatus
	var saved structs.IEntity
	var err error
	e := ent.(*entities.Fixing)
	if e.EventLog.LogType == enums.ELT2 {
		saved, err = s.Create(ent)
		if err != nil {
			return nil, err
		}
		if e.Fixer.WechatOpenid == "" {
			utils.Logger.Error("the fixer has not bond Wechat", zap.Error(err), zap.String("username", e.Fixer.Username))
		} else {
			go func() {
				admin, err := s.GetCurrentAdmin()
				if err != nil {
					utils.Logger.Error("get current admin failed", zap.Error(err))
					return
				}
				adm := admin.(*entities.Admin)
				msgData := make(map[string]*utils.TemplateDataItem)
				msgData["thing2"] = &utils.TemplateDataItem{Value: adm.RealName}
				msgData["thing3"] = &utils.TemplateDataItem{Value: e.Fixer.RealName}
				msgData["thing7"] = &utils.TemplateDataItem{Value: e.Device.Edges.DeviceInstallation[0].Edges.Area.Name}
				msgData["thing10"] = &utils.TemplateDataItem{Value: e.Event.EventType.Label()}
				_, err = utils.SendTemplateMsg(msgData, e.Fixer.WechatOpenid, "fu66XPPCgHpc9yKYrIb3kpwtPywI9k7Wficmb9RvwN8", fmt.Sprintf("%s/dashboard/ipc-event/%d", os.Getenv("DASHBOARD_SITE"), e.Event.ID))
				if err != nil {
					utils.Logger.Error("sending template msg failed", zap.Error(err), zap.String("username", e.Fixer.Username))
				} else {
					utils.Logger.Error("sending template successfully", zap.Error(err), zap.String("username", e.Fixer.Username))
				}
			}()
		}
	} else {
		switch e.EventLog.LogType {
		case enums.ELT5:
			eventStatus = enums.ES1
		case enums.ELT4:
			eventStatus = enums.ES3
			e.CompleteTime = time.Now()
		case enums.ELT3:
			e.FixTime = time.Now()
			eventStatus = enums.ES2
		default:
			eventStatus = enums.ES2
		}
		saved, err = s.Update(ent)
		if err != nil {
			return nil, utils.ErrorWithStack(err)
		}
	}

	// event log
	_, err = NewEventLogService(s.Ctx).Create(structs.ConvertTo[*dao.EventLog, entities.EventLog](e.EventLog))
	if err != nil {
		utils.Logger.Error("failed creating EventLog", zap.Error(err))
	}
	// update event status
	e.Event.EventStatus = eventStatus
	_, err = NewEventService(s.Ctx).Update(structs.ConvertTo[*dao.Event, entities.Event](e.Event))
	if err != nil {
		utils.Logger.Error("failed creating EventLog", zap.Error(err))
	}

	if e.EventLog.LogType == enums.ELT4 {
		if e.Edges.Creator.WechatOpenid == "" {
			utils.Logger.Error("the creator has not bond Wechat", zap.Error(err), zap.String("username", e.Edges.Creator.Username))
		} else {
			go func() {
				msgData := make(map[string]*utils.TemplateDataItem)
				msgData["thing9"] = &utils.TemplateDataItem{Value: e.Device.Name}
				msgData["thing2"] = &utils.TemplateDataItem{Value: e.Device.Edges.DeviceInstallation[0].Edges.Area.Name}
				msgData["thing5"] = &utils.TemplateDataItem{Value: e.Event.EventType.Label()}
				msgData["thing6"] = &utils.TemplateDataItem{Value: e.EventLog.Notes}
				msgData["thing4"] = &utils.TemplateDataItem{Value: e.Fixer.RealName}
				_, err = utils.SendTemplateMsg(msgData, e.Fixer.WechatOpenid, "nxke99tmcJomvFBfTZXHxKmc33vuUX9R-3_DztiLlt8", fmt.Sprintf("%s/dashboard/ipc-event/%d", os.Getenv("DASHBOARD_SITE"), e.Event.ID))
				if err != nil {
					utils.Logger.Error("sending template msg failed", zap.Error(err), zap.String("username", e.Fixer.Username))
				} else {
					utils.Logger.Error("sending template successfully", zap.Error(err), zap.String("username", e.Fixer.Username))
				}
			}()
		}
	}
	return saved, nil
}
