package routers

import (
	"github.com/gin-gonic/gin"
	"message-board/pkg/setting"
	"message-board/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
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