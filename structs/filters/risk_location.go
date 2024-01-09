package filters

import (
	"aisecurity/structs"
)

type RiskLocation struct {
	structs.StandardFilter
	Name string `form:"name"`
}
