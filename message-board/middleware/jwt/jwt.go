package jwt

import (
	"message-board/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"message-board/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "用户未登录", "msg": "请登录后再访问该内容访问该内容"})
			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "token解析错误", "msg": "请重新登录后再访问该内容访问该内容"})
			c.Abort()
			return
		}

		if claims.ExpiresAt < time.Now().Unix() {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "token超时", "msg": "请重新登录后再访问该内容访问该内容"})
			c.Abort()
			return
		}

		if !models.CheckAuth(claims.Username, claims.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "token解析错误", "msg": "请重新登录后再访问该内容访问该内容"})
			c.Abort()
			return
		}
		c.Set("username", claims.Username)

		c.Next()
	}
}
