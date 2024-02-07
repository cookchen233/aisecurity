package filters

import (
	"aisecurity/enums"
	"aisecurity/structs"
)

type Admin struct {
	structs.StandardFilter
	AdminStatus  enums.AdminStatus `form:"admin_status"`
	DepartmentID int               `form:"department_id"`
	OccupationID int               `form:"occupation_id"`
}
