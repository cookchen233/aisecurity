package types

import (
	"aisecurity/enums"
)

type IEEventStatusCounts struct {
	EventStatus enums.EventStatus `json:"event_status"`
	Label       string            `json:"label"`
	Count       int               `json:"count"`
}
