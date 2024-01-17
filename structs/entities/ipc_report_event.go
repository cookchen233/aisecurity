package entities

import (
	"aisecurity/ent/dao"
)

type IPCReportEvent struct {
	dao.IPCReportEvent
	EventTypeLabel   string `json:"event_type_label"`
	EventStatusLabel string `json:"event_status_label"`
}
