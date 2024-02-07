package entities

import (
	"aisecurity/ent/dao"
)

type Department struct {
	dao.Department
	EmployeeAdmins []*dao.Admin      `json:"employee_admins"`
	Permissions    []*dao.Permission `json:"permissions"`
	Parent         *dao.Department   `json:"parent"`
}
