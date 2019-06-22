package logger

import (
    "github.com/astaxie/beego/logs"
    "github.com/gin-gonic/gin"
	"log"
	"message-board/pkg/setting"
	"time"
    "fmt"
)

func Setup()  {
	switch setting.RunMode {
	case "debug":
		logs.SetLevel(logs.LevelDebug)
		if err := logs.SetLogger("console"); err != nil{
			log.Fatal(err)
		}
	case "release":
		logs.SetLevel(logs.LevelInformational)
		filePath := setting.AppSetting.LogSavePath
		fileName := time.Now().Format(setting.AppSetting.TimeFormat) + "." + setting.AppSetting.LogFileExt
		if err := logs.SetLogger(logs.AdapterFile,fmt.Sprintf(`{"filename":%s,"level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`, filePath + fileName)); err != nil{
			log.Fatal(err)
		}
	default:
		log.Fatal("no run mode")
	}
}


func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		logs.Info(c.Writer.Status(), " | ", time.Now().Sub(startTime), " | ", c.Request.RemoteAddr, " | ",c.Request.URL, " | ", c.Request.Method)
	}
}
