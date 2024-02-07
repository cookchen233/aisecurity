package filters

import (
	"aisecurity/properties/maintain_status"
	"aisecurity/structs"
	"time"
)

type Risk struct {
	structs.StandardFilter
	RiskCategoryIDs []int `form:"risk_category_ids[]"`
	RiskLocationIDs []int `form:"risk_location_ids[]"`
	MaintainerIDs   []int `form:"maintainer_ids[]"`
	CreatorID       int   `form:"creator_id"`
	DueTimeRange    struct {
		Start time.Time `form:"due_time_range[start]"`
		End   time.Time `form:"due_time_range[end]"`
	}
	CreateTimeRange struct {
		Start time.Time `form:"create_time_range[start]"`
		End   time.Time `form:"create_time_range[end]"`
	}
	MaintainStatus maintain_status.MaintainStatus `form:"maintain_status"`
	Status         maintain_status.MaintainStatus `form:"status"`
}
