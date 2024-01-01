package auth

import (
	"aisecurity/ent"
	"aisecurity/utils/db"
	"context"
	"fmt"
)

type AdminService struct {
	msgs     []byte
	isOnline bool
}

func NewAdminService() *AdminService {
	return &AdminService{}
}

var ()

func (service *AdminService) CreateOne() (*ent.Admin, error) {
	save, err := db.EntClient.Admin.Create().SetUsername("admin2").Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	return save, nil
}
