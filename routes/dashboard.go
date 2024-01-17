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
	indexHandler := dashboard.NewIndexHandler()
	{
		routes.RouterGroup.POST("/upload-images", handlers.Convert(indexHandler, indexHandler.UploadImages))
	}

	adminHandler := dashboard.NewAdminHandler()
	{
		routes.RouterGroup.POST("/admin/create", handlers.Convert(adminHandler, adminHandler.Create))
		routes.RouterGroup.POST("/admin/update", handlers.Convert(adminHandler, adminHandler.Update))
		routes.RouterGroup.GET("/admin/list", handlers.Convert(adminHandler, adminHandler.GetList))
	}

	adminRoleHandler := dashboard.NewAdminRoleHandler()
	{
		routes.RouterGroup.POST("/admin-role/create", handlers.Convert(adminRoleHandler, adminRoleHandler.Create))
		routes.RouterGroup.GET("/admin-role/list", handlers.Convert(adminRoleHandler, adminRoleHandler.GetList))
		routes.RouterGroup.GET("/admin-role/modules", handlers.Convert(adminRoleHandler, adminRoleHandler.GetModules))
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
		routes.RouterGroup.POST("/risk/update", handlers.Convert(riskHandler, riskHandler.Update))
		routes.RouterGroup.GET("/risk/list", handlers.Convert(riskHandler, riskHandler.GetList))
		routes.RouterGroup.GET("/risk/details", handlers.Convert(riskHandler, riskHandler.GetDetails))
	}

	riskCategoryHandler := dashboard.NewRiskCategoryHandler()
	{
		routes.RouterGroup.POST("/risk-category/create", handlers.Convert(riskCategoryHandler, riskCategoryHandler.Create))
		routes.RouterGroup.POST("/risk-category/update", handlers.Convert(riskCategoryHandler, riskCategoryHandler.Update))
		routes.RouterGroup.POST("/risk-category/delete", handlers.Convert(riskCategoryHandler, riskCategoryHandler.Delete))
		routes.RouterGroup.GET("/risk-category/list", handlers.Convert(riskCategoryHandler, riskCategoryHandler.GetList))

	}

	riskLocationHandler := dashboard.NewRiskLocationHandler()
	{
		routes.RouterGroup.POST("/risk-location/create", handlers.Convert(riskLocationHandler, riskLocationHandler.Create))
		routes.RouterGroup.POST("/risk-location/update", handlers.Convert(riskLocationHandler, riskLocationHandler.Update))
		routes.RouterGroup.GET("/risk-location/list", handlers.Convert(riskLocationHandler, riskLocationHandler.GetList))
		routes.RouterGroup.POST("/risk-location/delete", handlers.Convert(riskLocationHandler, riskLocationHandler.Delete))

	}
}
