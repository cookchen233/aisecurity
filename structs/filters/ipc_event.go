package filters

import (
	"aisecurity/enums"
	"aisecurity/structs"
	"time"
)

type Event struct {
	structs.StandardFilter
	Keywords       string            `form:"keywords"`
	FixerIDs       []int             `form:"fixer_ids[]"`
	VideoID        int               `form:"video_id"`
	DeviceIDs      []int             `form:"device_ids[]"`
	EventTypes     []enums.EventType `form:"event_types[]"`
	EventLevelIDs  []int             `form:"event_level_ids[]"`
	EventTimeRange struct {
		Start time.Time `form:"event_time_range[start]"`
		End   time.Time `form:"event_time_range[end]"`
	}
	FixTimeRange struct {
		Start time.Time `form:"fix_time_range[start]"`
		End   time.Time `form:"fix_time_range[end]"`
	}
	EventStatus enums.EventStatus `form:"status"`
}
