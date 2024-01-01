package security

import (
	"aisecurity/ent"
	"aisecurity/utils/db"
	"context"
	"fmt"
)

type RiskService struct {
	msgs     []byte
	isOnline bool
}

func NewRiskService() *RiskService {
	return &RiskService{}
}

var ()

func (service *RiskService) CreateRiskLocation(title string) (*ent.RiskLocation, error) {
	save, err := db.EntClient.RiskLocation.Create().SetTitle(title).Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed creating RiskLocation: %w", err)
	}
	return save, nil
}
