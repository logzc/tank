package main

import (
	_ "github.com/go-sql-driver/mysql"
	"tank/rest/user"
	"fmt"
)

func main() {
	user0 := user.FindById(1)
	fmt.Printf("email = %s", user0.Email)
}
