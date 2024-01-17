package filters

import (
	"aisecurity/enums"
	"aisecurity/structs"
)

type IPCReportEvent struct {
	structs.StandardFilter
	Keywords    string                     `form:"keywords"`
	VideoID     int                        `form:"video_id"`
	EventTypes  []enums.IPCReportEventType `form:"event_type"`
	EventStatus enums.IPCReportEventStatus `form:"event_status"`
}
