package auth

import (
	"aisecurity/ent/dao"
	"aisecurity/handlers"
	"aisecurity/services/admin/auth"
	"aisecurity/structs"
	"aisecurity/structs/request"
	"aisecurity/structs/service"
	"aisecurity/utils/http"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// AdminHandler Toycar handles
type AdminHandler struct {
	handlers.Handler
	service *auth.AdminService
}

func NewAdminHandler() *AdminHandler {
	h := &AdminHandler{}
	h.HandlerService = auth.NewAdminService()
	h.service = h.HandlerService.(*auth.AdminService)
	return h
}

func (handler *AdminHandler) CreateSuperAdmin(c *gin.Context) {
	var req request.Admin
	if err := c.BindJSON(&req); err != nil {
		http.Error(c, err, 900)
		return
	}
	role, err := auth.NewAdminRoleService().GetSuperRole()
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	saved, err := handler.service.Create(structs.ConvertTo[service.Admin](req), role)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}

func (handler *AdminHandler) Create(c *gin.Context) {
	var req request.Admin
	if err := c.BindJSON(&req); err != nil {
		http.Error(c, err, 900)
		return
	}
	if err := handler.Validate(req); err != nil {
		http.Error(c, errors.WithStack(errors.Wrap(err, "x")), 900)
		return
	}
	role, err := auth.NewAdminRoleService().GetByID(req.AdminRoleID)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	saved, err := handler.service.Create(structs.ConvertTo[service.Admin](req), role)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}

func (handler *AdminHandler) Update(c *gin.Context) {
	var req request.Admin
	if err := c.BindJSON(&req); err != nil {
		http.Error(c, err, 900)
		return
	}
	if err := handler.Validate(req); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := handler.service.Update(structs.ConvertTo[service.Admin](req))
	if err != nil {
		http.Error(c, errors.WithStack(err), 1000)
		return
	}
	http.Success(c, saved)
}

func (handler *AdminHandler) GetList(c *gin.Context) {
	list, err := handler.service.GetList()
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, list)
}

func (handler *AdminHandler) Login(c *gin.Context) {
	var req request.AdminLogin
	if err := c.BindJSON(&req); err != nil {
		http.Error(c, err, 900)
		return
	}
	if err := handler.Validate(req); err != nil {
		http.Error(c, err, 900)
		return
	}

	admin, err := handler.service.GetByUserName(req.Username)
	if err != nil {
		if dao.IsNotFound(err) {
			http.Error(c, fmt.Errorf("用户名或密码错误"), 900)
			return
		}
		http.Error(c, err, 900)
		return
	}

	// Compare the provided password with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		http.Error(c, fmt.Errorf("用户名或密码错误"), 900)
		return
	}

	// Create a session and set its values
	session := sessions.Default(c)
	session.Set("authenticated", true)
	session.Set("admin_id", admin.ID)
	persistSeconds := 0
	if req.PersistDays > 0 {
		persistSeconds = max(req.PersistDays, 30) * 86400
	}
	session.Options(sessions.Options{
		Path:     "/",
		MaxAge:   persistSeconds,
		HttpOnly: true,
	})
	err = session.Save()
	if err != nil {
		http.Error(c, err, 900)
		return
	}
	log.Println("admin", admin)
	http.Success(c, admin)
}
