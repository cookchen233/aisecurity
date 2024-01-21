package filters

import (
	"aisecurity/enums"
	"aisecurity/structs"
)

type IPCEvent struct {
	structs.StandardFilter
	Keywords    string            `form:"keywords"`
	FixerIDs    []int             `form:"fixer_ids[]"`
	VideoID     int               `form:"video_id"`
	EventTypes  []enums.EventType `form:"event_type"`
	EventStatus enums.EventStatus `form:"event_status"`
}
