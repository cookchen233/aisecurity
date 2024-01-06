package request

type RiskListFilter struct {
	Page            int    `form:"page"`
	Limit           int    `form:"limit"`
	ID              int    `form:"id"`
	Title           string `form:"title"`
	RiskCategoryIDs []int  `form:"risk_category_ids"`
	RiskLocationIDs []int  `form:"risk_location_ids"`
	Maintainer      int    `form:"maintainer"`
	CreatedBy       int    `form:"created_by"`
	MaintainStatus  int    `form:"maintain_status"`
}
