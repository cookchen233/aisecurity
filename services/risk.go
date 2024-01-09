package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/risk"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"github.com/gin-gonic/gin"
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

func (service *RiskService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Risk)
	save, err := db.EntClient.Risk.Create().
		SetReporterID(max(1, service.Ctx.(*gin.Context).GetInt("admin_id"))).
		SetTitle(e.Title).
		SetContent(e.Content).
		SetImages(e.Images).
		SetRiskCategoryID(e.RiskCategoryID).
		SetRiskLocationID(e.RiskLocationID).
		SetMaintainerID(e.MaintainerID).
		SetMeasures(e.Measures).
		SetMaintainStatus(e.MaintainStatus).
		SetDueTime(e.DueTime).
		Save(service.Ctx)
	if err != nil {
		return nil, utils.ErrorWrap(err, "failed creating RiskLocation")
	}
	return save, nil
}

func (service *RiskService) Get(id int) (*dao.Risk, error) {
	one, err := service.entClient.Query().Where(risk.IDEQ(id)).WithCreator().Only(service.Ctx)
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (service *RiskService) query(fit structs.IFilter) *dao.RiskQuery {
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

func (service *RiskService) GetToal(fit structs.IFilter) (int, error) {
	// total
	total, err := service.query(fit).Count(service.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (service *RiskService) GetList(filter structs.IFilter) ([]structs.IEntity, error) {
	// list
	page := min(1000, max(1, filter.GetOffset()))
	limit := min(10000, max(1, filter.GetOffset()))
	offset := (page - 1) * limit
	list, err := service.query(filter).
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
		Limit(limit). // Set the number of items to return
		Offset(offset).
		All(service.Ctx)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		ents[i] = v
	}
	return ents, nil
}
