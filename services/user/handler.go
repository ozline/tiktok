package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ozline/tiktok/services/user/attestation"
	"github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user"
	"github.com/ozline/tiktok/services/user/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// TiktokUserServiceImpl implements the last service interface defined in the IDL.
type TiktokUserServiceImpl struct{}

// 登录
func (s *TiktokUserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	//var user model.User
	//username := req.Username
	//password := req.Password
	////1.检查用户是否存在
	//if exist := model.LoginCheck(&user, username); exist == 1 {
	//	resp.StatusCode = 1
	//	*resp.StatusMsg = "该用户不存在!"
	//	return nil, err
	//}
	//2.用户名存在校验密码
	//err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	//if err != nil {
	//	resp.StatusCode = 1
	//	*resp.StatusMsg = "密码错误！"
	//	return nil, err
	//}
	//if err != nil {
	//
	//}
	return
}

// 注册
func (s *TiktokUserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	var user model.User
	//username := req.Username
	//password := req.Password
	username := "test1"
	password := "test123"
	//1.首先检查用户名是否已存在
	if exist := model.CheckUser(username); exist == 1 {
		resp.StatusCode = 1
		*resp.StatusMsg = "该用户名已存在！"
		return nil, nil
	}
	//2.密码非对称加密
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("加密失败:", err)
	}
	encodePWD := string(hash)
	//3.用户数据赋值
	uuserId := uuid.New()
	user.UserId = int64(uuserId.ID())
	user.Username = username
	user.Password = encodePWD
	user.FollowCount = 0
	user.FollowerCount = 0
	user.CreateDate = time.Now()
	//4.用户数据插入数据库
	if ok := model.AddUser(&user); ok == 1 {
		resp.StatusCode = 1
		*resp.StatusMsg = "用户注册失败!"
		return nil, nil
	}
	//5.查询注册用户的id
	id := model.SelecUser(username)
	user.UserId = id
	fmt.Println(id)
	//6.生成token
	token, err := attestation.CreateToken(id)
	if err != nil {
		resp.StatusCode = 1
		*resp.StatusMsg = "Token创建失败!"
		return nil, nil
	}
	//7.token放进缓存
	err = model.AddToken(token, &user)
	if err != nil {
		fmt.Println("结构体异常: ", err)
		resp.StatusCode = 1
		*resp.StatusMsg = "token保存失败！"
		return nil, nil
	}
	//8.注册成功
	resp.StatusCode = 0
	*resp.StatusMsg = "注册成功!"
	resp.UserId = id
	resp.Token = token
	return nil, nil
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
