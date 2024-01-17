package types

import (
	"aisecurity/enums"
)

type IREEventStatusCounts struct {
	EventStatus enums.IPCReportEventStatus `json:"event_status"`
	Label       string                     `json:"label"`
	Count       int                        `json:"count"`
}
