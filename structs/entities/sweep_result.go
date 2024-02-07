package entities

import (
	"aisecurity/ent/dao"
)

type SweepResult struct {
	dao.SweepResult
	SweepSchedule *dao.SweepSchedule `json:"sweep_schedule"`
}
