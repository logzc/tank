package user

import (
	"time"
)

type User struct {
	Id         uint `gorm:"primary_key"`
	Email      string
	Phone      string
	Nickname   string
	Password   string
	Gender     string
	City       string
	AvatarUrl  string
	LastIp     string
	LastTime   time.Time
	Sort       int64
	CreateTime time.Time
	Deleted    bool
}

// set User's table name to be `profiles`
func (User) TableName() string {
	return "tank_user"
}

