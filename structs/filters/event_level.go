package filters

import (
	"aisecurity/enums"
	"aisecurity/structs"
)

type EventLevel struct {
	structs.StandardFilter
	EventType enums.EventType `form:"event_type"`
}
