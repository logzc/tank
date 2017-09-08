package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"log"
	"tank/rest/user"
	"tank/rest/index"
	"fmt"
	"tank/rest/file"
)

func main() {
	http.HandleFunc("/", index.Main)
	http.HandleFunc("/user/login", user.Login)
	http.HandleFunc("/file/index", file.Index)
	http.HandleFunc("/file/upload", file.Upload)

	fmt.Println("App started at http://localhost:9090")

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
