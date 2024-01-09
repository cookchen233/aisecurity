package filters

import (
	"aisecurity/structs"
)

type RiskCategory struct {
	structs.StandardFilter
	Name string `form:"name"`
}
