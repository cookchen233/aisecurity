package handlers

import (
	"aisecurity/structs"
	"aisecurity/structs/entities"
	"aisecurity/utils"
	"aisecurity/utils/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type IDashboardHandler interface {
	IHandler
	Create(c *gin.Context)
	GetList(c *gin.Context)
	GetDetails(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type DashboardHandler struct {
	Handler
}

func (h *DashboardHandler) SetRequestContext(c *gin.Context, childHandler IHandler) {
	h.Service = childHandler.GetService(c)
	h.Filter = childHandler.GetFilter(c)
	h.Entity = childHandler.GetEntity(c)
	h.Handler.SetRequestContext(c, childHandler)
}

func (h *DashboardHandler) Create(c *gin.Context) {
	time.Sleep(time.Second * 1)
	if err := c.ShouldBindJSON(h.Handler.Entity); err != nil {
		http.Error(c, err, 900)
		return
	}
	if err := h.Validate(h.Handler.Entity); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := h.Handler.Service.Create(h.Handler.Entity)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}

func (h *DashboardHandler) Update(c *gin.Context) {
	time.Sleep(time.Second * 1)
	if err := c.ShouldBindJSON(h.Handler.Entity); err != nil {
		http.Error(c, err, 900)
		return
	}
	if err := h.Validate(h.Handler.Entity); err != nil {
		http.Error(c, err, 900)
		return
	}
	saved, err := h.Handler.Service.Update(h.Handler.Entity)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, saved)
}

func (h *DashboardHandler) GetList(c *gin.Context) {
	time.Sleep(time.Second * 1)
	var err error
	utils.Logger.Info("called default GetList")
	if err := c.ShouldBindQuery(h.Handler.Filter); err != nil {
		http.Error(c, err, 900)
		return
	}
	utils.Logger.Info("bound filter", zap.Any("filter", h.Handler.Filter))
	var resp = struct {
		Total int               `json:"total"`
		List  []structs.IEntity `json:"list"`
	}{}

	resp.Total, err = h.Handler.Service.GetTotal(h.Handler.Filter)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	utils.Logger.Info("called total", zap.Int("total", resp.Total))
	if resp.Total == 0 {
		http.Success(c, resp)
		return
	}

	resp.List, err = h.Handler.Service.GetList(h.Handler.Filter)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	utils.Logger.Info("called service.GetList", zap.Int("length", len(resp.List)))
	http.Success(c, resp)
}

func (h *DashboardHandler) GetDetails(c *gin.Context) {
	time.Sleep(time.Second * 1)

	if err := c.ShouldBindQuery(h.Handler.Filter); err != nil {
		http.Error(c, err, 900)
		return
	}
	details, err := h.Handler.Service.GetDetails(h.Handler.Filter)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, details)
}

func (h *DashboardHandler) Delete(c *gin.Context) {
	time.Sleep(time.Second * 1)
	var ent *entities.ID
	if err := c.ShouldBindJSON(&ent); err != nil {
		http.Error(c, err, 900)
		return
	}
	if err := h.Validate(ent); err != nil {
		http.Error(c, err, 900)
		return
	}
	err := h.Handler.Service.Delete(ent)
	if err != nil {
		http.Error(c, err, 1000)
		return
	}
	http.Success(c, nil)
}
