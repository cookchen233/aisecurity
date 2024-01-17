package services

import (
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/utils"
	"context"
	"go.uber.org/zap"
)

type IService interface {
	SetContext(context.Context)
	Create(structs.IEntity) (structs.IEntity, error)
	Update(structs.IEntity) (structs.IEntity, error)
	GetList(structs.IFilter) ([]structs.IEntity, error)
	GetTotal(structs.IFilter) (int, error)
	GetDetails(structs.IFilter) (structs.IEntity, error)
	Delete(structs.IEntity) error
}

type Service struct {
	Ctx context.Context
}

func (s *Service) SetContext(ctx context.Context) {
	s.Ctx = ctx
}

func (s *Service) Create(structs.IEntity) (structs.IEntity, error) {
	utils.Logger.Error("called empty service method", zap.String("method", utils.GetMethodName()))
	return nil, utils.ErrorWithStack(expects.NewNotImplementedMethod())
}
func (s *Service) Update(structs.IEntity) (structs.IEntity, error) {
	utils.Logger.Error("called empty service method", zap.String("method", utils.GetMethodName()))
	return nil, utils.ErrorWithStack(expects.NewNotImplementedMethod())
}
func (s *Service) GetList(structs.IFilter) ([]structs.IEntity, error) {
	utils.Logger.Error("called empty service method", zap.String("method", utils.GetMethodName()))
	return nil, utils.ErrorWithStack(expects.NewNotImplementedMethod())
}
func (s *Service) GetTotal(structs.IFilter) (int, error) {
	utils.Logger.Error("called empty service method", zap.String("method", utils.GetMethodName()))
	return 0, nil
}
func (s *Service) GetDetails(structs.IFilter) (structs.IEntity, error) {
	utils.Logger.Error("called empty service method", zap.String("method", utils.GetMethodName()))
	return nil, utils.ErrorWithStack(expects.NewNotImplementedMethod())
}
func (s *Service) Delete(structs.IEntity) error {
	utils.Logger.Error("called empty service method", zap.String("method", utils.GetMethodName()))
	return utils.ErrorWithStack(expects.NewNotImplementedMethod())
}
