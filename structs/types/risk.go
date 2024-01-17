package types

import (
	"aisecurity/properties/maintain_status"
)

type MaintainStatusCounts struct {
	MaintainStatus maintain_status.MaintainStatus `json:"maintain_status"`
	Label          string                         `json:"label"`
	Count          int                            `json:"count"`
}
