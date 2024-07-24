package dashboard

import (
	"aisecurity/ent/dao"
	"aisecurity/expects"
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
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
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
func (h *IndexHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewAdminService(c)
	h.Service.Ctx = c
	h.Filter = &filters.Admin{}
	h.Entity = &entities.Admin{}
	h.DashboardHandler.SetRequestContext(c, h)
}

func (h *IndexHandler) CreateSuperAdmin(c *gin.Context) {
	saved, err := h.Service.CreateSuperAdmin()
	if err != nil {
		http.Error(c, utils.ErrorWithStack(err), 2000)
		return
	}
	http.Success(c, saved)
}

func (h *IndexHandler) Login(c *gin.Context) {
	var p posts.DashboardLogin
	if err := c.BindJSON(&p); err != nil {
		http.Error(c, err, 1000)
		return
	}
	if err := h.Validate(p); err != nil {
		http.Error(c, err, 1000)
		return
	}

	admin, err := h.Service.GetByUserName(p.Username)
	if err != nil {
		if dao.IsNotFound(err) {
			http.Error(c, utils.ErrorWithStack(expects.NewClientError("用户名或密码错误")))
			return
		}
		http.Error(c, err, 1000)
		return
	}

	a := admin.(*entities.Admin)
	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(p.Password)); err != nil {
		http.Error(c, utils.ErrorWithStack(expects.NewClientError("用户名或密码错误")))
		return
	}

	// Create a session and set its values
	session := sessions.Default(c)
	session.Set("admin_id", a.ID)
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
		http.Error(c, err, 1000)
		return
	}

	var accessIDs []string
	for _, v := range a.Edges.Permissions {
		accessIDs = append(accessIDs, v.AccessIds...)
	}
	var jwtAdmin = types.JWTAdmin{
		ID:          a.ID,
		DisplayName: a.Nickname,
		PhotoURL:    a.Avatar.URL,
		PhoneNumber: a.Mobile,
		AccessIDs:   accessIDs,
	}
	resp, err := h.GetJWTResponse(jwtAdmin, time.Duration(max(3600*24, persistSeconds))*time.Second)
	if err != nil {
		http.Error(c, err, 2000)
		return
	}

	http.Success(c, resp)
}

func (h *IndexHandler) GetCurrentAdmin(c *gin.Context) {
	jwtData, ex := c.Get("jwt_data")
	if !ex {
		http.Error(c, utils.ErrorWithStack(fmt.Errorf("jwt data does not exist")), 2000)
		return
	}
	j := jwtData.(*types.JWTData)

	resp, err := h.GetJWTResponse(j.Admin, gconv.Duration(j.Persist))
	if err != nil {
		http.Error(c, err, 2000)
		return
	}

	http.Success(c, resp)
}

func (h *IndexHandler) GetJWTResponse(jwtAdmin types.JWTAdmin, persist time.Duration) (*types.JWTResponse, error) {

	jwtData := types.JWTData{
		Admin:   jwtAdmin,
		Persist: persist,
	}
	// Create a new map for jwt claims
	claims := jwt.MapClaims{
		"exp":  time.Now().Add(persist).Unix(),
		"data": jwtData,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SESSION_KEY")))
	if err != nil {
		return nil, err
	}
	return &types.JWTResponse{
		AccessToken: tokenString,
		Admin:       jwtAdmin,
	}, nil
}

func (h *IndexHandler) UploadImages(c *gin.Context) {
	var basePath = "./data/uploads/images"
	var maxSize int64 = 1024 * 1024 * 3
	var allowedMimeTypes = types.NewAllowedMimeTypes([]string{
		"image/jpeg",
		"image/png",
	})
	form, _ := c.MultipartForm()
	fileHeaders := form.File["files"]
	uploadedFiles, err := h.UploadFile(c, basePath, fileHeaders, maxSize, allowedMimeTypes)
	if err != nil {
		http.Error(c, utils.ErrorWithStack(err))
		return
	}
	http.Success(c, uploadedFiles)
}
