package dashboard

import (
	"aisecurity/ent/dao"
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/structs/posts"
	"aisecurity/structs/types"
	"aisecurity/utils"
	"aisecurity/utils/http"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type IndexHandler struct {
	handlers.DashboardHandler
	Service *services.AdminService
}

func NewIndexHandler() *IndexHandler { return &IndexHandler{} }
func (h *IndexHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *IndexHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *IndexHandler) GetEntity(c *gin.Context) structs.IEntity { return h.Entity }
func (h *IndexHandler) SetRequestContext(c *gin.Context, h2 handlers.IHandler) {
	h.Service = services.NewAdminService()
	h.Service.Ctx = c
	h.Filter = &filters.Admin{}
	h.Entity = &entities.Admin{}
	h.DashboardHandler.SetRequestContext(c, h)
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
	http.Success(c, admin)
}

func (h *IndexHandler) Logout(c *gin.Context) {

}

func (h *IndexHandler) UploadImages(c *gin.Context) {
	var basePath = "./data/uploads/images"
	var maxSize int64 = 1024 * 1024 * 1
	var allowedMimeTypes = types.NewAllowedMimeTypes([]string{
		"image/jpeg",
		"image/png",
	})
	form, _ := c.MultipartForm()
	fileHeaders := form.File["files"]
	uploadedFiles, err, code := h.UploadFile(c, basePath, fileHeaders, maxSize, allowedMimeTypes)
	if err != nil {
		http.Error(c, utils.ErrorWithStack(err), code)
		return
	}
	http.Success(c, uploadedFiles)
}
