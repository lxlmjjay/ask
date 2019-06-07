package models

import (
"database/sql"
"fmt"
_ "github.com/go-sql-driver/mysql"
	"log"
	"message-board/pkg/setting"
	"strconv"
)

var (
	db *sql.DB
)
func initDatabase() {
	var (
		err error
		dbType, dbName, user, password, host string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
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

func Register(userId, userPass int) (err error) {
	initDatabase()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO user(user_id,user_pass) VALUES(?,?)")
	_, err = stmt.Exec(userId, userPass)
	return
}

func Login(userIdStr, userPassStr string) (err error) {
	initDatabase()
	defer db.Close()
	userId,_ := strconv.Atoi(userIdStr)
	userPass,_ := strconv.Atoi(userPassStr)
	var rightPass int
	err = db.QueryRow("select user_pass from user where user_id = ?", userId).Scan(&rightPass)
	fmt.Println(userPass, rightPass)
	if err != nil {
		if err == sql.ErrNoRows {  //如果未查询到对应字段则...
			return fmt.Errorf("not exist")
		} else {
			return
		}
	}
	if userPass == rightPass{
		fmt.Println("right")
		return
	}
	return fmt.Errorf("wrong password")
}

func GetAllMessages() (ms []Message, err error) {
	initDatabase()
	defer db.Close()
	rows, err := db.Query("select * from message")
	defer rows.Close()
	for rows.Next() {  //next需要与scan配合完成读取，取第一行也要先next
		m := Message{}
		err = rows.Scan(&m.MId, &m.UserId, &m.MCont)
		ms = append(ms, m)
	}
	err = rows.Err()  //返回迭代过程中出现的错误
	return
}

func AddMessage(userIdStr, userPassStr, mCont string) (mId string, err error) {
	initDatabase()
	defer db.Close()
	userId,_ := strconv.Atoi(userIdStr)
	strconv.Atoi(userPassStr)  //需添加验证密码
	stmt, err := db.Prepare("INSERT INTO message(user_id,m_cont) VALUES(?,?)")
	res, err := stmt.Exec(userId, mCont)
	mIdInt, err := res.LastInsertId()
	return strconv.Itoa(int(mIdInt)), err
}

func DeleteMessage(userIdStr, userPassStr, mIdStr string) (err error) {
	initDatabase()
	defer db.Close()
	strconv.Atoi(userIdStr)
	strconv.Atoi(userPassStr)  //需添加验证密码
	mId,_ := strconv.Atoi(mIdStr)
	stmt, err := db.Prepare("delete from message where m_id=?")
	fmt.Println(mId)
	_, err = stmt.Exec(mId)
	return
}
