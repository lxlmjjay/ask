package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"message-board/pkg/setting"
	"message-board/routers"
	"net/http"
)

func main() {

	logs.SetLogger("console")
	logs.SetLevel(7)
	//logs.SetLogger(logs.AdapterFile,`{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	//f, _ := os.Open("project.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
