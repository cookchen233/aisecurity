package posts

import (
	"aisecurity/enums"
	"aisecurity/structs"
	"aisecurity/structs/entities"
)

type FlowFixing struct {
	structs.StandardPost
	entities.Fixing
	EventLogType enums.EventLogType `json:"event_log_type"`
}
