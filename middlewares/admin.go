package middlewares

import (
	"aisecurity/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"os"
	"strings"
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
		//s := sessions.Default(c)
		//adminId := s.Get("admin_id")
		//if adminId == nil || adminId.(int) == 0 {
		//	http.Error(c, expects.NewNotLoggedIn())
		//	return
		//}

		token, err := jwt.Parse(strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1), func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				utils.Logger.Error("failed validating token", zap.Error(fmt.Errorf("unexpected signing method: %v", token.Header["alg"])))
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SESSION_KEY")), nil
		})
		if err != nil {
			//http.Error(c, err)
			//return
		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				data := claims["data"].(map[string]any)
				c.Set("admin_id", gconv.Int(data["admin"].(map[string]any)["id"]))
				c.Set("jwt_data", data)
			} else {
				//http.Error(c, err)
				//return
			}
		}

		c.Next() // Proceed to the next handler if the user has the permission
	}
}
