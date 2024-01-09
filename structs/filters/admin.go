package filters

import (
	"aisecurity/structs"
)

type Admin struct {
	structs.StandardFilter
	Nickname string `form:"nickname"`
}
