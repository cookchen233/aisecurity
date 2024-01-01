package auth

import (
	"aisecurity/services/admin/auth"
	"aisecurity/utils/http"
	"github.com/gin-gonic/gin"
)

// AdminHandler Toycar handles
type AdminHandler struct {
	msgs         []byte
	isOnline     bool
	adminService *auth.AdminService
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{
		adminService: auth.NewAdminService(),
	}
}

type postData struct {
	Title string `json:"title"`
}

// Create one record
func (handler *AdminHandler) Create(c *gin.Context) {
	modules, err := handler.adminService.CreateOne()
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, modules)
}
