package entities

import (
	"aisecurity/ent/dao"
)

type SweepResultDetails struct {
	dao.SweepResultDetails
	Sweep         *dao.Sweep         `json:"sweep"`
	SweepSchedule *dao.SweepSchedule `json:"sweep_schedule"`
}
