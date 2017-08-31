package main

import (
	_ "github.com/go-sql-driver/mysql"
	"tank/rest/user"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", user.Login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
