package user

import (
	"net/http"
	"fmt"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println(username + ":" + password)

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	if "" == username || "" == password {
		fmt.Fprintf(w, "请输入用户名和密码")
		return
	}
}
