package auth

import (
	"aisecurity/services/admin/auth"
	"aisecurity/utils/http"
	"github.com/gin-gonic/gin"
)

// RoleHandler Toycar handles
type RoleHandler struct {
	msgs        []byte
	isOnline    bool
	roleService *auth.RoleService
}

func NewRoleHandler() *RoleHandler {
	return &RoleHandler{
		roleService: auth.NewRoleService(),
	}
}

// Index 首页
func (handler *RoleHandler) Index(c *gin.Context) {
	modules, err := handler.roleService.GetModules()
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, modules)
}
