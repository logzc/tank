package user

import (
	"net/http"
	"fmt"
	"strings"
	"tank/result"
)

func Login(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Println(email + ":" + password)

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
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
	return

}
