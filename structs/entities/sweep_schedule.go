package entities

import (
	"aisecurity/ent/dao"
)

type SweepSchedule struct {
	dao.SweepSchedule
	Sweep   *dao.Sweep   `json:"sweep"`
	Workers []*dao.Admin `json:"workers"`
}
