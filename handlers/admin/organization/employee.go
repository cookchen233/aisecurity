package organization

import (
	"aisecurity/ent/dao"
	"aisecurity/services/admin/organization"
	"aisecurity/utils/http"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	EmployeeService *organization.EmployeeService
}

func NewEmployeeHandler() *EmployeeHandler {
	return &EmployeeHandler{
		EmployeeService: organization.NewEmployeeService(),
	}
}

func (handler *EmployeeHandler) Create(c *gin.Context) {
	var data dao.Employee
	if err := c.BindJSON(&data); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := handler.EmployeeService.Create(&data)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}

func (handler *EmployeeHandler) GetList(c *gin.Context) {
	list, err := handler.EmployeeService.GetList()
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, list)
}

// CreateDepartment 创建部门
func (handler *EmployeeHandler) CreateDepartment(c *gin.Context) {
	var data dao.Department
	if err := c.BindJSON(&data); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := handler.EmployeeService.CreateDepartment(data.Title)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}
