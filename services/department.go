package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/department"
	"aisecurity/ent/dao/employee"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/utils/db"
	"fmt"
	"log"
)

type DepartmentService struct {
	Service
}

func NewDepartmentService() *DepartmentService {
	return &DepartmentService{}
}

var ()

func (service *DepartmentService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Department)
	c := db.EntClient.Department.Create().
		SetName(e.Name)
	if e.ParentID > 0 {
		c.SetParentID(e.ParentID)
	}
	saved, err := c.Save(service.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating Department: %w", err)
	}
	return saved, nil
}

// List returns a list of departments with all employees
func (service *DepartmentService) List() ([]*dao.Department, error) {
	topLevelDepartments, err := db.EntClient.Department.Query().
		WithChildren().
		WithParent().
		All(service.Ctx)
	if err != nil {
		return nil, err
	}
	for _, d := range topLevelDepartments {
		employeeIDs, err := service.getAllEmployeeIDs(d)
		if err != nil {
			log.Fatalf("failed to get employees for department %v: %v", d.ID, err)
		}
		employees, err := db.EntClient.Employee.
			Query().
			Where(employee.IDIn(employeeIDs...)).
			WithAdmin().
			WithOccupations().
			All(service.Ctx)
		if err != nil {
			log.Fatalf("failed querying employees: %v", err)
		}
		d.Edges.Employees = employees
	}
	return topLevelDepartments, nil
}

func (service *DepartmentService) Tree() ([]*dao.Department, error) {
	list, err := db.EntClient.Department.Query().
		Where(department.Not(department.HasParent())).
		WithEmployees(func(query *dao.EmployeeQuery) {
			query.WithOccupations()
		}).
		All(service.Ctx)
	if err != nil {
		return nil, err
	}
	for _, d := range list {
		err := service.getNestedChildren(d)
		if err != nil {
			return nil, err
		}
	}
	return list, nil
}

func (service *DepartmentService) getNestedChildren(dept *dao.Department) error {
	children, err := dept.QueryChildren().
		WithEmployees(func(query *dao.EmployeeQuery) {
			query.WithOccupations()
		}).
		All(service.Ctx)
	if err != nil {
		return err
	}
	for _, ch := range children {
		err := service.getNestedChildren(ch)
		if err != nil {
			return err
		}
	}
	dept.Edges.Children = children
	return nil
}

func (service *DepartmentService) getAllEmployeeIDs(d *dao.Department) ([]int, error) {
	// Base case: if the department has no children, return its direct employees
	if !d.QueryChildren().ExistX(service.Ctx) {
		return d.QueryEmployees().IDs(service.Ctx)
	}

	// Recursive case: accumulate employee IDs from all child departments
	var employeeIDs []int
	children, err := d.QueryChildren().All(service.Ctx)
	if err != nil {
		return nil, err
	}

	for _, child := range children {
		childEmployeeIDs, err := service.getAllEmployeeIDs(child)
		if err != nil {
			return nil, err
		}
		employeeIDs = append(employeeIDs, childEmployeeIDs...)
	}

	// Append the direct employees of the current department
	directEmployeeIDs, err := d.QueryEmployees().IDs(service.Ctx)
	if err != nil {
		return nil, err
	}
	employeeIDs = append(employeeIDs, directEmployeeIDs...)

	return employeeIDs, nil
}
