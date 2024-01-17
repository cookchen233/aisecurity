package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/risk"
	"aisecurity/expects"
	"aisecurity/properties/maintain_status"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/structs/types"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"fmt"
	"github.com/gin-gonic/gin"
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
		SetReporterID(max(1, s.Ctx.(*gin.Context).GetInt("admin_id"))).
		SetTitle(e.Title).
		SetContent(e.Content)
	if e.Images != nil {
		c.SetImages(e.Images)
	}
	save, err := c.SetRiskCategoryID(e.RiskCategoryID).
		SetRiskLocationID(e.RiskLocationID).
		SetMaintainerID(e.MaintainerID).
		SetMeasures(e.Measures).
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
	save, err := s.entClient.UpdateOneID(e.ID).
		SetTitle(e.Title).
		SetContent(e.Content).
		SetImages(e.Images).
		SetRiskCategoryID(e.RiskCategoryID).
		SetRiskLocationID(e.RiskLocationID).
		SetMaintainerID(e.MaintainerID).
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
		WithRiskLocation().
		WithRiskCategory().
		WithMaintainer(func(q *dao.EmployeeQuery) {
			q.WithDepartment().WithOccupations().WithAdmin(func(q *dao.AdminQuery) {
				q.WithAdminRoles()
			})
		}).
		WithReporter(func(q *dao.EmployeeQuery) {
			q.WithDepartment().WithOccupations().WithAdmin(func(q *dao.AdminQuery) {
				q.WithAdminRoles()
			})
		}).
		Limit(fit.GetLimit()).
		Offset(fit.GetOffset()).
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
	if f.CreatedBy != 0 {
		q = q.Where(risk.CreatedBy(f.CreatedBy))
	}
	maintainerIDs := utils.FilterZerosFromArray(f.MaintainerIDs)
	if len(maintainerIDs) > 0 {
		q = q.Where(risk.MaintainerIDIn(maintainerIDs...))
	}
	reporterIDs := utils.FilterZerosFromArray(f.ReporterIDs)
	if len(reporterIDs) > 0 {
		q = q.Where(risk.ReporterIDIn(reporterIDs...))
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
	return q.Clone()
}

func (s *RiskService) GetTotal(fit structs.IFilter) (int, error) {
	// total
	fmt.Println("risk get total", fit)
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *RiskService) GetMaintainStatusCounts(fit structs.IFilter) ([]*types.MaintainStatusCounts, error) {
	// status counts
	var counts []*types.MaintainStatusCounts
	err := s.query(fit).GroupBy(risk.FieldMaintainStatus).
		Aggregate(dao.Count()).
		Scan(s.Ctx, &counts)
	if err != nil {
		return counts, utils.ErrorWithStack(err)
	}
	for _, status := range maintain_status.Unknown.GetAll() {
		if status == maintain_status.Unknown {
			continue
		}
		var ex bool
		for _, count := range counts {
			if count.MaintainStatus == status {
				count.Label = status.Label()
				ex = true
				break
			}
		}
		if !ex {
			counts = append(counts, &types.MaintainStatusCounts{
				MaintainStatus: status,
				Count:          0,
				Label:          status.Label(),
			})
		}
	}
	return counts, nil
}
