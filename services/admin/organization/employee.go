package organization

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/department"
	"aisecurity/services"
	"aisecurity/utils/db"
	"fmt"
)

type EmployeeService struct {
	services.Service
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{}
}

var ()

func (service *EmployeeService) Create(data *dao.Employee) (*dao.Employee, error) {
	save, err := db.EntClient.Employee.Create().
		SetAdminID(data.AdminID).
		SetDepartmentID(data.DepartmentID).
		Save(service.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating Employee: %w", err)
	}
	return save, nil
}

func (service *EmployeeService) GetList() ([]*dao.Employee, error) {
	all, err := db.EntClient.Employee.Query().WithAdmin().WithDepartment().All(service.Ctx)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (service *EmployeeService) CreateDepartment(department dao.Department) (*dao.Department, error) {
	saved, err := db.EntClient.Department.Create().
		SetTitle(department.Title).
		SetParentID(department.ParentID).
		Save(service.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating Department: %w", err)
	}
	return saved, nil
}

func (service *EmployeeService) GetDepartmentList() ([]*dao.Department, error) {
	all, err := db.EntClient.Department.Query().
		Where(department.ParentIDEQ(1)).
		Where(department.IDNEQ(1)).
		WithChildren().All(service.Ctx)
	if err != nil {
		return nil, err
	}
	return all, nil
}
