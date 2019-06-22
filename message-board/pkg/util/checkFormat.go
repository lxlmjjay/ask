package util

import "github.com/astaxie/beego/validation"

func CheckAccountFormat(username, password string) bool {
	valid := validation.Validation{}
	valid.MinSize(username, 6, "username")
	valid.MaxSize(username, 50, "username")
	valid.MinSize(password, 6, "password")
	valid.MaxSize(password, 50, "password")
	return !valid.HasErrors()
}

func CheckMessageFormat(title, content string) bool {
	valid := validation.Validation{}
	valid.MaxSize(title, 100, "title")
	valid.MaxSize(content, 2000000000, "content")
	return !valid.HasErrors()
}
