package security

import (
	"aisecurity/ent/dao"
	"aisecurity/handlers"
	"aisecurity/services/admin/security"
	"aisecurity/utils/http"
	"github.com/gin-gonic/gin"
)

type RiskHandler struct {
	handlers.Handler
	service *security.RiskService
}

func NewRiskHandler() *RiskHandler {
	h := &RiskHandler{}
	h.HandlerService = security.NewRiskService()
	h.service = h.HandlerService.(*security.RiskService)
	return h
}

func (handler *RiskHandler) Create(c *gin.Context) {
	var data dao.Risk
	if err := c.BindJSON(&data); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := handler.service.Create(&data)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}

func (handler *RiskHandler) GetList(c *gin.Context) {
	list, err := handler.service.GetList()
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, list)
}

// CreateRiskLocation 创建风险地点
func (handler *RiskHandler) CreateRiskLocation(c *gin.Context) {
	var data dao.RiskLocation
	if err := c.BindJSON(&data); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := handler.service.CreateRiskLocation(data.Title)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}

// CreateRiskCategory 创建风险分类
func (handler *RiskHandler) CreateRiskCategory(c *gin.Context) {
	var data dao.RiskCategory
	if err := c.BindJSON(&data); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := handler.service.CreateRiskCategory(data.Title)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}
