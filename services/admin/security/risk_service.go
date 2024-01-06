package security

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/employee"
	"aisecurity/ent/dao/risk"
	"aisecurity/services"
	"aisecurity/structs/request"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"github.com/pkg/errors"
)

type RiskService struct {
	services.Service
	entClient *dao.RiskClient
}

func NewRiskService() *RiskService {
	return &RiskService{
		entClient: db.EntClient.Risk,
	}
}

func (service *RiskService) Create(data *dao.Risk) (*dao.Risk, error) {
	save, err := db.EntClient.Risk.Create().
		SetTitle(data.Title).
		SetContent(data.Content).
		SetImages(data.Images).
		SetRiskCategoryID(data.RiskCategoryID).
		SetRiskLocationID(data.RiskLocationID).
		SetMaintainerID(data.MaintainerID).
		SetMeasures(data.Measures).
		SetMaintainStatus(data.MaintainStatus).
		SetDueTime(data.DueTime).
		Save(service.Ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed creating RiskLocation")
	}
	return save, nil
}

func (service *RiskService) GetList(filter request.RiskListFilter) ([]*dao.Risk, error) {
	q := db.EntClient.Risk.Query()
	if filter.ID != 0 {
		q.Where(risk.IDEQ(filter.ID))
	}
	if filter.Title != "" {
		q.Where(risk.TitleContainsFold(filter.Title))
	}
	if filter.CreatedBy != 0 {
		q.Where(risk.HasCreatorWith(employee.IDEQ(filter.CreatedBy)))
	}
	if filter.Maintainer != 0 {
		q.Where(risk.HasMaintainerWith(employee.IDEQ(filter.Maintainer)))
	}
	if filter.MaintainStatus != 0 {

	}
	categoryIDs := utils.FilterZerosFromArray(filter.RiskCategoryIDs)
	if len(categoryIDs) > 0 {
		q.Where(risk.RiskCategoryIDIn(categoryIDs...))
	}
	locationIDs := utils.FilterZerosFromArray(filter.RiskLocationIDs)
	if len(locationIDs) > 0 {
		q.Where(risk.RiskLocationIDIn(locationIDs...))
	}
	q.Where(risk.RiskLocationIDIn(filter.RiskLocationIDs...))
	page := min(1000, max(1, filter.Page))
	limit := min(10000, max(1, filter.Limit))
	offset := (page - 1) * limit
	all, err := q.
		WithRiskLocation().
		WithRiskCategory().
		WithMaintainer(func(query *dao.EmployeeQuery) {
			query.WithDepartment().WithAdmin(func(query *dao.AdminQuery) {
				query.WithAdminRoles().Limit(1)
			})
		}).
		WithCreator(func(query *dao.EmployeeQuery) {
			query.WithDepartment().WithAdmin(func(query *dao.AdminQuery) {
				query.WithAdminRoles().Limit(1)
			})
		}).
		Limit(limit). // Set the number of items to return
		Offset(offset).
		All(service.Ctx)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (service *RiskService) CreateRiskLocation(location dao.RiskLocation) (*dao.RiskLocation, error) {
	saved, err := db.EntClient.RiskLocation.Create().SetTitle(location.Title).Save(service.Ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed creating RiskLocation")
	}
	return saved, nil
}

func (service *RiskService) CreateRiskCategory(category dao.RiskCategory) (*dao.RiskCategory, error) {
	saved, err := db.EntClient.RiskCategory.Create().SetTitle(category.Title).Save(service.Ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed creating RiskCategory")
	}
	return saved, nil
}

func (service *RiskService) GetByID(id int) (*dao.Risk, error) {
	one, err := service.entClient.Query().Where(risk.IDEQ(id)).WithCreator().Only(service.Ctx)
	if err != nil {
		return nil, err
	}
	return one, nil
}
