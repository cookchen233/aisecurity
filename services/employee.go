package services

import (
	"aisecurity/ent/dao"
	"aisecurity/ent/dao/admin"
	"aisecurity/ent/dao/employee"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/db"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
)

type EmployeeService struct {
	Service
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{}
}

var ()

func (service *EmployeeService) Create(ent structs.IEntity) (structs.IEntity, error) {
	e := ent.(*entities.Employee)
	save, err := db.EntClient.Employee.Create().
		SetAdminID(e.AdminID).
		SetDepartmentID(e.DepartmentID).
		AddOccupationIDs(e.OccupationID).
		Save(service.Ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating Employee: %w", err)
	}
	return save, nil
}

func (service *EmployeeService) query(fit structs.IFilter) *dao.EmployeeQuery {
	f := fit.(*filters.Employee)
	q := db.EntClient.Employee.Query()
	if f.ID != 0 {
		q = q.Where(employee.IDEQ(f.ID))
	}
	if f.Name != "" {
		q = q.Where(employee.HasAdminWith(admin.NicknameContainsFold(f.Name)))
	}
	return q.Clone()
}

func (service *EmployeeService) GetToal(fit structs.IFilter) (int, error) {
	// total
	total, err := service.query(fit).Count(service.Ctx)
	if err != nil {
		return 0, utils.ErrorWithStack(err)
	}
	return total, nil
}

func (service *EmployeeService) GetList(fit structs.IFilter) ([]structs.IEntity, error) {
	list, err := db.EntClient.Employee.Query().
		WithAdmin().
		WithDepartment().
		WithOccupations().
		Offset(fit.GetOffset()).
		Limit(fit.GetLimit()).
		All(service.Ctx)
	if err != nil {
		return nil, err
	}
	ents := make([]structs.IEntity, len(list))
	for i, v := range list {
		top, err := service.getTopDepartment(v.Edges.Department)
		if err != nil {
			return nil, err
		}
		v2 := new(entities.Employee)
		err = gconv.Struct(v, v2)
		if err != nil {
			return nil, err
		}
		v2.Edges.TopDepartment = top
		ents[i] = v2
	}
	return ents, nil
}

func (service *EmployeeService) getTopDepartment(dept *dao.Department) (*dao.Department, error) {
	for dept != nil {
		parent, err := dept.QueryParent().Only(service.Ctx)
		if err != nil {
			if dao.IsNotFound(err) {
				break
			}
			return nil, err
		}
		dept = parent
	}
	return dept, nil
}