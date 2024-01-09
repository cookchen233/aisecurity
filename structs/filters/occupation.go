package filters

import (
	"aisecurity/structs"
)

type Occupation struct {
	structs.StandardFilter
	Name string `form:"name"`
}
