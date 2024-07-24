package filters

import (
	"aisecurity/enums"
	"aisecurity/structs"
)

type Admin struct {
	structs.StandardFilter
	AdminStatus  enums.EnabledStatus `form:"admin_status"`
	DepartmentID int                 `form:"department_id"`
	OccupationID int                 `form:"occupation_id"`
}
