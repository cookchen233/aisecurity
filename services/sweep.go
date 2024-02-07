package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/sweep"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
)

type SweepService struct {
	Service
	entClient *dao.SweepClient
}

func NewSweepService(ctx context.Context) *SweepService {
	s := &SweepService{entClient: db.EntClient.Sweep}
	s.Ctx = ctx
	return s
}

var ()

func (s *SweepService) Create(ent structs.IEntity) (structs.IEntity, error) {
	tx, err := db.EntClient.Tx(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(err))
	}
	e := ent.(*entities.Sweep)
	c := tx.Sweep.Create().
		SetName(e.Name).
		SetRiskLocation(e.RiskLocation).
		SetRiskCategory(e.RiskCategory).
		SetSweepJobs(e.SweepJobs)
	saved, err := c.Save(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "failed creating Sweep"))
	}
	err = tx.Commit()
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "commit failed"))
	}
	return structs.ConvertTo[*dao.Sweep, entities.Sweep](saved), nil
}

func (s *SweepService) Update(ent structs.IEntity) (structs.IEntity, error) {
	tx, err := db.EntClient.Tx(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(err))
	}
	e := ent.(*entities.Sweep)
	u := tx.Sweep.UpdateOneID(e.ID).
		SetName(e.Name).
		SetRiskLocation(e.RiskLocation).
		SetRiskCategory(e.RiskCategory).
		SetSweepJobs(e.SweepJobs)
	saved, err := u.Save(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "failed updating Sweep"))
	}
	err = tx.Commit()
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "commit failed"))
	}
	return structs.ConvertTo[*dao.Sweep, entities.Sweep](saved), nil
}

func (s *SweepService) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
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

func (s *SweepService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	list, err := s.query(fit).
		WithCreator(func(q *dao.AdminQuery) {
			q.WithEmployee(func(q2 *dao.EmployeeQuery) {
				q2.WithDepartment().WithOccupation()
			})
		}).
		WithRiskCategory().
		WithRiskLocation().
		Order(dao.Desc(sweep.FieldID)).
		All(s.Ctx)
	if err != nil {
		return nil, err
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		ents[i] = structs.ConvertTo[*dao.Sweep, entities.Sweep](v)
	}
	return ents, nil
}

func (s *SweepService) query(fit structs.IFilter) *dao.SweepQuery {
	f := fit.(*filters.Sweep)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(sweep.IDEQ(f.ID))
	}
	if !f.CreateTimeRange.Start.IsZero() {
		q = q.Where(sweep.CreateTimeGTE(f.CreateTimeRange.Start))
	}
	if !f.CreateTimeRange.End.IsZero() {
		q = q.Where(sweep.CreateTimeLTE(f.CreateTimeRange.End))
	}
	if f.Name != "" {
		q = q.Where(sweep.NameEQ(f.Name))
	}
	categoryIDs := utils.FilterZerosFromArray(f.RiskCategoryIDs)
	if len(categoryIDs) > 0 {
		q = q.Where(sweep.RiskCategoryIDIn(categoryIDs...))
	}
	locationIDs := utils.FilterZerosFromArray(f.RiskLocationIDs)
	if len(locationIDs) > 0 {
		q = q.Where(sweep.RiskLocationIDIn(locationIDs...))
	}
	return q.Clone()
}

func (s *SweepService) GetTotal(fit structs.IFilter) (int, error) {
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}
