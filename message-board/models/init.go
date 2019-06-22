package models

import (
	"database/sql"
	"github.com/astaxie/beego/logs"
	"message-board/pkg/setting"
	"fmt"
)

var db *sql.DB

func Setup()  {
	var err error
	db, err = sql.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host, setting.DatabaseSetting.Name))
	if err != nil {
		logs.Warn(err)
	}
}
