package user

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"tank/config"
)

type UserDao struct {
	db *gorm.DB
}

//构造函数
func NewUserDao() *UserDao {

	var userDao = &UserDao{}

	return userDao;
}

//析构函数。
func (this *UserDao) openDb() {

	var err error = nil
	this.db, err = gorm.Open("mysql", config.MysqlUrl)

	if err != nil {
		panic("failed to connect mysql database")
	}

}

//析构函数。
func (this *UserDao) closeDb() {

	if this.db != nil {
		this.db.Close()
	}
}

//按照Id查询用户
func (this *UserDao) FindById(id uint) (*User) {

	// Read
	var user User

	this.openDb()
	this.db.First(&user, id)
	this.closeDb()

	return &user
}

//按照邮箱查询用户。
func (this *UserDao) FindByEmail(email string) (*User) {

	var user User

	this.openDb()
	this.db.Where(&User{Email: email}).First(&user)
	this.closeDb()

	return &user
}
