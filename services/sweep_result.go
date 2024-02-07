package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/sweepresult"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
	"time"
)

type SweepResultService struct {
	Service
	entClient *dao.SweepResultClient
}

func NewSweepResultService(ctx context.Context) *SweepResultService {
	s := &SweepResultService{entClient: db.EntClient.SweepResult}
	s.Ctx = ctx
	return s
}

var ()

func (s *SweepResultService) Create(ent structs.IEntity) (structs.IEntity, error) {
	tx, err := db.EntClient.Tx(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(err))
	}
	e := ent.(*entities.SweepResult)
	c := tx.SweepResult.Create().
		SetSweepSchedule(e.SweepSchedule).
		SetSweep(e.SweepSchedule.Edges.Sweep).
		SetSweepJobs(e.SweepSchedule.Edges.Sweep.SweepJobs)
	saved, err := c.Save(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "failed creating SweepResult"))
	}
	err = tx.Commit()
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "commit failed"))
	}
	return structs.ConvertTo[*dao.SweepResult, entities.SweepResult](saved), nil
}

func (s *SweepResultService) Update(ent structs.IEntity) (structs.IEntity, error) {
	tx, err := db.EntClient.Tx(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(err))
	}
	e := ent.(*entities.SweepResult)
	if e.CheckInTime.IsZero() {
		e.CheckInTime = time.Now()
	}
	u := tx.SweepResult.UpdateOneID(e.ID)
	for _, v := range e.SweepJobs {
		if v.Result > 0 {
			v.UpdaterID = s.GetCurrentAdminID()
			v.UpdateTime = time.Now()
		}
	}
	u = u.SetSweepJobs(e.SweepJobs)
	saved, err := u.Save(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "failed updating SweepResult"))
	}
	err = tx.Commit()
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "commit failed"))
	}
	return structs.ConvertTo[*dao.SweepResult, entities.SweepResult](saved), nil
}

func (s *SweepResultService) CheckIn(ent structs.IEntity) (structs.IEntity, error) {
	tx, err := db.EntClient.Tx(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(err))
	}
	e := ent.(*entities.SweepResult)
	u := tx.SweepResult.UpdateOneID(e.ID).
		SetCheckInTime(time.Now())
	saved, err := u.Save(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "failed updating SweepResult"))
	}
	err = tx.Commit()
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "commit failed"))
	}
	return structs.ConvertTo[*dao.SweepResult, entities.SweepResult](saved), nil
}

func (s *SweepResultService) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
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

func (s *SweepResultService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	list, err := s.query(fit).
		WithCreator(func(q *dao.AdminQuery) {
			q.WithEmployee(func(q2 *dao.EmployeeQuery) {
				q2.WithDepartment().WithOccupation()
			})
		}).
		WithSweep(func(q *dao.SweepQuery) {
			q.WithRiskLocation().WithRiskCategory()
		}).
		WithSweepSchedule(func(q *dao.SweepScheduleQuery) {
			q.WithWorkers()
		}).
		Order(dao.Desc(sweepresult.FieldID)).
		All(s.Ctx)
	if err != nil {
		return nil, err
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		ents[i] = structs.ConvertTo[*dao.SweepResult, entities.SweepResult](v)
	}
	return ents, nil
}

func (s *SweepResultService) query(fit structs.IFilter) *dao.SweepResultQuery {
	f := fit.(*filters.SweepResult)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(sweepresult.IDEQ(f.ID))
	}
	return q.Clone()
}

func (s *SweepResultService) GetTotal(fit structs.IFilter) (int, error) {
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *SweepResultService) StartWork(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.SweepResult)
	result, err := s.GetDetails(&filters.SweepResult{StandardFilter: structs.StandardFilter{ID: e.ID}})
	if err != nil {
		if !expects.IsDataNotFound(err) {
			return nil, utils.ErrorWithStack(err)
		}
		e.CheckInTime = time.Now()
		return s.Create(e)
	}
	r := result.(*entities.SweepResult)
	if r.CheckInTime.IsZero() {
		return result, nil
	}
	return s.Create(ent)
}
