package entities

import (
	"aisecurity/ent/dao"
)

type Risk struct {
	dao.Risk
	Maintainer          *dao.Admin        `json:"maintainer"`
	RiskLocation        *dao.RiskLocation `json:"risk_location"`
	RiskCategory        *dao.RiskCategory `json:"risk_category"`
	MaintainStatusLabel string            `json:"maintain_status_label"`
}
