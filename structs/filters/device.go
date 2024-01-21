package filters

import (
	"aisecurity/structs"
)

type Device struct {
	structs.StandardFilter
	SN string `form:"sn"`
}
