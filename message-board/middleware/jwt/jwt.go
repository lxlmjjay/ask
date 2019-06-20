package jwt

import (
"message-board/models"
"net/http"
"time"

"github.com/gin-gonic/gin"

"message-board/pkg/e"
"message-board/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS
		token,err := c.Cookie("token")
		if err != nil || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code" : http.StatusUnauthorized, "err" : "用户未登录","msg" : "请登录后再访问该内容访问该内容"})
			c.Abort()
			return
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			} else {
				username := claims.Username
				password := claims.Password
				if models.CheckAuth(username, password){
					c.Set("username", username)
				}else {
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{"err" : e.GetMsg(code)})
			c.Abort()
			return
		}

		c.Next()
	}
}
