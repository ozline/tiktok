package main

import (
	"context"

	"github.com/ozline/tiktok/pkg/errno"
	auth "github.com/ozline/tiktok/services/auth/kitex_gen/tiktok/auth"
	"github.com/ozline/tiktok/services/auth/pack"
	"github.com/ozline/tiktok/services/auth/service"
)

// TiktokAuthServiceImpl implements the last service interface defined in the IDL.
type TiktokAuthServiceImpl struct{}

// Ping implements the TiktokAuthServiceImpl interface.
func (s *TiktokAuthServiceImpl) Ping(ctx context.Context, req *auth.PingRequest) (resp *auth.PingResponse, err error) {
	resp = new(auth.PingResponse)
	resp.Message = "Pong"
	return
}

// GetToken implements the TiktokAuthServiceImpl interface.
func (s *TiktokAuthServiceImpl) GetToken(ctx context.Context, req *auth.GetTokenRequest) (resp *auth.GetTokenResponse, err error) {
	resp = new(auth.GetTokenResponse)

	if len(req.Username) == 0 || req.UserId == 0 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}

	token, err := service.NewAuthService(ctx).GetToken(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.Token = token
	return resp, nil
}

// CheckToken implements the TiktokAuthServiceImpl interface.
func (s *TiktokAuthServiceImpl) CheckToken(ctx context.Context, req *auth.CheckTokenRequest) (resp *auth.CheckTokenResponse, err error) {
	resp = new(auth.CheckTokenResponse)

	if len(req.Token) == 0 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}
	claims, err := service.NewAuthService(ctx).CheckToken(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}

	resp.Info = &auth.Auth{
		UserId:    claims.UserId,
		Username:  claims.Username,
		ExpiresAt: claims.ExpiresAt,
		NotBefore: claims.NotBefore,
	}
	resp.Base = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
