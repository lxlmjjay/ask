package main

import (
"fmt"
"message-board/pkg/log"
"message-board/pkg/setting"
"message-board/routers"
"net/http"
)

func init()  {
	setting.Setup()
	log.Setup()
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
