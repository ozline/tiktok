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

func CheckUser(username string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	if user.UserId > 0 {
		return 1 //1代表用户已存在
	}
	return 0 //0代表用户不存在
}

func AddUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return 1
	}
	return 0
}

func SelecUser(username string) int64 {
	var user User
	db.Where("username = ?", username).First(&user)
	return user.UserId
}
