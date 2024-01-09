package entities

import (
	"aisecurity/ent/dao"
)

type Risk struct {
	dao.Risk
	MaintainStatusLabel string `json:"maintain_status_label"`
}
