package filters

import (
	"aisecurity/structs"
)

type Fixing struct {
	structs.StandardFilter
	EventID int `json:"event_id"`
}
