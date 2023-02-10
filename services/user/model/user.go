package model

import (
	"fmt"
	"github.com/ozline/tiktok/services/user/configs"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID       uint `gorm:"primarykey"`
	CreateAt time.Time
	updateAt time.Time
	deleteAt time.Time `gorm:"index"`
}
type User struct {
	Username      string
	Password      string
	FollowCount   int64
	FollowerCount int64
	gorm.Model
}

// 注册检查
func CheckUser(username string) int {
	var user User
	configs.Db.Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return 1 //1代表用户已存在
	}
	return 0 //0代表用户不存在
}

// 注册用户
func AddUser(data *User) int {
	err := configs.Db.Create(&data).Error
	fmt.Println(err)
	if err != nil {
		return 1
	}
	return 0
}

// 查询注册用户id
func SelecUser(username string) int64 {
	var user User
	configs.Db.Where("username = ?", username).First(&user)
	return int64(user.ID)
}

// 登录检查
func LoginCheck(data *User, username string) int {
	configs.Db.Where("username = ?", username).First(&data)
	if data.ID <= 0 {
		return 1 //1代表用户不存在
	}
	return 0
}

// 通过id查询用户信息
func GetUserById(userid int64) User {
	var user User
	configs.Db.Where("id = ?", userid).First(&user)
	return user
}
