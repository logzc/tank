package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"log"
	"tank/rest/user"
	"tank/rest/index"
)

func main() {
	http.HandleFunc("/", index.Main)
	http.HandleFunc("/user/login", user.Login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
