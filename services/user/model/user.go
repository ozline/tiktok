package model

import (
	"fmt"

	"github.com/ozline/tiktok/services/user/configs"
	"gorm.io/gorm"
)

type User struct {
	ID            int64
	Username      string
	Password      string
	FollowCount   int64
	FollowerCount int64
	gorm.Model
}

// 注册检查
func CheckUser(username string) bool {
	var count int64
	configs.Db.Where("username = ?", username).Count(&count)

	return count > 0 // 存在数目则表示已经注册
}

// 注册用户
func AddUser(data *User) bool {
	err := configs.Db.Create(&data).Error
	fmt.Println(err)

	return err != nil
}

// 查询注册用户id
func SelecUser(username string) int64 {
	var user User
	configs.Db.Where("username = ?", username).First(&user)
	return user.ID
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
