package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/kitex_gen/tiktok/user"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils/snowflake"
	"github.com/ozline/tiktok/services/user/model"
	"github.com/ozline/tiktok/services/user/pack"
	"github.com/ozline/tiktok/services/user/service"
	"github.com/ozline/tiktok/services/user/utils"
	"golang.org/x/crypto/bcrypt"
)

// TiktokUserServiceImpl implements the last service interface defined in the IDL.
type TiktokUserServiceImpl struct {
}

// 登录
func (s *TiktokUserServiceImpl) Login(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	//0.创建返回对象
	resp = new(user.UserLoginResponse)
	var user model.User
	username := req.Username
	password := req.Password
	//1.检查用户是否存在
	if exist := model.LoginCheck(&user, username); exist == 1 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}
	//2.用户名存在则校验密码
	result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if result != nil {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}
	//3.登陆验证通过，生成token
	token, err := utils.CreateToken(user.ID)
	if err != nil {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}
	//4.返回
	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.UserId = int64(user.ID)
	resp.Token = token
	return resp, nil
}

// 注册
func (s *TiktokUserServiceImpl) Register(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	//0.创建返回对象
	resp = new(user.UserRegisterResponse)
	var user model.User
	//1.首先检查用户名是否已存在
	if exist := model.CheckUser(req.Username); exist {
		resp.Base = pack.BuildBaseResp(errno.UserExistedError)
		return resp, nil
	}
	//2.密码非对称加密
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		resp.Base = pack.BuildBaseResp(errno.ServiceInternalError)
		return resp, nil
	}
	encodePWD := string(hash)
	//3.用户数据赋值
	sf, err := snowflake.NewSnowflake(constants.SnowflakeWorkerID, constants.SnowflakeDatacenterID)
	if err != nil {
		resp.Base = pack.BuildBaseResp(errno.ServiceInternalError)
		return resp, nil
	}
	id := sf.NextVal()
	user.ID = id
	user.Username = req.Username
	user.Password = encodePWD
	//4.用户数据插入数据库
	if ok := model.AddUser(&user); ok {
		resp.Base = pack.BuildBaseResp(errno.ServiceInternalError)
		return resp, nil
	}
	//6.生成token
	token, err := utils.CreateToken(id)
	if err != nil {
		klog.Info("Token ERROR")
		resp.Base = pack.BuildBaseResp(errno.ServiceInternalError)
		return resp, nil
	}
	//7.注册成功
	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.UserId = id
	resp.Token = token
	return resp, nil
}

// 用户信息
func (s *TiktokUserServiceImpl) Info(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	//0.创建返回对象
	resp = new(user.UserResponse)
	//1.获取id
	id := req.UserId
	//2.通过用户id查询对应用户
	userInfo := model.GetUserById(id)
	//3.返回结果
	resp.User = &user.User{
		Id:            int64(userInfo.ID),
		Name:          userInfo.Username,
		FollowCount:   userInfo.FollowCount,
		FollowerCount: userInfo.FollowerCount,
		IsFollow:      true,
	}
	resp.Base = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// 获取Token
func (s *TiktokUserServiceImpl) GetToken(ctx context.Context, req *user.GetTokenRequest) (resp *user.GetTokenResponse, err error) {
	//0.创建返回对象
	resp = new(user.GetTokenResponse)
	//1.判断用户是否登录
	if len(req.Username) == 0 || req.UserId == 0 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}
	//2.登录则生成token
	token, err := service.NewAuthService(ctx).GetToken(req)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	//3.返回结果
	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.Token = token
	return resp, nil
}

// 检查Token
func (s *TiktokUserServiceImpl) CheckToken(ctx context.Context, req *user.CheckTokenRequest) (resp *user.CheckTokenResponse, err error) {
	//0.创建返回对象
	resp = new(user.CheckTokenResponse)
	//1.判断token是否存在
	if len(req.Token) == 0 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}
	//2.token存在则鉴权
	claims, err := service.NewAuthService(ctx).CheckToken(req)
	if err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}
	//3.返回结果
	resp.Info = &user.Auth{
		UserId:    claims.UserId,
		Username:  claims.Username,
		NotBefore: claims.ExpiresAt,
		ExpiresAt: claims.NotBefore,
	}
	resp.Base = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
