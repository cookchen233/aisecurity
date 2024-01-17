package handlers

import (
	"aisecurity/expects"
	"aisecurity/utils"
	"aisecurity/utils/http"
	"github.com/gin-gonic/gin"
)

type IIPCHandler interface {
	IHandler
	ReportEvent(c *gin.Context)
}

type IPCHandler struct {
	Handler
}

func (h *IPCHandler) SetRequestContext(c *gin.Context, childHandler IHandler) {
	h.Service = childHandler.GetService(c)
	h.Filter = childHandler.GetFilter(c)
	h.Entity = childHandler.GetEntity(c)
	h.Handler.SetRequestContext(c, childHandler)
}

func (h *IPCHandler) ReportEvent(c *gin.Context) {
	http.Error(c, utils.ErrorWithStack(expects.NewNotImplementedMethod()), 500)
}
