package types

import "aisecurity/properties"

type MaintainStatusCounts struct {
	MaintainStatus properties.MaintainStatus `json:"maintain_status"`
	Count          int                       `json:"count"`
}
