package main

import (
"fmt"
	"message-board/middleware/logger"
	"message-board/models"
    "message-board/pkg/setting"
    "message-board/routers"
    "net/http"
)

func init()  {
	setting.Setup()
	logger.Setup()
	models.Setup()
}

func main() {

	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
