package log

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"log"
	"message-board/pkg/setting"
)

func Setup()  {
	if setting.RunMode == "debug" {
		logs.SetLevel(logs.LevelDebug)
		if err := logs.SetLogger("console"); err != nil{
			log.Fatal(err)
		}
	}else {
		logs.SetLevel(logs.LevelInformational)
		filename :=  getLogFilePath() + "/" + getLogFileName()
		if err := logs.SetLogger(logs.AdapterFile,fmt.Sprintf(`{"filename":%s,"level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`, filename)); err != nil{
			log.Fatal(err)
		}
	}
	//filename :=  getLogFilePath() + "/" + getLogFileName()
	//logs.SetLogger(logs.AdapterFile,`{"filename":filename,"level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	//f, _ := os.Open("project.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
