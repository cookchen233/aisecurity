package security

import (
	"aisecurity/ent"
	"aisecurity/ent/admin"
	"aisecurity/ent/riskreporting"
	"aisecurity/utils/db"
	"context"
	"fmt"
	"log"
)

type RiskReportingService struct {
	msgs     []byte
	isOnline bool
}

func NewRiskReportingService() *RiskReportingService {
	return &RiskReportingService{}
}

type Item struct {
	Title        string
	CategoryId   int64
	LocationId   int64
	Handler      int64
	HandlerDept  int64
	CreateBy     int64
	CreateByDept int64
	Mesures      string
	HandleStatus int64
}

func (service *RiskReportingService) Create(data *ent.RiskReporting) (*ent.RiskReporting, error) {
	adminData, _ := db.EntClient.Admin.Query().
		Where(admin.IDEQ(1)).
		Only(context.Background())
	log.Printf("data.Measures %v", data.Measures)
	save, err := db.EntClient.RiskReporting.Create().
		SetTitle(data.Title).
		SetContent(data.Content).
		SetAdminRiskReporting(adminData).
		SetImages(data.Images).
		SetRiskCategoryID(data.RiskCategoryID).
		SetRiskLocationID(data.RiskLocationID).
		SetMaintainer(data.Maintainer).
		//SetMeasures(*data.Measures).
		SetDueTime(data.DueTime).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed creating RiskLocation: %w", err)
	}
	return save, nil
}

func (service *RiskReportingService) GetList() ([]*ent.RiskReporting, error) {
	all, err := db.EntClient.RiskReporting.Query().Select(riskreporting.FieldTitle).All(context.Background())
	if err != nil {
		return nil, err
	}
	return all, nil
}
