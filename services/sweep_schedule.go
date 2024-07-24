package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/sweepschedule"
	"aisecurity/enums"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/structs/types"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"time"
)

type SweepScheduleService struct {
	Service
	entClient *dao.SweepScheduleClient
}

func NewSweepScheduleService(ctx context.Context) *SweepScheduleService {
	s := &SweepScheduleService{entClient: db.EntClient.SweepSchedule}
	s.Ctx = ctx
	return s
}

var ()

func (s *SweepScheduleService) Create(ent structs.IEntity) (structs.IEntity, error) {
	tx, err := db.EntClient.Tx(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(err))
	}
	e := ent.(*entities.SweepSchedule)
	c := tx.SweepSchedule.Create().
		SetName(e.Name).
		SetSweep(e.Sweep).
		SetRepeat(e.Repeat).
		SetActionTime(e.ActionTime).
		SetRemind(e.Remind).
		AddWorkers(e.Workers...)
	saved, err := c.Save(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "failed creating SweepSchedule"))
	}
	err = tx.Commit()
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "commit failed"))
	}
	return structs.ConvertTo[*dao.SweepSchedule, entities.SweepSchedule](saved), nil
}

func (s *SweepScheduleService) Update(ent structs.IEntity) (structs.IEntity, error) {
	tx, err := db.EntClient.Tx(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(err))
	}
	e := ent.(*entities.SweepSchedule)
	u := tx.SweepSchedule.UpdateOneID(e.ID).ClearWorkers().
		SetName(e.Name).
		SetSweep(e.Sweep).
		SetRepeat(e.Repeat).
		SetActionTime(e.ActionTime).
		SetRemind(e.Remind).
		AddWorkers(e.Workers...)
	saved, err := u.Save(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "failed updating SweepSchedule"))
	}
	err = tx.Commit()
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "commit failed"))
	}
	return structs.ConvertTo[*dao.SweepSchedule, entities.SweepSchedule](saved), nil
}

func (s *SweepScheduleService) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
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

func (s *SweepScheduleService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	list, err := s.query(fit).
		WithCreator(func(q *dao.AdminQuery) {
			q.WithEmployee(func(q2 *dao.EmployeeQuery) {
				q2.WithDepartment().WithOccupation()
			})
		}).
		WithWorkers().
		WithSweep(func(q *dao.SweepQuery) {
			q.WithRiskLocation()
		}).
		WithSweepResult().
		Order(dao.Desc(sweepschedule.FieldID)).
		All(s.Ctx)
	if err != nil {
		return nil, err
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		ents[i] = structs.ConvertTo[*dao.SweepSchedule, entities.SweepSchedule](v)
	}
	return ents, nil
}

func (s *SweepScheduleService) query(fit structs.IFilter) *dao.SweepScheduleQuery {
	f := fit.(*filters.SweepSchedule)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(sweepschedule.IDEQ(f.ID))
	}
	if !f.CreateTimeRange.Start.IsZero() {
		q = q.Where(sweepschedule.CreateTimeGTE(f.CreateTimeRange.Start))
	}
	if !f.CreateTimeRange.End.IsZero() {
		q = q.Where(sweepschedule.CreateTimeLTE(f.CreateTimeRange.End))
	}
	if f.Name != "" {
		q = q.Where(sweepschedule.NameEQ(f.Name))
	}
	if !f.ActionTimeRange.Start.IsZero() {
		q = q.Where(sweepschedule.ActionTimeGTE(f.ActionTimeRange.Start))
	}
	if !f.ActionTimeRange.End.IsZero() {
		q = q.Where(sweepschedule.ActionTimeLTE(f.ActionTimeRange.End))
	}
	if f.SweepID > 0 {
		q = q.Where(sweepschedule.SweepID(f.SweepID))
	}
	sweepIDs := utils.FilterZerosFromArray(f.SweepIDs)
	if len(sweepIDs) > 0 {
		q = q.Where(sweepschedule.SweepIDIn(sweepIDs...))
	}
	if f.WorkerID > 0 {
		q = q.Where(sweepschedule.HasWorkersWith(admin.ID(f.WorkerID)))
	}
	workerIDs := utils.FilterZerosFromArray(f.WorkerIDs)
	if len(workerIDs) > 0 {
		q = q.Where(sweepschedule.HasWorkersWith(admin.IDIn(workerIDs...)))
	}
	return q.Clone()
}

func (s *SweepScheduleService) GetTotal(fit structs.IFilter) (int, error) {
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *SweepScheduleService) SetCrontab() {
	c := cron.New()
	var taskRegistry = map[string]types.ScheduledTask{}
	_, err := c.AddFunc("*/1 * * * *", func() {
		var fit = &filters.SweepSchedule{
			EnabledStatus: enums.ENS1,
		}
		list, err := s.GetList(fit)
		if err != nil {
			utils.Logger.Error("failed getting sweep schedule list", zap.Error(err))
			return
		}
		for _, v := range list {
			schedule := v.(*entities.SweepSchedule)

			// Determine the cron schedule string based on the RepeatType
			if schedule.Repeat.ExpiresType == enums.EXT2 && time.Now().After(schedule.Repeat.EndTime) {
				continue
			}
			var cronSchedule string
			var remindTime = schedule.ActionTime.Add(-time.Duration(schedule.Remind.AdvancedSeconds) * time.Second)
			switch schedule.Repeat.RepeatType {
			case enums.RT1:
				// "不重复" - This should be handled as a one-time event, perhaps not with cron
				continue
			case enums.RT2:
				// "每日"
				cronSchedule = fmt.Sprintf("%d %d * * *", remindTime.Minute(), remindTime.Hour())
			case enums.RT3:
				// "每周"
				cronSchedule = fmt.Sprintf("%d %d * * %d", remindTime.Minute(), remindTime.Hour(), remindTime.Weekday())
			case enums.RT4:
				// "每月"
				cronSchedule = fmt.Sprintf("%d %d %d * *", remindTime.Minute(), remindTime.Hour(), remindTime.Day())
			case enums.RT5:
				// "每隔?日" - Repeat every 'Interval' days
				cronSchedule = fmt.Sprintf("%d %d */%d * *", remindTime.Minute(), remindTime.Hour(), schedule.Repeat.Interval)
			case enums.RT6:
				// "每隔?周" - Since cron does not support week intervals directly, we handle it as a daily check within the job.
				// This requires logic inside the job to check if the current week matches the interval.
				// This placeholder does not implement the logic but sets up a daily check.
				cronSchedule = fmt.Sprintf("%d %d * * *", remindTime.Minute(), remindTime.Hour())
			case enums.RT7:
				// "每隔?月" - Repeat on the same day every 'Interval' months. Cron does not support month intervals, so this is a placeholder.
				// Actual implementation may require storing the last execution date and comparing it.
				cronSchedule = fmt.Sprintf("%d %d %d */%d *", remindTime.Minute(), remindTime.Hour(), remindTime.Day(), schedule.Repeat.Interval)
			default:
				continue
			}

			c2 := cron.New()

			utils.Logger.Debug("=======k", zap.String("dd", cronSchedule))
			if _, exists := taskRegistry[cronSchedule]; exists {
				continue
			}

			_, err := c2.AddFunc(cronSchedule, func() {
				//m := gomail.NewMessage()
				//m.SetHeader("From", m.FormatAddress(hook.User, "Golang App Error")) //这种方式可以添加别名，即“XX官方”
				////说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
				////m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用m.SetHeader("From",mailConn["user"])
				////m.SetHeader("From", mailConn["user"])
				//reg1 := regexp.MustCompile(`(.*?)<(.*?)>`)
				//var to []string
				//for _, v := range mailTo {
				//    res := reg1.FindAllStringSubmatch(v, -1)
				//    if len(res) > 0 {
				//        to = append(to, m.FormatAddress(res[0][2], res[0][1]))
				//    }
				//}
				//m.SetHeader("To", to...)
				//m.SetHeader("Subject", subject)
				//m.SetBody("text/html", body)
				//port, _ := strconv.Atoi(hook.Port)
				//d := gomail.NewDialer(hook.Host, port, hook.User, hook.Pass)
				//err := d.DialAndSend(m)
				//return err
				utils.Logger.Debug("cron job executed")
				msgData := make(map[string]*utils.TemplateDataItem)
				msgData["thing3"] = &utils.TemplateDataItem{Value: schedule.Name}
				msgData["thing2"] = &utils.TemplateDataItem{Value: schedule.Edges.Sweep.Edges.RiskLocation.Name}
				for _, worker := range schedule.Edges.Workers {
					wk := structs.ConvertTo[*dao.Admin, entities.Admin](worker)
					if wk.WechatOpenid == "" {
						utils.Logger.Error("the worker has not bond Wechat", zap.Error(err), zap.String("username", wk.Username))
						continue
					}
					_, err = utils.SendTemplateMsg(msgData, wk.WechatOpenid, "s_VhFCUCnRlAI9jSLm-AkSCdJRb-kb03BM62xpWAKyg", "https://baidu.com")
					if err != nil {
						utils.Logger.Error("sending template msg failed", zap.Error(err), zap.String("username", wk.Username))
					} else {
						utils.Logger.Error("sending template successfully", zap.Error(err), zap.String("username", wk.Username))
					}
				}
			})
			if err != nil {
				utils.Logger.Error("failed adding cron task", zap.Error(err))
			} else {
				taskRegistry[cronSchedule] = types.ScheduledTask{
					SweepScheduleID: schedule.ID,
					CronSchedule:    cronSchedule,
				}
				c2.Start()
			}
		}
	})
	if err != nil {
		utils.Logger.Error("failed executing AddFunc", zap.Error(err))
	} else {
		c.Start()
	}
}
