package admin

import (
	"aisecurity/handlers"
	"aisecurity/handlers/admin/organization"
	"github.com/gin-gonic/gin"
)

type OrganizationRoutes struct {
	RouterGroup *gin.RouterGroup
}

func NewOrganizationRoutes(router *gin.RouterGroup) *OrganizationRoutes {
	return &OrganizationRoutes{
		RouterGroup: router,
	}
}

// Setup Register security handlers
func (routes *OrganizationRoutes) Setup() {
	employeeHandler := organization.NewEmployeeHandler()
	{
		routes.RouterGroup.POST("/employee/create", handlers.Convert(employeeHandler, employeeHandler.Create))
		routes.RouterGroup.GET("/employee/get-list", handlers.Convert(employeeHandler, employeeHandler.GetList))

		routes.RouterGroup.POST("/department/create", handlers.Convert(employeeHandler, employeeHandler.CreateDepartment))
		routes.RouterGroup.GET("/department/get-list", handlers.Convert(employeeHandler, employeeHandler.GetDepartmentList))
	}
}
