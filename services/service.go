package services

import (
	"aisecurity/structs"
	"context"
)

type IService interface {
	SetContext(context.Context)
	Create(structs.IEntity) (structs.IEntity, error)
	Update(structs.IEntity) (structs.IEntity, error)
	GetList(structs.IFilter) ([]structs.IEntity, error)
	GetDetails(structs.IFilter) (structs.IEntity, error)
	Delete(structs.IEntity) error
}

type Service struct {
	Ctx context.Context
}

func (service *Service) SetContext(ctx context.Context) {
	service.Ctx = ctx
}

func (service *Service) Create(structs.IEntity) (structs.IEntity, error)     { return nil, nil }
func (service *Service) Update(structs.IEntity) (structs.IEntity, error)     { return nil, nil }
func (service *Service) GetList(structs.IFilter) ([]structs.IEntity, error)  { return nil, nil }
func (service *Service) GetDetails(structs.IFilter) (structs.IEntity, error) { return nil, nil }
func (service *Service) Delete(structs.IEntity) error                        { return nil }
