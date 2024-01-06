package security

import (
	"aisecurity/ent/dao"
	"aisecurity/handlers"
	"aisecurity/properties"
	"aisecurity/services/admin/security"
	"aisecurity/structs"
	"aisecurity/structs/request"
	"aisecurity/structs/response"
	"aisecurity/utils/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
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
	time.Sleep(1000 * time.Millisecond)
	var req request.RiskListFilter
	if err := c.ShouldBindQuery(&req); err != nil {
		http.Error(c, err, 900)
		return
	}
	fmt.Println("req", req)
	list, err := handler.service.GetList(req)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	// add labels to the list
	var newList []response.Risk
	for _, v := range list {
		v2 := structs.ConvertTo[response.Risk](v)
		v2.MaintainStatusLabel = v2.MaintainStatus.Label()
		newList = append(newList, structs.ConvertTo[response.Risk](v2))
	}
	http.Success(c, newList)
}

func (handler *RiskHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http.Error(c, err, 900)
		return
	}
	list, err := handler.service.GetByID(id)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, list)
}

// CreateRiskLocation 创建风险地点
func (handler *RiskHandler) CreateRiskLocation(c *gin.Context) {
	var req dao.RiskLocation
	if err := c.ShouldBindJSON(&req); err != nil {
		http.Error(c, err, 900)
		return
	}
	if err := handler.Validate(req); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := handler.service.CreateRiskLocation(req)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}

// CreateRiskCategory 创建风险分类
func (handler *RiskHandler) CreateRiskCategory(c *gin.Context) {
	var req dao.RiskCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		http.Error(c, err, 900)
		return
	}
	if err := handler.Validate(req); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := handler.service.CreateRiskCategory(req)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}

func (handler *RiskHandler) GetEnums(c *gin.Context) {
	var p properties.MaintainStatus
	enums := map[string]interface{}{
		"maintain_status": p.GetLabeledEnumList(),
	}
	http.Success(c, enums)
}
