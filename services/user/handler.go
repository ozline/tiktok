package main

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils/snowflake"
	"github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user"
	"github.com/ozline/tiktok/services/user/model"
	"github.com/ozline/tiktok/services/user/utils"
	"golang.org/x/crypto/bcrypt"
)

// TiktokUserServiceImpl implements the last service interface defined in the IDL.
type TiktokUserServiceImpl struct {
}

// 登录
func (s *TiktokUserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	//0.创建返回对象
	resp = new(user.DouyinUserLoginResponse)
	var user model.User
	username := req.Username
	password := req.Password
	//1.检查用户是否存在
	if exist := model.LoginCheck(&user, username); exist == 1 {
		resp.StatusCode = 1
		resp.StatusMsg = "该用户不存在!"
		return resp, nil
	}
	//2.用户名存在则校验密码
	result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if result != nil {
		fmt.Println("密码错误！")
		resp.StatusCode = 1
		resp.StatusMsg = "密码错误！"
		return resp, nil
	}
	//3.登陆验证通过，生成token
	token, err := utils.CreateToken(int64(user.ID))
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "Token Create failed"
		return resp, nil
	}
	//4.返回
	resp.StatusCode = 0 //0代表成功其他代表失败
	resp.StatusMsg = "登录成功！"
	resp.UserId = int64(user.ID)
	resp.Token = token
	return resp, err
}

// 注册
func (s *TiktokUserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	//0.创建返回对象
	resp = new(user.DouyinUserRegisterResponse)
	var user model.User
	username := req.Username
	password := req.Password
	//1.首先检查用户名是否已存在
	if exist := model.CheckUser(username); exist == 1 {
		resp.StatusCode = 1
		resp.StatusMsg = "该用户名已存在！"
		return resp, nil
	}
	//2.密码非对称加密
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("加密失败:", err)
	}
	encodePWD := string(hash)
	//3.用户数据赋值
	sf, err := snowflake.NewSnowflake(constants.SnowflakeWorkerID, constants.SnowflakeDatacenterID)
	if err != nil {
		glog.Error(err)
		return
	}
	id := sf.NextVal()
	user.ID = uint(id)
	user.Username = username
	user.Password = encodePWD
	//4.用户数据插入数据库
	if ok := model.AddUser(&user); ok == 1 {
		resp.StatusCode = 1
		resp.StatusMsg = "用户注册失败!"
		return resp, nil
	}
	//5.查询注册用户的id
	userid := model.SelecUser(username)
	//6.生成token
	token, err := utils.CreateToken(userid)
	if err != nil {
		resp.StatusCode = 1
		resp.StatusMsg = "Token创建失败！"
		return resp, nil
	}
	//7.注册成功
	resp.StatusCode = 0
	resp.StatusMsg = "注册成功!"
	resp.UserId = id
	resp.Token = token
	return resp, nil
}

// 用户信息
func (s *TiktokUserServiceImpl) Info(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	//0.创建返回对象
	resp = new(user.DouyinUserResponse)
	//1.获取id
	id := req.UserId
	//2.通过用户id查询对应用户
	userInfo := model.GetUserById(id)
	fmt.Println(userInfo)
	//3.返回结果
	resp.User = &user.User{
		Id:            int64(userInfo.ID),
		Name:          userInfo.Username,
		FollowCount:   userInfo.FollowCount,
		FollowerCount: userInfo.FollowerCount,
		IsFollow:      true,
	}
	fmt.Println(resp.User)
	resp.StatusCode = 0
	resp.StatusMsg = "成功获取用户信息！"
	return resp, nil
}

// 框架运行测试
func (s *TiktokUserServiceImpl) PingPong(ctx context.Context, req *user.Request1) (resp *user.Response1, err error) {
	resp = &user.Response1{}
	resp.Message = req.Message
	return
}
