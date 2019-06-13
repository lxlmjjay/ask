package logger

import (
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		logs.Info(c.Writer.Status(), " | ", time.Now().Sub(startTime), " | ", c.Request.RemoteAddr, " | ",c.Request.URL, " | ", c.Request.Method)
	}
}
