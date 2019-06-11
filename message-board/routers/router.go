package routers

import (
	"github.com/gin-gonic/gin"
	"message-board/pkg/setting"
	"message-board/routers/api/v1"
	"net/http"
	"fmt"
)
//跨域中间件
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		fmt.Println(method)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors())  //允许跨域
	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/register", v1.Register)
		apiv1.POST("/login", v1.Login)
		apiv1.GET("/messages", v1.GetAllMessages)
		apiv1.POST("/messages", v1.AddMessage)
		apiv1.DELETE("/messages", v1.DeleteMessage)
	}

	return r
}
