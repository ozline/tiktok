package rpc

import (
	"context"

	"github.com/ozline/tiktok/api-gateway/biz/model/model"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user"
)

// Login returns token
func UserLogin(ctx context.Context, req *user.UserLoginRequest) (int64, string, error) {
	resp, err := userClient.Login(ctx, req)

	if err != nil {
		return -1, "", errno.ServiceInternalError
	}

	if resp.Base.Code != errno.SuccessCode {
		return -1, "", errno.NewErrNo(int64(resp.Base.Code), resp.Base.Msg)
	}

	return resp.UserId, resp.Token, nil
}

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
