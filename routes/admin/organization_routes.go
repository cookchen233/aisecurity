package admin

import (
	"aisecurity/handlers/admin/organization"
	"github.com/gin-gonic/gin"
)

type OrganizationRoutes struct {
	RouterGroup *gin.RouterGroup
}

func NewOrganizationRoutes(router *gin.RouterGroup) *SecurityRoutes {
	return &SecurityRoutes{
		RouterGroup: router,
	}
}

// Setup Register security handlers
func (routes *OrganizationRoutes) Setup() {
	employeeHandler := organization.NewEmployeeHandler()
	routes.RouterGroup.POST("/employee/create", employeeHandler.Create)
	routes.RouterGroup.GET("/employee/get-list", employeeHandler.GetList)
	routes.RouterGroup.POST("/employee/create-department", employeeHandler.CreateDepartment)
}
