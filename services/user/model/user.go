package model

import (
	"time"
)

type User struct {
	UserId        int64
	Username      string
	Password      string
	FollowCount   int64
	FollowerCount int64
	CreateDate    time.Time
}
type Follower struct {
	Id         int64
	UserId     int64
	ToUserId   int64
	CreateDate time.Time
}

// 注册检查
func CheckUser(username string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.UserId > 0 {
		return 1 //1代表用户已存在
	}
	return 0 //0代表用户不存在
}

// 注册用户
func AddUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return 1
	}
	return 0
}

// 查询注册用户id
func SelecUser(username string) int64 {
	var user User
	db.Where("username = ?", username).First(&user)
	return user.UserId
}

// 登录检查
func LoginCheck(data *User, username string) int {
	db.Where("username = ?", username).First(&data)
	if data.UserId <= 0 {
		return 1 //1代表用户不存在
	}
	return 0
}

// 通过id查询用户信息
func GetUserById(userid int64) User {
	var user User
	db.Where("user_id = ?", userid).First(&user)
	return user
}
