package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/risk"
	"aisecurity/enums"
	"aisecurity/expects"
	"aisecurity/properties/maintain_status"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/structs/types"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"github.com/gogf/gf/v2/util/gconv"
	"go.uber.org/zap"
)

type RiskService struct {
	Service
	entClient *dao.RiskClient
}

func NewRiskService() *RiskService {
	return &RiskService{
		entClient: db.EntClient.Risk,
	}
}

func (s *RiskService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Risk)
	c := s.entClient.Create().
		SetTitle(e.Title).
		SetContent(e.Content)
	if e.Images != nil {
		c.SetImages(e.Images)
	}
	if e.Maintainer != nil {
		c.SetMaintainer(e.Maintainer)
	}
	save, err := c.
		SetMeasures(e.Measures).
		SetRiskLocation(e.RiskLocation).
		SetRiskCategory(e.RiskCategory).
		SetMaintainStatus(maintain_status.Ready).
		SetDueTime(e.DueTime).
		Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating RiskLocation")
	}
	return save, nil
}

func (s *RiskService) Update(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Risk)
	u := s.entClient.UpdateOneID(e.ID)
	if e.Maintainer != nil {
		u = u.SetMaintainer(e.Maintainer)
	}
	save, err := u.SetTitle(e.Title).
		SetContent(e.Content).
		SetImages(e.Images).
		SetRiskLocation(e.RiskLocation).
		SetRiskCategory(e.RiskCategory).
		SetMeasures(e.Measures).
		//SetMaintainStatus(e.MaintainStatus).
		SetDueTime(e.DueTime).
		Save(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed updating RiskLocation")
	}
	return save, nil
}

func (s *RiskService) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
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

func (s *RiskService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	// list
	list, err := s.query(fit).
		WithCreator(func(aQuery *dao.AdminQuery) {
			aQuery.WithEmployee(func(eQuery *dao.EmployeeQuery) {
				eQuery.WithDepartment().WithOccupation()
			})
		}).
		WithRiskLocation().
		WithRiskCategory().
		WithMaintainer(func(q *dao.AdminQuery) {
			q.WithEmployee(func(q2 *dao.EmployeeQuery) {
				q2.WithDepartment().WithOccupation()
			})
		}).
		Limit(fit.GetLimit()).
		Offset(fit.GetOffset()).
		Order(dao.Desc(risk.FieldID)).
		All(s.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		v2 := new(entities.Risk)
		ents[i] = v
		err := gconv.Struct(v, v2)
		if err != nil {
			utils.Logger.Warn("convert error", zap.Error(err))
		} else {
			v2.MaintainStatusLabel = v2.MaintainStatus.Label()
			ents[i] = v2
		}
	}
	return ents, nil
}

func (s *RiskService) query(fit structs.IFilter) *dao.RiskQuery {
	utils.Logger.Info("risk query", zap.Any("fit", fit))
	f := fit.(*filters.Risk)
	q := db.EntClient.Risk.Query()
	if f.ID != 0 {
		q = q.Where(risk.IDEQ(f.ID))
	}
	if f.Title != "" {
		q = q.Where(risk.TitleContainsFold(f.Title))
	}
	if f.CreatorID != 0 {
		q = q.Where(risk.CreatorID(f.CreatorID))
	}
	maintainerIDs := utils.FilterZerosFromArray(f.MaintainerIDs)
	if len(maintainerIDs) > 0 {
		q = q.Where(risk.MaintainerIDIn(maintainerIDs...))
	}
	categoryIDs := utils.FilterZerosFromArray(f.RiskCategoryIDs)
	if len(categoryIDs) > 0 {
		q = q.Where(risk.RiskCategoryIDIn(categoryIDs...))
	}
	locationIDs := utils.FilterZerosFromArray(f.RiskLocationIDs)
	if len(locationIDs) > 0 {
		q = q.Where(risk.RiskLocationIDIn(locationIDs...))
	}
	if f.MaintainStatus != 0 {
		q = q.Where(risk.MaintainStatusIn(f.MaintainStatus))
	}
	if f.Status != 0 {
		q = q.Where(risk.MaintainStatusIn(f.Status))
	}
	if !f.CreateTimeRange.Start.IsZero() {
		q = q.Where(risk.CreateTimeGTE(f.CreateTimeRange.Start))
	}
	if !f.CreateTimeRange.End.IsZero() {
		q = q.Where(risk.CreateTimeLTE(f.CreateTimeRange.End))
	}
	if !f.DueTimeRange.Start.IsZero() {
		q = q.Where(risk.CreateTimeGTE(f.DueTimeRange.Start))
	}
	if !f.DueTimeRange.End.IsZero() {
		q = q.Where(risk.CreateTimeLTE(f.DueTimeRange.End))
	}
	return q.Clone()
}

func (s *RiskService) GetTotal(fit structs.IFilter) (int, error) {
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *RiskService) GetStatusCounts(fit structs.IFilter) ([]*types.StatusCount, error) {
	// status counts
	var queryCounts []struct {
		MaintainStatus enums.MaintainStatus `json:"maintain_status"`
		Count          int
	}
	err := s.query(fit).GroupBy(risk.FieldMaintainStatus).
		Aggregate(dao.Count()).
		Scan(s.Ctx, &queryCounts)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	var statusCounts []*types.StatusCount
	for _, s := range enums.MaintainStatus(0).GetAll() {
		var c int
		for _, q := range queryCounts {
			if q.MaintainStatus == s {
				c = q.Count
				break
			}
			if s == enums.MSUnknown {
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
