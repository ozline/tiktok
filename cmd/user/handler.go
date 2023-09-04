package main

import (
	"context"

	"github.com/ozline/tiktok/cmd/user/pack"
	"github.com/ozline/tiktok/cmd/user/service"
	user "github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = new(user.RegisterResponse)

	// 对传入的数据做判断
	if len(req.Username) == 0 || len(req.Username) > 255 || len(req.Password) == 0 || len(req.Password) > 255 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}

	// 发给业务层
	userResp, err := service.NewUserService(ctx).CreateUser(req)

	// 包装返回值
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	token, err := utils.CreateToken(userResp.Id)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.UserId = userResp.Id
	resp.Token = token
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)

	if len(req.Username) == 0 || len(req.Username) > 255 || len(req.Password) == 0 || len(req.Password) > 255 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}

	userResp, err := service.NewUserService(ctx).CheckUser(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	token, err := utils.CreateToken(userResp.Id)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.User = pack.User(userResp)
	resp.Token = token
	return
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, req *user.InfoRequest) (resp *user.InfoResponse, err error) {
	resp = new(user.InfoResponse)

	if req.UserId < 10000 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}

	// 校验token
	if _, err := utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}

	userResp, err := service.NewUserService(ctx).GetUser(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.User = userResp
	return
}
