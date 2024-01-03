package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuditMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		session := sessions.Default(c)
		// Store request-related information
		c.Set("user_agent", c.Request.UserAgent())
		c.Set("referrer_url", c.Request.Referer())
		c.Set("resource_accessed", c.Request.URL.Path)
		c.Set("http_method", c.Request.Method)
		c.Set("client_ip", c.Request.RemoteAddr)
		c.Set("client_ip2", c.ClientIP())
		c.Set("admin_id", session.Get("admin_id"))

		// Continue processing the request
		c.Next()
	}
}
