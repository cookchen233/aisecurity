package organization

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/department"
	"aisecurity/utils/db"
	"context"
	"fmt"
)

type EmployeeService struct {
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{}
}

var ()

func (service *EmployeeService) Create(data *dao.Employee) (*dao.Employee, error) {
	adminData, _ := db.EntClient.Admin.Query().Where(admin.IDEQ(data.AdminID)).Only(context.Background())
	departmentData, _ := db.EntClient.Department.Query().Where(department.IDEQ(data.DepartmentID)).Only(context.Background())
	save, err := db.EntClient.Employee.Create().
		SetAdmin(adminData).
		SetDepartment(departmentData).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed creating Employee: %w", err)
	}
	return save, nil
}

func (service *EmployeeService) GetList() ([]*dao.Employee, error) {
	all, err := db.EntClient.Employee.Query().WithAdmin().WithDepartment().All(context.Background())
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (service *EmployeeService) CreateDepartment(title string) (*dao.Department, error) {
	saved, err := db.EntClient.Department.Create().SetTitle(title).Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed creating Department: %w", err)
	}
	return saved, nil
}
