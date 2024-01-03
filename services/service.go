package services

import "context"

type Service struct {
	Ctx context.Context
}

func (service *Service) SetContext(ctx context.Context) {
	service.Ctx = ctx
}
