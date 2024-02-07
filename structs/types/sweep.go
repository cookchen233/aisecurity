package types

import (
	"aisecurity/enums"
	"time"
)

type SweepJob struct {
	Title       string                      `json:"title"`
	ResultTypes []*enums.SweepJobResultType `json:"result_types"`
	Result      enums.SweepJobResultType    `json:"result"`
	UpdateTime  time.Time                   `json:"update_time"`
	UpdaterID   int                         `json:"updater_id"`
}

type ScheduleRepeat struct {
	RepeatType  enums.RepeatType  `json:"repeat_type"`
	Interval    int               `json:"interval"`
	ExpiresType enums.ExpiresType `json:"expires_type"`
	Times       int               `json:"times"`
	EndTime     time.Time         `json:"end_time"`
}

type ScheduleRemind struct {
	RemindTypes     []*enums.RemindType `json:"remind_types"`
	AdvancedSeconds int                 `json:"advanced_seconds"`
}
