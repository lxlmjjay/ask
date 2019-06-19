package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"message-board/pkg/setting"
	"strconv"
	"time"
)

type User struct {
	Id int `json:"id"`
	Username string `json:"username" valid: MaxSize(50)"`
	Password string `json:"password" valid:"MinSize(6); MaxSize(50)"`
}

type Message struct {
	MessageId int `json:"message_id"`
	Username int `json:"username"`
	Title string `json:"title"`
	Content string `json:"content"`
	ImageUrl string `json:"image_url"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func Register(username, password string) (err error) {
	db, err := sql.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host, setting.DatabaseSetting.Name))
	if err != nil {
		logs.Warn(err)
	}
	defer db.Close()
	valid := validation.Validation{}
	valid.Required(username, "username").Message("username不能为空")
	valid.Required(password, "password").Message("password不能为空")
	valid.MaxSize(username, 50, "username").Message("username长度需要小于50个字符")
	valid.MinSize(password, 6, "password").Message("password长度需要大于6个字符")
	valid.MaxSize(password, 50, "password").Message("password长度需要小于50个字符")

	if valid.HasErrors() {
		fmt.Println(valid.ErrorsMap)
		err = fmt.Errorf("账号格式错误")
		return
	}

	stmt, err := db.Prepare("INSERT INTO user (username,password) VALUES(?,?)")
	_, err = stmt.Exec(username, password)
	return
}

func DeleteUser(username string) (err error) {
	db, err := sql.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host, setting.DatabaseSetting.Name))
	if err != nil {
		logs.Warn(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("delete from user where username=?")
	_, err = stmt.Exec(username)

	return
}

func Login(username string, password string) (err error) {
	db, err := sql.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host, setting.DatabaseSetting.Name))
	if err != nil {
		logs.Warn(err)
	}
	defer db.Close()
	var rightPass string
	err = db.QueryRow("select password from user where username = ?", username).Scan(&rightPass)
	if err != nil {
		if err == sql.ErrNoRows {  //如果未查询到对应字段则...
			return fmt.Errorf("not exist")
		} else {
			return

	}}
	if password == rightPass{
		return
	}
	return fmt.Errorf("wrong password")
}

func GetMessages(pageStr, perPageStr string) (messages []Message, err error) {
	db, err := sql.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host, setting.DatabaseSetting.Name))
	if err != nil {
		logs.Warn(err)
		return
	}
	defer db.Close()
    if pageStr == "" {
		rows, err1 := db.Query("select * from message")
		defer rows.Close()
		if err1 != nil{
			err = err1
			return
		}
		for rows.Next() { //next需要与scan配合完成读取，取第一行也要先next
			message := Message{}
			err = rows.Scan(&message.MessageId, &message.Username, &message.Title, &message.Content, &message.ImageUrl, &message.CreatedOn, &message.ModifiedOn)
			messages = append(messages, message)
		}
		err1 = rows.Err() //返回迭代过程中出现的错误
		err = err1
		return
	}
	page,err := strconv.Atoi(pageStr)
	perPage, err := strconv.Atoi(perPageStr)
	offset := page * perPage
	rows, err := db.Query("select * from message limit ?,?", offset, perPage)
	defer rows.Close()
	for rows.Next() { //next需要与scan配合完成读取，取第一行也要先next
		message := Message{}
		err = rows.Scan(&message.MessageId, &message.Username, &message.Title, &message.Content, &message.ImageUrl, &message.CreatedOn, &message.ModifiedOn)
		messages = append(messages, message)
	}
	err = rows.Err() //返回迭代过程中出现的错误
	return
}

func GetMessageById(messageIdStr string) (message Message, err error) {
	db, err := sql.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host, setting.DatabaseSetting.Name))
	if err != nil {
		logs.Warn(err)
	}
	defer db.Close()
	messageId,err := strconv.Atoi(messageIdStr)
	err = db.QueryRow("select * from message where message_id = ?", messageId).Scan(&message)
	if err != nil {
		if err == sql.ErrNoRows {  //如果未查询到对应字段则...
		    err = fmt.Errorf("该message不存在")
			return
		} else {
			return
		}}
	return
}

func AddMessage(username, title, content, imageUrl string) (message Message, err error) {
	db, err := sql.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host, setting.DatabaseSetting.Name))
	if err != nil {
		logs.Warn(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO message(username,title,content,image_url,created_on,modified_on) VALUES(?,?,?,?,?,?)")
	res, err := stmt.Exec(username, title, content, imageUrl, int(time.Now().Unix()), int(time.Now().Unix()))
	messageId, err := res.LastInsertId()  //LastInsertId只在自增列时有效
	if err != nil {
		log.Fatal(err)
	}
	err = db.QueryRow("select * from message where message_id = ?", messageId).Scan(&message)
	if err != nil {
		err = fmt.Errorf("添加失败")
		return
		}
	return
}

func ModifyMessage(messageId, title, content, imageUrl string) (message Message, err error) {
	db, err := sql.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host, setting.DatabaseSetting.Name))
	if err != nil {
		logs.Warn(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("update message set title=?,content=?,image_url=?,modified_on=? where message_id=?")
	_, err = stmt.Exec(title, content, imageUrl, int(time.Now().Unix()), messageId)
	if err != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	err = db.QueryRow("select * from message where message_id = ?", messageId).Scan(&message)
	if err != nil {
		err = fmt.Errorf("修改失败")
		return
	}
	return

}

func DeleteMessage(messageIdStr, username string) (err error) {
	db, err := sql.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User, setting.DatabaseSetting.Password, setting.DatabaseSetting.Host, setting.DatabaseSetting.Name))
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	var rightName string
	messageId,err := strconv.Atoi(messageIdStr)
	err = db.QueryRow("select username from message where message_id = ?", messageId).Scan(&rightName)
	if err != nil {
		if err == sql.ErrNoRows {  //如果未查询到对应字段则...
			return fmt.Errorf("该message不存在")
		} else {
			return
		}}
	if username != rightName{
		return fmt.Errorf("该用户没有权限")
	}
	stmt, err := db.Prepare("delete from message where message_id=?")
	_, err = stmt.Exec(messageId)
	return
}
