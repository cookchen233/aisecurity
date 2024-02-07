package filters

import (
	"aisecurity/structs"
)

type EventLog struct {
	structs.StandardFilter
	DeviceID int `json:"device_id"`
	EventID  int `json:"event_id"`
}
