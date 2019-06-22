package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"message-board/pkg/util"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Message struct {
	MessageId int `json:"message_id"`
	Username string `json:"username"`
	Title string `json:"title"`
	Content string `json:"content"`
	ImageUrl string `json:"image_url"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func Register(username, password string) (code int, err error, msg string) {
	if !util.CheckAccountFormat(username, password) {
		err = fmt.Errorf("账号格式错误")
		return http.StatusBadRequest, fmt.Errorf("账号格式错误"), "username和password长度需在6到50个字符之间"
	}
	stmt, _ := db.Prepare("INSERT INTO user (username,password) VALUES(?,?)")
	defer stmt.Close()
	_, err = stmt.Exec(username, password)
	if err != nil{
		return http.StatusInternalServerError, fmt.Errorf("操作失败"),"请稍后再试"
	}
	return
}

func ModifyUser(username, newUsername, newPassword string) (code int, err error, msg string) {
	if !util.CheckAccountFormat(newUsername, newPassword) {
		err = fmt.Errorf("账号格式错误")
		return http.StatusBadRequest, fmt.Errorf("账号格式错误"), "username和password长度需在6到50个字符之间"
	}
	stmt, _ := db.Prepare("update user set username=?,password=? where username=?")
	defer stmt.Close()
	_, err = stmt.Exec(newUsername, newPassword, username)
	if err != nil{
		return http.StatusInternalServerError, fmt.Errorf("操作失败"),"请稍后再试"
	}
	return
}

func DeleteUser(username string) (code int, err error, msg string) {
	stmt, _ := db.Prepare("update user set username=?,password=? where username=?")
	defer stmt.Close()
	_, err = stmt.Exec("", "", username)
	if err != nil{
		return http.StatusInternalServerError, fmt.Errorf("操作失败"),"请稍后再试"
	}
	return
}

func Login(username string, password string) (code int, err error, msg string) {
	var rightPass string
	err = db.QueryRow("select password from user where username = ?", username).Scan(&rightPass)
	if err != nil {
		if err == sql.ErrNoRows {  //如果未查询到对应字段则...
			return http.StatusBadRequest, fmt.Errorf("登录失败"), "该用户不存在"
		}
		return http.StatusInternalServerError, fmt.Errorf("操作失败"),"请稍后再试"
	}
	if password != rightPass{
		err = fmt.Errorf("登录失败")
		return http.StatusBadRequest, fmt.Errorf("登录失败"), "密码错误"
	}
	return
}

func GetMessages(pageStr, perPageStr string) (messages []Message, code int, err error, msg string) {
    if pageStr == "" {
		rows, _ := db.Query("select * from message")
		for rows.Next() { //next需要与scan配合完成读取，取第一行也要先next
			message := Message{}
			rows.Scan(&message.MessageId, &message.Username, &message.Title, &message.Content, &message.ImageUrl, &message.CreatedOn, &message.ModifiedOn)
			messages = append(messages, message)
	    }
		err = rows.Err()
		if err != nil{  //返回迭代过程中出现的错误
			return  messages, http.StatusInternalServerError, fmt.Errorf("操作失败"),"请稍后再试"
		}
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1{
		err = fmt.Errorf("请求格式不对")
		return messages, http.StatusBadRequest, fmt.Errorf("请求格式不对"),"page必须是大于0的整数"
	}
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1{
		err = fmt.Errorf("请求格式不对")
		return messages, http.StatusBadRequest, fmt.Errorf("请求格式不对"),"per_page必须是大于0的整数"
	}
	offset := (page-1) * perPage
	rows, _ := db.Query("select * from message limit ?,?", offset, perPage)
	for rows.Next() { //next需要与scan配合完成读取，取第一行也要先next
		message := Message{}
		rows.Scan(&message.MessageId, &message.Username, &message.Title, &message.Content, &message.ImageUrl, &message.CreatedOn, &message.ModifiedOn)
		messages = append(messages, message)
	}
	err = rows.Err()  //返回迭代过程中出现的错误
	if err != nil{
		return  messages, http.StatusInternalServerError, fmt.Errorf("操作失败"),"请稍后再试"
	}
	return
}

func GetMessageById(messageIdStr string) (message Message, code int, err error, msg string) {
	messageId, err := strconv.Atoi(messageIdStr)
	if err != nil || messageId < 1{
		err = fmt.Errorf("请求格式不对")
		return message, http.StatusBadRequest, fmt.Errorf("请求格式不对"),"地址messages/xxx中，xxx必须是大于0的整数"
	}
	err = db.QueryRow("select * from message where message_id = ?", messageId).Scan(&message.MessageId, &message.Username, &message.Title, &message.Content, &message.ImageUrl, &message.CreatedOn, &message.ModifiedOn)
	if err != nil {
		if err == sql.ErrNoRows {  //如果未查询到对应字段则...
		    return message, http.StatusNotFound, fmt.Errorf("该message不存在"), "查询错误或已被删除"
		}
		return  message, http.StatusInternalServerError, fmt.Errorf("操作失败"),"请稍后再试"
	}
	return
}

func AddMessage(username, title, content, imageUrl string) (message Message, code int, err error, msg string) {
	if !util.CheckMessageFormat(title, content) {
		err = fmt.Errorf("格式错误")
		return message, http.StatusBadRequest, fmt.Errorf("格式错误"), "1.title长度需要小于100个字符, 2.content长度需要小于2G个字符"
	}
	stmt, _ := db.Prepare("INSERT INTO message (username,title,content,image_url,created_on,modified_on) VALUES(?,?,?,?,?,?)")
	defer stmt.Close()
	now := int(time.Now().Unix())
	res, err := stmt.Exec(username, title, content, imageUrl, now, now)
	if err != nil {
		return message, http.StatusInternalServerError, fmt.Errorf("提交失败"), "请稍后再试"
	}
	messageId, err := res.LastInsertId()  //LastInsertId只在自增列时有效
	if err != nil {
		log.Fatal(err)
	}
	db.QueryRow("select * from message where message_id = ?", messageId).Scan(&message.MessageId, &message.Username, &message.Title, &message.Content, &message.ImageUrl, &message.CreatedOn, &message.ModifiedOn)
	return
}

func ModifyMessage(messageIdStr, title, content, imageUrl string) (message Message, code int, err error, msg string) {
	if !util.CheckMessageFormat(title, content) {
		err = fmt.Errorf("格式错误")
		return message, http.StatusBadRequest, fmt.Errorf("格式错误"), "1.title长度需要小于100个字符, 2.content长度需要小于2G个字符"
	}
	messageId, err := strconv.Atoi(messageIdStr)
	if err != nil || messageId < 1{
		err = fmt.Errorf("请求格式不对")
		return message, http.StatusBadRequest, fmt.Errorf("请求格式不对"),"地址messages/xxx中，xxx必须是大于0的整数"
	}
	stmt, _ := db.Prepare("update message set title=?,content=?,image_url=?,modified_on=? where message_id=?")
	defer stmt.Close()
	_, err = stmt.Exec(title, content, imageUrl, int(time.Now().Unix()), messageId)
	if err != nil {
		return message, http.StatusInternalServerError, fmt.Errorf("修改失败"), "请稍后再试"
	}
	db.QueryRow("select * from message where message_id = ?", messageId).Scan(&message.MessageId, &message.Username, &message.Title, &message.Content, &message.ImageUrl, &message.CreatedOn, &message.ModifiedOn)
	return

}

func DeleteMessage(messageIdStr string) (code int, err error, msg string) {
	messageId, err := strconv.Atoi(messageIdStr)
	if err != nil || messageId < 1{
		err = fmt.Errorf("请求格式不对")
		return http.StatusBadRequest, fmt.Errorf("请求格式不对"),"地址messages/xxx中，xxx必须是大于0的整数"
	}
	stmt, _ := db.Prepare("delete from message where message_id=?")
	defer stmt.Close()
	_, err = stmt.Exec(messageId)
	if err != nil{
		return  http.StatusInternalServerError, fmt.Errorf("删除失败"),"请稍后再试"
	}
	return
}
