package dashboard

import (
	"aisecurity/ent/dao"
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/structs/posts"
	"aisecurity/utils/http"
	"context"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type IndexHandler struct {
	handlers.Handler
	Service *services.AdminService
}

func NewIndexHandler() *IndexHandler {
	h := &IndexHandler{}
	h.Service = services.NewAdminService()
	h.Handler.Service = h.Service
	return h
}
func (h *IndexHandler) ResetRequest(ctx context.Context) {
	h.Filter = &filters.Admin{}
	h.Handler.Filter = h.Filter
	h.Entity = &entities.Admin{}
	h.Handler.Entity = h.Entity
}

func (h *IndexHandler) Login(c *gin.Context) {
	var p posts.DashboardLogin
	if err := c.BindJSON(&p); err != nil {
		http.Error(c, err, 900)
		return
	}
	if err := h.Validate(p); err != nil {
		http.Error(c, err, 900)
		return
	}

	admin, err := h.Service.GetByUserName(p.Username)
	if err != nil {
		if dao.IsNotFound(err) {
			http.Error(c, errors.Wrap(fmt.Errorf(""), "用户名或密码错误"), 901)
			return
		}
		http.Error(c, err, 900)
		return
	}

	// Compare the provided password with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(p.Password)); err != nil {
		http.Error(c, errors.Wrap(fmt.Errorf(""), "用户名或密码错误"), 901)
		return
	}

	// Create a session and set its values
	session := sessions.Default(c)
	session.Set("authenticated", true)
	session.Set("admin_id", admin.ID)
	persistSeconds := 0
	if p.PersistDays > 0 {
		persistSeconds = max(p.PersistDays, 30) * 86400
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
