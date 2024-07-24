package middlewares

import (
	"aisecurity/expects"
	"aisecurity/structs/types"
	"aisecurity/utils"
	"aisecurity/utils/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"os"
	"strings"
)

func IsAdminAuthorized(checkPermission bool) gin.HandlerFunc {
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
			if strings.Contains(err.Error(), "token is expired") {
				http.Error(c, expects.NewNotLoggedIn())
				return
			}
			http.Error(c, err)
			return
		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				data := claims["data"].(map[string]any)
				var data2 *types.JWTData
				err := gconv.Struct(data, &data2)
				if err != nil {
					http.Error(c, err)
					return
				}
				admin := data2.Admin
				c.Set("admin_id", admin.ID)
				c.Set("jwt_data", data2)
				if checkPermission && admin.ID != 1 {
					isAuthorized, err := checkAdminPermission(c, admin.AccessIDs)
					if err != nil {
						http.Error(c, err)
						return
					}
					if !isAuthorized {
						http.Error(c, utils.ErrorWithStack(expects.NewUnauthorized(fmt.Sprintf("unauthorized access to %s", c.Request.URL.Path))))
						return
					}
				}
			} else {
				http.Error(c, err)
				return
			}
		}

		c.Next() // Proceed to the next handler if the user has the permission
	}
}

func checkAdminPermission(c *gin.Context, accessIDs []string) (bool, error) {
	pathSegments := strings.Split(c.Request.URL.Path, "/")
	if len(pathSegments) < 3 {
		return false, utils.ErrorWithStack(fmt.Errorf("invalid request path"))
	}
	lastTwoSegments := pathSegments[len(pathSegments)-2] + "/" + pathSegments[len(pathSegments)-1]
	isAuthorized := false
	for _, accessID := range accessIDs {
		if accessID == lastTwoSegments {
			isAuthorized = true
			break
		}
	}

	if !isAuthorized {
		return false, nil
	}
	return true, nil
}
