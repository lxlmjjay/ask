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
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			} else {
				userId := claims.UserId
				userPass := claims.UserPass
				if models.CheckAuth(userId, userPass){
					c.Set("userId", userId)
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
