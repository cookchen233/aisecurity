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
		routes.RouterGroup.GET("/admin/details", handlers.Convert(adminHandler, adminHandler.GetDetails))
		routes.RouterGroup.GET("/admin/invite-qr-code", handlers.Convert(adminHandler, adminHandler.GetInviteQRCode))
	}

	permissionHandler := dashboard.NewPermissionHandler()
	{
		routes.RouterGroup.POST("/permission/create", handlers.Convert(permissionHandler, permissionHandler.Create))
		routes.RouterGroup.POST("/permission/update", handlers.Convert(permissionHandler, permissionHandler.Update))
		routes.RouterGroup.GET("/permission/list", handlers.Convert(permissionHandler, permissionHandler.GetList))
		routes.RouterGroup.GET("/permission/details", handlers.Convert(permissionHandler, permissionHandler.GetDetails))
		routes.RouterGroup.GET("/permission/access-groups", handlers.Convert(permissionHandler, permissionHandler.GetAccessGroups))
	}

	departmentHandler := dashboard.NewDepartmentHandler()
	{
		routes.RouterGroup.POST("/department/create", handlers.Convert(departmentHandler, departmentHandler.Create))
		routes.RouterGroup.POST("/department/update", handlers.Convert(departmentHandler, departmentHandler.Update))
		routes.RouterGroup.GET("/department/list", handlers.Convert(departmentHandler, departmentHandler.GetList))
		routes.RouterGroup.GET("/department/details", handlers.Convert(departmentHandler, departmentHandler.GetDetails))
	}

	employeeHandler := dashboard.NewEmployeeHandler()
	{
		routes.RouterGroup.POST("/employee/create", handlers.Convert(employeeHandler, employeeHandler.Create))
		routes.RouterGroup.GET("/employee/list", handlers.Convert(employeeHandler, employeeHandler.GetList))
	}

	occupationHandler := dashboard.NewOccupationHandler()
	{
		routes.RouterGroup.POST("/occupation/create", handlers.Convert(occupationHandler, occupationHandler.Create))
		routes.RouterGroup.POST("/occupation/update", handlers.Convert(occupationHandler, occupationHandler.Update))
		routes.RouterGroup.POST("/occupation/delete", handlers.Convert(occupationHandler, occupationHandler.Delete))
		routes.RouterGroup.GET("/occupation/list", handlers.Convert(occupationHandler, occupationHandler.GetList))
		routes.RouterGroup.GET("/occupation/details", handlers.Convert(occupationHandler, occupationHandler.GetDetails))
	}

	sweepHandler := dashboard.NewSweepHandler()
	{
		routes.RouterGroup.POST("/sweep/create", handlers.Convert(sweepHandler, sweepHandler.Create))
		routes.RouterGroup.POST("/sweep/update", handlers.Convert(sweepHandler, sweepHandler.Update))
		routes.RouterGroup.POST("/sweep/delete", handlers.Convert(sweepHandler, sweepHandler.Delete))
		routes.RouterGroup.GET("/sweep/list", handlers.Convert(sweepHandler, sweepHandler.GetList))
		routes.RouterGroup.GET("/sweep/details", handlers.Convert(sweepHandler, sweepHandler.GetDetails))
	}

	sweepScheduleHandler := dashboard.NewSweepScheduleHandler()
	{
		routes.RouterGroup.POST("/sweep-schedule/create", handlers.Convert(sweepScheduleHandler, sweepScheduleHandler.Create))
		routes.RouterGroup.POST("/sweep-schedule/update", handlers.Convert(sweepScheduleHandler, sweepScheduleHandler.Update))
		routes.RouterGroup.POST("/sweep-schedule/delete", handlers.Convert(sweepScheduleHandler, sweepScheduleHandler.Delete))
		routes.RouterGroup.GET("/sweep-schedule/list", handlers.Convert(sweepScheduleHandler, sweepScheduleHandler.GetList))
		routes.RouterGroup.GET("/sweep-schedule/details", handlers.Convert(sweepScheduleHandler, sweepScheduleHandler.GetDetails))
	}

	sweepResultHandler := dashboard.NewSweepResultHandler()
	{
		routes.RouterGroup.POST("/sweep-result/create", handlers.Convert(sweepResultHandler, sweepResultHandler.Create))
		routes.RouterGroup.POST("/sweep-result/update", handlers.Convert(sweepResultHandler, sweepResultHandler.Update))
		routes.RouterGroup.POST("/sweep-result/delete", handlers.Convert(sweepResultHandler, sweepResultHandler.Delete))
		routes.RouterGroup.POST("/sweep-result/check-in", handlers.Convert(sweepResultHandler, sweepResultHandler.CheckIn))
		routes.RouterGroup.POST("/sweep-result/start-work", handlers.Convert(sweepResultHandler, sweepResultHandler.StartWork))
		routes.RouterGroup.GET("/sweep-result/list", handlers.Convert(sweepResultHandler, sweepResultHandler.GetList))
		routes.RouterGroup.GET("/sweep-result/details", handlers.Convert(sweepResultHandler, sweepResultHandler.GetDetails))
	}

	riskHandler := dashboard.NewRiskHandler()
	{
		routes.RouterGroup.POST("/risk/create", handlers.Convert(riskHandler, riskHandler.Create))
		routes.RouterGroup.POST("/risk/update", handlers.Convert(riskHandler, riskHandler.Update))
		routes.RouterGroup.GET("/risk/list", handlers.Convert(riskHandler, riskHandler.GetList))
		routes.RouterGroup.GET("/risk/details", handlers.Convert(riskHandler, riskHandler.GetDetails))
		routes.RouterGroup.GET("/risk/status-counts", handlers.Convert(riskHandler, riskHandler.GetStatusCounts))
		routes.RouterGroup.GET("/risk/risk-category-counts", handlers.Convert(riskHandler, riskHandler.GetRiskCategoryCounts))
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

	ipcEventHandler := dashboard.NewEventHandler()
	{
		routes.RouterGroup.GET("/ipc-event/list", handlers.Convert(ipcEventHandler, ipcEventHandler.GetList))
		routes.RouterGroup.GET("/ipc-event/details", handlers.Convert(ipcEventHandler, ipcEventHandler.GetDetails))
		routes.RouterGroup.GET("/ipc-event/status-counts", handlers.Convert(ipcEventHandler, ipcEventHandler.GetStatusCounts))
		routes.RouterGroup.GET("/ipc-event/event-type-counts", handlers.Convert(ipcEventHandler, ipcEventHandler.GetEventTypeCounts))
		routes.RouterGroup.GET("/ipc-event/event-device-counts", handlers.Convert(ipcEventHandler, ipcEventHandler.GetEventDeviceCounts))
		routes.RouterGroup.GET("/ipc-event/event-time-counts", handlers.Convert(ipcEventHandler, ipcEventHandler.GetEventTimeCounts))
	}

	eventLevelHandler := dashboard.NewEventLevelHandler()
	{
		routes.RouterGroup.POST("/event-level/create", handlers.Convert(eventLevelHandler, eventLevelHandler.Create))
		routes.RouterGroup.POST("/event-level/update", handlers.Convert(eventLevelHandler, eventLevelHandler.Update))
		routes.RouterGroup.GET("/event-level/list", handlers.Convert(eventLevelHandler, eventLevelHandler.GetList))
		routes.RouterGroup.POST("/event-level/delete", handlers.Convert(eventLevelHandler, eventLevelHandler.Delete))
		routes.RouterGroup.GET("/event-level/event-type-list",
			//middlewares.CacheMemory(24*time.Hour),
			handlers.Convert(eventLevelHandler, eventLevelHandler.GetListForEachEventType),
		)
	}

	deviceInstallationHandler := dashboard.NewIDeviceInstallationHandler()
	{
		routes.RouterGroup.POST("/device-installation/create", handlers.Convert(deviceInstallationHandler, deviceInstallationHandler.Create))
		routes.RouterGroup.POST("/device-installation/update", handlers.Convert(deviceInstallationHandler, deviceInstallationHandler.Update))
		routes.RouterGroup.GET("/device-installation/list", handlers.Convert(deviceInstallationHandler, deviceInstallationHandler.GetList))
		routes.RouterGroup.POST("/device-installation/delete", handlers.Convert(deviceInstallationHandler, deviceInstallationHandler.Delete))
	}

	eventLogHandler := dashboard.NewEventLogHandler()
	{
		routes.RouterGroup.GET("/event-log/list", handlers.Convert(eventLogHandler, eventLogHandler.GetList))
	}

	fixingHandler := dashboard.NewFixingHandler()
	{
		routes.RouterGroup.POST("/fixing/flow", handlers.Convert(fixingHandler, fixingHandler.Flow))
	}
}
