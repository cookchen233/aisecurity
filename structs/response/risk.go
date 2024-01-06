package response

import (
	"aisecurity/ent/dao"
	"aisecurity/properties"
	"time"
)

type Risk struct {
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 创建时间
	CreatedAt time.Time `json:"created_at"`
	// 创建者
	CreatedBy int `json:"created_by"`
	// 删除时间
	DeletedAt *time.Time `json:"deleted_at"`
	// 最后更新者
	UpdatedBy int `json:"updated_by"`
	// 最后更新时间
	UpdatedAt time.Time `json:"updated_at"`
	// 标题
	Title string `json:"title"`
	// 内容
	Content string `json:"content"`
	// 图片
	Images []struct {
		Title string "json:\"title\""
		URL   string "json:\"url\""
	} `json:"images"`
	// 风险类别
	RiskCategoryID int `json:"risk_category_id"`
	// 地点
	RiskLocationID int `json:"risk_location_id"`
	// 整改人
	MaintainerID int `json:"maintainer_id"`
	// 整改措施
	Measures string `json:"measures"`
	// 整改状态
	MaintainStatus      properties.MaintainStatus `json:"maintain_status"`
	MaintainStatusLabel string                    `json:"maintain_status_label"`
	// 计划完成日期
	DueTime time.Time `json:"due_time"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RiskQuery when eager-loading is set.
	Edges dao.RiskEdges `json:"edges"`
}
