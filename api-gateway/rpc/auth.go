package rpc

import (
	"context"

	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/services/auth/kitex_gen/tiktok/auth"
)

func GetToken(ctx context.Context, req *auth.GetTokenRequest) (string, error) {
	resp, err := authClient.GetToken(ctx, req)

	if err != nil {
		return "", err
	}

	if resp.Base.Code != errno.SuccessCode {
		return "", errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Token, nil
}
