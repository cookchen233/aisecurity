package filters

import (
	"aisecurity/enums"
	"aisecurity/structs"
	"time"
)

type SweepSchedule struct {
	structs.StandardFilter
	CreateTimeRange struct {
		Start time.Time `form:"create_time_range[start]"`
		End   time.Time `form:"create_time_range[end]"`
	}
	ActionTimeRange struct {
		Start time.Time `form:"action_time_range[start]"`
		End   time.Time `form:"action_time_range[end]"`
	}
	EnabledStatus enums.EnabledStatus `form:"enabled_status"`
	SweepID       int                 `form:"sweep_id"`
	SweepIDs      []int               `form:"sweep_ids[]"`
	WorkerID      int                 `form:"worker_id"`
	WorkerIDs     []int               `form:"worker_ids[]"`
}
