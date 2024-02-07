package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/department"
	employee2 "aisecurity/ent/dao/employee"
	"aisecurity/expects"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go.uber.org/zap"
	"log"
)

type DepartmentService struct {
	Service
	entClient *dao.DepartmentClient
}

func NewDepartmentService(ctx context.Context) *DepartmentService {
	s := &DepartmentService{entClient: db.EntClient.Department}
	s.Ctx = ctx
	return s
}

var ()

func (s *DepartmentService) Create(ent structs.IEntity) (structs.IEntity, error) {
	tx, err := db.EntClient.Tx(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(err))
	}
	e := ent.(*entities.Department)
	c := tx.Department.Create().
		SetName(e.Name).
		SetNotes(e.Notes)
	if e.Parent != nil {
		c.SetParent(e.Parent)
	}
	if len(e.Permissions) > 0 {
		c.AddPermissions(e.Permissions...)
	}
	saved, err := c.Save(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "failed creating Department"))
	}
	if len(e.EmployeeAdmins) > 0 {
		var creates []*dao.EmployeeCreate
		for _, admin := range e.EmployeeAdmins {
			employee, err := db.EntClient.Employee.Query().Where(employee2.AdminID(admin.ID)).First(s.Ctx)
			if err != nil {
				if !dao.IsNotFound(err) {
					return nil, s.rollback(tx, err)
				}
			} else {
				_, err := tx.Employee.UpdateOne(employee).SetDepartment(saved).Save(s.Ctx)
				if err != nil {
					return nil, s.rollback(tx, utils.ErrorWrap(err, "failed updating Employee"))
				}
				continue
			}
			create := tx.Employee.Create().
				SetAdmin(admin).
				SetDepartment(saved)
			creates = append(creates, create)
		}
		if len(creates) > 0 {
			_, err := tx.Employee.CreateBulk(creates...).Save(s.Ctx)
			if err != nil {
				return nil, s.rollback(tx, utils.ErrorWrap(err, "failed creating Employees"))
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "commit failed"))
	}
	return structs.ConvertTo[*dao.Department, entities.Department](saved), nil
}

func (s *DepartmentService) Update(ent structs.IEntity) (structs.IEntity, error) {
	tx, err := db.EntClient.Tx(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWithStack(err))
	}
	e := ent.(*entities.Department)
	u := tx.Department.UpdateOneID(e.ID).
		SetName(e.Name).
		SetNotes(e.Notes).
		ClearParent().ClearPermissions().ClearEmployees()
	if e.Parent != nil {
		u.SetParent(e.Parent)
	}
	if len(e.Permissions) > 0 {
		u.AddPermissions(e.Permissions...)
	}
	saved, err := u.Save(s.Ctx)
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "failed updating Department"))
	}
	if len(e.EmployeeAdmins) > 0 {
		var creates []*dao.EmployeeCreate
		for _, admin := range e.EmployeeAdmins {
			employee, err := db.EntClient.Employee.Query().Where(employee2.AdminID(admin.ID)).First(s.Ctx)
			if err != nil {
				if !dao.IsNotFound(err) {
					return nil, s.rollback(tx, err)
				}
			} else {
				_, err := tx.Employee.UpdateOne(employee).SetDepartment(saved).Save(s.Ctx)
				if err != nil {
					return nil, s.rollback(tx, utils.ErrorWrap(err, "failed updating Employee"))
				}
				continue
			}
			create := db.EntClient.Employee.Create().
				SetAdmin(admin).
				SetDepartment(saved)
			creates = append(creates, create)
		}
		if len(creates) > 0 {
			_, err := tx.Employee.CreateBulk(creates...).Save(s.Ctx)
			if err != nil {
				return nil, s.rollback(tx, utils.ErrorWrap(err, "failed creating Employees"))
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, s.rollback(tx, utils.ErrorWrap(err, "commit failed"))
	}
	return structs.ConvertTo[*dao.Department, entities.Department](saved), nil
}

func (s *DepartmentService) GetDetails(fit structs.IFilter) (structs.IEntity, error) {
	fit.SetPage(1)
	fit.SetLimit(1)
	list, err := s.GetList(fit)
	if err != nil {
		return nil, utils.ErrorWithStack(err)
	}
	if len(list) == 0 {
		return nil, utils.ErrorWithStack(expects.NewDataNotFound())
	}
	return list[0], nil
}

func (s *DepartmentService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	list, err := s.query(fit).
		WithChildren().
		WithParent().
		WithPermissions().
		WithEmployees(func(query *dao.EmployeeQuery) {
			query.WithAdmin()
		}).
		Order(dao.Desc(department.FieldID)).
		All(s.Ctx)
	if err != nil {
		return nil, err
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		allEmployees, err := s.getAllEmployees(v)
		if err != nil {
			log.Fatalf("failed to get employees for department %v: %v", v.ID, err)
		}
		v.AllEmployees = allEmployees
		v2 := new(entities.Department)
		ents[i] = v
		err = gconv.Struct(v, v2)
		if err != nil {
			utils.Logger.Warn("convert error", zap.Error(err))
		}
	}
	return ents, nil
}

func (s *DepartmentService) query(fit structs.IFilter) *dao.DepartmentQuery {
	f := fit.(*filters.Department)
	q := s.entClient.Query()
	if f.ID != 0 {
		q = q.Where(department.IDEQ(f.ID))
	}
	if f.Name != "" {
		q = q.Where(department.NameEQ(f.Name))
	}
	return q.Clone()
}

func (s *DepartmentService) GetTotal(fit structs.IFilter) (int, error) {
	total, err := s.query(fit).Count(s.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (s *DepartmentService) Tree() ([]*dao.Department, error) {
	list, err := db.EntClient.Department.Query().
		Where(department.Not(department.HasParent())).
		WithEmployees(func(query *dao.EmployeeQuery) {
			query.WithOccupation()
		}).
		All(s.Ctx)
	if err != nil {
		return nil, err
	}
	for _, d := range list {
		err := s.getNestedChildren(d)
		if err != nil {
			return nil, err
		}
	}
	return list, nil
}

func (s *DepartmentService) getNestedChildren(dept *dao.Department) error {
	children, err := dept.QueryChildren().
		WithEmployees(func(query *dao.EmployeeQuery) {
			query.WithOccupation()
		}).
		All(s.Ctx)
	if err != nil {
		return err
	}
	for _, ch := range children {
		err := s.getNestedChildren(ch)
		if err != nil {
			return err
		}
	}
	dept.Edges.Children = children
	return nil
}

func (s *DepartmentService) getAllEmployees(d *dao.Department) ([]*dao.Employee, error) {
	// Base case: if the department has no children, return its direct employees
	if !d.QueryChildren().ExistX(s.Ctx) {
		return d.QueryEmployees().All(s.Ctx)
	}

	// Recursive case: accumulate employee IDs from all child departments
	var employees []*dao.Employee
	children, err := d.QueryChildren().All(s.Ctx)
	if err != nil {
		return nil, err
	}

	for _, child := range children {
		childEmployees, err := s.getAllEmployees(child)
		if err != nil {
			return nil, err
		}
		employees = append(employees, childEmployees...)
	}

	// Append the direct employees of the current department
	directEmployees, err := d.QueryEmployees().WithAdmin().All(s.Ctx)
	if err != nil {
		return nil, err
	}
	employees = append(employees, directEmployees...)

	return employees, nil
}
