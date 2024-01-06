package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckAdminPermission(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to determine if the user has the requiredPermission.
		// This might involve checking the user's session, a database, etc.

		c.Next() // Proceed to the next handler if the user has the permission
	}
}

func IsAdminAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)
		adminId := s.Get("admin_id")
		if adminId == nil || adminId.(int) == 0 {
			//http2.Error(c, errors.Wrap(fmt.Errorf(""), "请先登录"), 901)
			//return
		}
		c.Next() // Proceed to the next handler if the user has the permission
	}
}
