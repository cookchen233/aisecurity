package filters

import (
	"aisecurity/structs"
)

type Employee struct {
	structs.StandardFilter
	Name string `form:"name"`
}
