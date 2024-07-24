package dashboard

import (
	"aisecurity/handlers"
	"aisecurity/services"
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/structs/filters"
	"aisecurity/utils"
	"aisecurity/utils/http"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
)

// AdminHandler Toycar handles
type AdminHandler struct {
	handlers.DashboardHandler
	Service *services.AdminService
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{}
}
func (h *AdminHandler) GetService(c *gin.Context) services.IService {
	return h.Service
}
func (h *AdminHandler) GetFilter(c *gin.Context) structs.IFilter {
	return h.Filter
}
func (h *AdminHandler) GetEntity(c *gin.Context) structs.IEntity {
	return h.Entity
}
func (h *AdminHandler) SetRequestContext(c *gin.Context, childHandler handlers.IHandler) {
	h.Service = services.NewAdminService(c)
	h.Service.Ctx = c
	h.Filter = &filters.Admin{}
	h.Entity = &entities.Admin{}
	h.DashboardHandler.SetRequestContext(c, h)
}

func (h *AdminHandler) GetInviteQRCode(c *gin.Context) {
	ticket, _, err := utils.GetQRCode("create-admin_1_" + gconv.String(h.Service.GetCurrentAdminID()))
	if err != nil {
		http.Error(c, utils.ErrorWithStack(err), 2000)
		return
	}
	http.Success(c, "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket="+ticket)
}
