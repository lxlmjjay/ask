package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/logs"
	"message-board/pkg/setting"
)

type Auth struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	db, err := sql.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host, setting.DatabaseSetting.Name))
	if err != nil {
		logs.Warn(err)
		return false
	}
	defer db.Close()
	var rightPass string
	err = db.QueryRow("select password from user where username = ?", username).Scan(&rightPass)
	if err != nil {
		return false
	}
	if password == rightPass{
		return true
	}else{
		return false
	}
}
