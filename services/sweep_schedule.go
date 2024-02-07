package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/sweepschedule"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
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
	u := tx.SweepSchedule.UpdateOneID(e.ID).
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
		WithSweep().
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
