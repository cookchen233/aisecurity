package routes

import (
	"aisecurity/handlers"
	"aisecurity/handlers/dashboard"
	"github.com/gin-gonic/gin"
)

type DashboardRoutes struct {
	RouterGroup *gin.RouterGroup
}

func NewDashboardRoutes(router *gin.RouterGroup) *DashboardRoutes {
	return &DashboardRoutes{
		RouterGroup: router,
	}
}

// Setup Register admin handlers
func (routes *DashboardRoutes) Setup() {
	adminHandler := dashboard.NewAdminHandler()
	{
		routes.RouterGroup.POST("/admin/create", handlers.Convert(adminHandler, adminHandler.Create))
		routes.RouterGroup.POST("/admin/update", handlers.Convert(adminHandler, adminHandler.Update))
		routes.RouterGroup.GET("/admin/list", handlers.Convert(adminHandler, adminHandler.GetList))
	}

	adminRoleHandler := dashboard.NewAdminRoleHandler()
	{
		routes.RouterGroup.POST("/admin-role/create", adminRoleHandler.Create)
		routes.RouterGroup.GET("/admin-role/list", adminRoleHandler.GetList)
		routes.RouterGroup.GET("/admin-role/modules", adminRoleHandler.GetModules)
	}

	departmentHandler := dashboard.NewDepartmentHandler()
	{
		routes.RouterGroup.POST("/department/create", handlers.Convert(departmentHandler, departmentHandler.Create))
		routes.RouterGroup.GET("/department/list", handlers.Convert(departmentHandler, departmentHandler.GetList))
	}

	employeeHandler := dashboard.NewEmployeeHandler()
	{
		routes.RouterGroup.POST("/employee/create", handlers.Convert(employeeHandler, employeeHandler.Create))
		routes.RouterGroup.GET("/employee/list", handlers.Convert(employeeHandler, employeeHandler.GetList))
	}

	occupationHandler := dashboard.NewOccupationHandler()
	{
		routes.RouterGroup.POST("/occupation/create", handlers.Convert(occupationHandler, occupationHandler.Create))
		routes.RouterGroup.GET("/occupation/list", handlers.Convert(occupationHandler, occupationHandler.GetList))
	}

	riskHandler := dashboard.NewRiskHandler()
	{
		routes.RouterGroup.POST("/risk/create", handlers.Convert(riskHandler, riskHandler.Create))
		routes.RouterGroup.GET("/risk/list", handlers.Convert(riskHandler, riskHandler.GetList))
	}

	riskCategoryHandler := dashboard.NewRiskCategoryHandler()
	{
		routes.RouterGroup.POST("/risk-category/create", handlers.Convert(riskCategoryHandler, riskCategoryHandler.Create))
	}

	riskLocationHandler := dashboard.NewRiskLocationHandler()
	{
		routes.RouterGroup.POST("/risk-location/create", handlers.Convert(riskLocationHandler, riskLocationHandler.Create))
	}
}
