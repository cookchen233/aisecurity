package organization

import (
	"aisecurity/ent/dao"
	"aisecurity/handlers"
	"aisecurity/services/admin/organization"
	"aisecurity/utils/http"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	handlers.Handler
	service *organization.EmployeeService
}

func NewEmployeeHandler() *EmployeeHandler {
	h := &EmployeeHandler{}
	h.HandlerService = organization.NewEmployeeService()
	h.service = h.HandlerService.(*organization.EmployeeService)
	return h
}

func (handler *EmployeeHandler) Create(c *gin.Context) {
	var req dao.Employee
	if err := c.ShouldBind(&req); err != nil {
		http.Error(c, err, 900)
		return
	}
	if err := handler.Validate(req); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := handler.service.Create(&req)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}

func (handler *EmployeeHandler) GetList(c *gin.Context) {
	list, err := handler.service.GetList()
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, list)
}

// CreateDepartment 创建部门
func (handler *EmployeeHandler) CreateDepartment(c *gin.Context) {
	var req dao.Department
	if err := c.ShouldBindJSON(&req); err != nil {
		http.Error(c, err, 900)
		return
	}
	if err := handler.Validate(req); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := handler.service.CreateDepartment(req)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}

// GetDepartmentList 获取部门
func (handler *EmployeeHandler) GetDepartmentList(c *gin.Context) {
	var req dao.Department
	if err := c.ShouldBindJSON(&req); err != nil {
		http.Error(c, err, 900)
		return
	}
	if err := handler.Validate(req); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := handler.service.GetDepartmentList()
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}
