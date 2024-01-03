package auth

import (
	"aisecurity/ent/dao"
	"aisecurity/services/admin/auth"
	"aisecurity/utils/http"
	"github.com/gin-gonic/gin"
)

// AdminRoleHandler Toycar handles
type AdminRoleHandler struct {
	service *auth.AdminRoleService
}

func NewAdminRoleHandler() *AdminRoleHandler {
	return &AdminRoleHandler{
		service: auth.NewAdminRoleService(),
	}
}

func (handler *AdminRoleHandler) Create(c *gin.Context) {
	var data dao.AdminRole
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

func (handler *AdminRoleHandler) GetList(c *gin.Context) {
	list, err := handler.service.GetList()
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, list)
}

func (handler *AdminRoleHandler) GetModules(c *gin.Context) {
	modules, err := handler.service.GetModules()
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, modules)
}
