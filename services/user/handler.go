package main

import (
	"context"
	"fmt"
	"github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user"
	"github.com/ozline/tiktok/services/user/model"
	"time"
)

// TiktokUserServiceImpl implements the last service interface defined in the IDL.
type TiktokUserServiceImpl struct{}

// Login implements the TiktokUserServiceImpl interface.
func (s *TiktokUserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// 注册
func (s *TiktokUserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	var user model.User
	username := req.Username
	password := req.Password
	//1.首先检查用户名是否已存在
	if exist := model.CheckUser(username); exist == 1 {
		resp.StatusCode = 1
		*resp.StatusMsg = "用户已存在"
		return nil, err
	}
	//2.用户数据赋值
	user.Username = username
	user.Password = password
	user.FollowCount = 0
	user.FollowerCount = 0
	user.CreateDate = time.Now()
	//3.用户数据插入数据库
	if ok := model.AddUser(&user); ok == 1 {
		resp.StatusCode = 1
		*resp.StatusMsg = "用户注册失败!"
		return nil, err
	}
	//4.查询注册用户的id
	id := model.SelecUser(username)
	user.UserId = id
	fmt.Println(id)
	return nil, err
}

// Info implements the TiktokUserServiceImpl interface.
func (s *TiktokUserServiceImpl) Info(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}

// 测试一下
func (s *TiktokUserServiceImpl) PingPong(ctx context.Context, req *user.Request1) (resp *user.Response1, err error) {
	resp = &user.Response1{}
	resp.Message = req.Message
	return
}
