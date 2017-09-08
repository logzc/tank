package user

import (
	"net/http"
	"tank/result"
)

func Login(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")

	if "" == email || "" == password {
		result.OutputError(w, "请输入邮箱和密码")
		return
	}

	userDao := NewUserDao()
	user := userDao.FindByEmail(email)
	if user == nil {
		result.OutputError(w, "邮箱或密码错误")
		return
	} else {
		if user.Password != password {
			result.OutputError(w, "邮箱或密码错误")
			return
		}
	}

	result.OutputSuccess(w, "登录成功")

}
