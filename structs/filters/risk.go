package filters

import (
	"aisecurity/properties"
	"aisecurity/structs"
)

type Risk struct {
	structs.StandardFilter
	Title           string                    `form:"title"`
	RiskCategoryIDs []int                     `form:"risk_category_ids[]"`
	RiskLocationIDs []int                     `form:"risk_location_ids[]"`
	MaintainerIDs   []int                     `form:"maintainer_ids[]"`
	ReporterIDs     []int                     `form:"reporter_ids[]"`
	CreatedBy       int                       `form:"created_by"`
	MaintainStatus  properties.MaintainStatus `form:"maintain_status"`
}