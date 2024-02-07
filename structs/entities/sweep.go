package entities

import (
	"aisecurity/ent/dao"
	"aisecurity/structs/types"
)

type Sweep struct {
	dao.Sweep
	RiskLocation *dao.RiskLocation `json:"risk_location"`
	RiskCategory *dao.RiskCategory `json:"risk_category"`
	SweepJobs    []*types.SweepJob `json:"sweep_jobs"`
}
