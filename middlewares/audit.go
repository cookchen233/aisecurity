package middlewares

import (
	"github.com/gin-gonic/gin"
)

func DatabaseAudit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		// Store request-related information
		c.Set("is_audit", true)
		c.Set("user_agent", c.Request.UserAgent())
		c.Set("referrer_url", c.Request.Referer())
		c.Set("resource_accessed", c.Request.URL.Path)
		c.Set("http_method", c.Request.Method)
		c.Set("client_ip", c.Request.RemoteAddr)
		c.Set("client_ip2", c.ClientIP())

		// Continue processing the request
		c.Next()
	}
}
