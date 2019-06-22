package models

import (
	"fmt"
	"net/http"
	"strconv"
)

type Auth struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var rightPass string
	err := db.QueryRow("select password from user where username = ?", username).Scan(&rightPass)
	if err != nil {
		return false
	}
	if password == rightPass{
		return true
	}else{
		return false
	}
}

func CheckMessageAuth(username, messageIdStr string) (code int, err error, msg string) {
	messageId, err := strconv.Atoi(messageIdStr)
	if err != nil{
		return http.StatusBadRequest, fmt.Errorf("请求格式不对"),"请求地址messages/xxx中,xxx必须是数字"
	}
	var rightName string
	err = db.QueryRow("select username from message where message_id = ?", messageId).Scan(&rightName)
	if err != nil {
		return http.StatusNotFound, fmt.Errorf("此message不存在"), "请求错误或已被删除"
	}
	if username == rightName {
		return http.StatusOK, nil, ""
	}
	return http.StatusForbidden, fmt.Errorf("没有权限"),"不能修改其他账号的message"
}
