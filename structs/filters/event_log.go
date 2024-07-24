package filters

import (
	"aisecurity/structs"
)

type EventLog struct {
	structs.StandardFilter
	DeviceID int `form:"device_id"`
	EventID  int `form:"event_id"`
}
