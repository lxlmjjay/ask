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
	r.Use(gin.Recovery())
	//r.Use(cors.Cors())  //允许跨域
	gin.SetMode(setting.RunMode)
	r.GET("/",logger.Logger(), Info)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(logger.Logger())
	{
		apiv1.POST("/user", v1.Register)
		apiv1.PUT("/user",jwt.JWT(), v1.ModifyUser)
		apiv1.DELETE("/user",jwt.JWT(), v1.DeleteUser)

		apiv1.POST("/auth", v1.Login)
		apiv1.DELETE("/auth",jwt.JWT(), v1.Logout)

		apiv1.GET("/images/:image_name",jwt.JWT(), v1.GetImage)

		apiv1.GET("/messages",jwt.JWT(), v1.GetMessages)
		apiv1.POST("/messages",jwt.JWT(), v1.AddMessage)

		apiv1.GET("/messages/:message_id",jwt.JWT(), v1.GetMessageById)
		apiv1.PUT("/messages/:message_id",jwt.JWT(), v1.ModifyMessage)
		apiv1.DELETE("/messages/:message_id",jwt.JWT(), v1.DeleteMessage)
	}

	return r
}
