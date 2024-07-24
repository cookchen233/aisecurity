package filters

import (
	"aisecurity/structs"
)

type Fixing struct {
	structs.StandardFilter
	EventID int `form:"event_id"`
}
