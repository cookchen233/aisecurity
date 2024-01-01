package security

import (
	"aisecurity/ent"
	"aisecurity/services/admin/security"
	"aisecurity/utils/http"
	"github.com/gin-gonic/gin"
)

type RiskReportingHandler struct {
	msgs                 []byte
	isOnline             bool
	RiskReportingService security.RiskReportingService
	RiskService          security.RiskService
}

type postData struct {
}

type locationData struct {
	Title string `json:"title"`
}

func NewRiskReportingHandler() *RiskReportingHandler {
	return &RiskReportingHandler{}
}

// Create 创建隐患汇报
func (handler *RiskReportingHandler) Create(c *gin.Context) {
	var data ent.RiskReporting
	if err := c.BindJSON(&data); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := handler.RiskReportingService.Create(&data)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}

// GetList 获取列表
func (handler *RiskReportingHandler) GetList(c *gin.Context) {
	list, err := handler.RiskReportingService.GetList()
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, list)
}

// CreateRiskLocation 创建风险地点
func (handler *RiskReportingHandler) CreateRiskLocation(c *gin.Context) {
	var data locationData
	if err := c.BindJSON(&data); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := handler.RiskService.CreateRiskLocation(data.Title)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}
