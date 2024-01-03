package security

import (
	"aisecurity/ent/dao"
	"aisecurity/services"
	"aisecurity/utils/db"
	"fmt"
)

type RiskService struct {
	services.Service
}

func NewRiskService() *RiskService {
	return &RiskService{}
}

func (service *RiskService) Create(data *dao.Risk) (*dao.Risk, error) {
	save, err := db.EntClient.Risk.Create().
		SetTitle(data.Title).
		SetContent(data.Content).
		SetImages(data.Images).
		SetRiskCategoryID(data.RiskCategoryID).
		SetRiskLocationID(data.RiskLocationID).
		SetMaintainer(data.Maintainer).
		SetMeasures(data.Measures).
		SetDueTime(data.DueTime).
		Save(service.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating RiskLocation: %w", err)
	}
	return save, nil
}

func (service *RiskService) GetList() ([]*dao.Risk, error) {
	all, err := db.EntClient.Risk.Query().WithCreator().All(service.Ctx)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (service *RiskService) CreateRiskLocation(title string) (*dao.RiskLocation, error) {
	saved, err := db.EntClient.RiskLocation.Create().SetTitle(title).Save(service.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating RiskLocation: %w", err)
	}
	return saved, nil
}

func (service *RiskService) CreateRiskCategory(title string) (*dao.RiskCategory, error) {
	saved, err := db.EntClient.RiskCategory.Create().SetTitle(title).Save(service.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating RiskLocation: %w", err)
	}
	return saved, nil
}
