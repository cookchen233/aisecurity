package filters

import (
	"aisecurity/structs"
	"time"
)

type Sweep struct {
	structs.StandardFilter
	CreateTimeRange struct {
		Start time.Time `form:"create_time_range[start]"`
		End   time.Time `form:"create_time_range[end]"`
	}
	RiskCategoryIDs []int `form:"risk_category_ids[]"`
	RiskLocationIDs []int `form:"risk_location_ids[]"`
}
