package rpc

import (
	"context"

	"github.com/ozline/tiktok/api-gateway/biz/model/model"
	"github.com/ozline/tiktok/kitex_gen/tiktok/user"
	"github.com/ozline/tiktok/pkg/errno"
)

// GetToken return token
func GetToken(ctx context.Context, req *user.GetTokenRequest) (string, error) {
	resp, err := userClient.GetToken(ctx, req)

	if err != nil {
		return "", err
	}

	if resp.Base.Code != errno.SuccessCode {
		return "", errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Token, nil
}

// CheckToken return userid or -1(exist error)
func CheckToken(ctx context.Context, req *user.CheckTokenRequest) (int64, error) {
	resp, err := userClient.CheckToken(ctx, req)

	if err != nil {
		return -1, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return -1, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Info.UserId, nil
}

// Login returns [userid, token]
func UserLogin(ctx context.Context, req *user.UserLoginRequest) (int64, string, error) {
	resp, err := userClient.Login(ctx, req) // RPC调用

	if err != nil {
		return -1, "", errno.ServiceInternalError
	}

	if resp.Base.Code != errno.SuccessCode {
		return -1, "", errno.NewErrNo(int64(resp.Base.Code), resp.Base.Msg)
	}

	return resp.UserId, resp.Token, nil
}

// Register returns [userid, token]
func UserRegister(ctx context.Context, req *user.UserRegisterRequest) (int64, string, error) {
	resp, err := userClient.Register(ctx, req)

	if err != nil {
		return -1, "", errno.ServiceInternalError
	}

	if resp.Base.Code != errno.SuccessCode {
		return -1, "", errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp.UserId, resp.Token, nil
}

// GetInfo returns [User]
func UserGetInfo(ctx context.Context, req *user.UserRequest) (*model.User, error) {
	resp, err := userClient.Info(ctx, req)

	if err != nil {
		return nil, errno.ServiceInternalError
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return &model.User{
		Id:            resp.User.Id,
		Name:          resp.User.Name,
		FollowCount:   resp.User.FollowCount,
		FollowerCount: resp.User.FollowerCount,
	}, nil
}
