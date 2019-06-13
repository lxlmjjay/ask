package routers

import (
	"github.com/gin-gonic/gin"
	"message-board/middleware/jwt"
	"message-board/middleware/logger"
	"message-board/pkg/setting"
	"message-board/routers/api/v1"
)


func InitRouter() *gin.Engine {
	r := gin.New()
	//r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//r.Use(cors.Cors())  //允许跨域
	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(logger.Logger())
	{
		apiv1.POST("/register",v1.Register)
		apiv1.POST("/login", v1.Login)
		apiv1.GET("/messages",jwt.JWT(), v1.GetAllMessages)
		apiv1.POST("/messages",jwt.JWT(), v1.AddMessage)
		apiv1.DELETE("/messages",jwt.JWT(), v1.DeleteMessage)
	}

	return r
}
