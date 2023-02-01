package main

import (
	"context"
	user "github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user"
)

// TiktokUserServiceImpl implements the last service interface defined in the IDL.
type TiktokUserServiceImpl struct{}

// Login implements the TiktokUserServiceImpl interface.
func (s *TiktokUserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// Register implements the TiktokUserServiceImpl interface.
func (s *TiktokUserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// Info implements the TiktokUserServiceImpl interface.
func (s *TiktokUserServiceImpl) Info(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}

// PingPong implements the TiktokUserServiceImpl interface.
func (s *TiktokUserServiceImpl) PingPong(ctx context.Context, req *user.Request1) (resp *user.Response1, err error) {
	// TODO: Your code here...
	return
}
