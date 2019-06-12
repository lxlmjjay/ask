package models

type Auth struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(userId, userPass int) bool {
	initDatabase()
	defer db.Close()
	var rightPass int
	err = db.QueryRow("select user_pass from user where user_id = ?", userId).Scan(&rightPass)
	if err != nil {return false}
	if userPass == rightPass{
		return true
	}else{
		return false
	}
}
