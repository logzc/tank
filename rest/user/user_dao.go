package user

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"tank/config"
)

func FindById(id uint) (*User) {
	db, err := gorm.Open("mysql", config.MysqlUrl)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Read
	var user User
	db.First(&user, id)

	return &user

}
