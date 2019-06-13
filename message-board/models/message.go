package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/validation"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"message-board/pkg/setting"
	"strconv"
)

var (
	db *sql.DB
	err error
)
func initDatabase() {
	var (
		dbType, dbName, user, password, host string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Println("Fail to get section 'database': ", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()

	db, err = sql.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName))

	if err != nil {
		log.Println(err)
	}
}

type User struct {
	UserId int `json:"user_id"`
	UserPass int `json:"user_pass"`
}

type Message struct {
	MId int `json:"m_id"`
	UserId int `json:"user_id"`
	MCont string `json:"m_cont"`
}

func Register(userIdStr, userPassStr string) (err error) {
	initDatabase()
	defer db.Close()
	userId,err := strconv.Atoi(userIdStr)
	userPass,err := strconv.Atoi(userPassStr)
	valid := validation.Validation{}
	valid.Required(userId, "user_id").Message("user_id不能为空")
	valid.Range(userId, 0,10000, "user_id").Message("user_id取值范围为0-10000")
	valid.Required(userPass, "user_pass").Message("user_pass不能为空")
	valid.Range(userId, 1000,999999, "user_id").Message("user_pass需要为4到6位的数字")
	//valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	//valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	if valid.HasErrors() {
		fmt.Println(valid.ErrorsMap)
		return fmt.Errorf("params form wrong")
	}

	stmt, err := db.Prepare("INSERT INTO user (user_id,user_pass) VALUES(?,?)")
	_, err = stmt.Exec(userId, userPass)
	return
}

func Login(userIdStr, userPassStr string) (err error) {
	initDatabase()
	defer db.Close()
	userId,err := strconv.Atoi(userIdStr)
	userPass,err := strconv.Atoi(userPassStr)
	var rightPass int
	err = db.QueryRow("select user_pass from user where user_id = ?", userId).Scan(&rightPass)
	if err != nil {
		if err == sql.ErrNoRows {  //如果未查询到对应字段则...
			return fmt.Errorf("not exist")
		} else {
			return

	}}
	if userPass == rightPass{
		return
	}
	return fmt.Errorf("wrong password")
}

func GetAllMessages() (messages []Message, err error) {
	initDatabase()
	defer db.Close()
	rows, err := db.Query("select * from message")
	defer rows.Close()
	for rows.Next() {  //next需要与scan配合完成读取，取第一行也要先next
		message := Message{}
		err = rows.Scan(&message.MId, &message.UserId, &message.MCont)
		messages = append(messages, message)
	}
	err = rows.Err()  //返回迭代过程中出现的错误
	return
}

func AddMessage(userId int, mCont string) (mId string, err error) {
	initDatabase()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO message(user_id,m_cont) VALUES(?,?)")
	res, err := stmt.Exec(userId, mCont)
	mIdInt, err := res.LastInsertId()
	return strconv.Itoa(int(mIdInt)), err
}

func DeleteMessage(mIdStr string) (err error) {
	initDatabase()
	defer db.Close()
	mId,err := strconv.Atoi(mIdStr)
	stmt, err := db.Prepare("delete from message where m_id=?")
	fmt.Println(mId)
	_, err = stmt.Exec(mId)
	return
}
