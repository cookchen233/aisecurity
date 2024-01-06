package services

import "context"

type IService interface {
	SetContext(c context.Context)
}

type Service struct {
	Ctx context.Context
}

func (service *Service) SetContext(ctx context.Context) {
	service.Ctx = ctx
}
