package rpc

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/services/auth/kitex_gen/tiktok/auth"
	"github.com/ozline/tiktok/services/auth/kitex_gen/tiktok/auth/tiktokauthservice"
)

func initAuthRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdEndpoints})

	if err != nil {
		panic(err)
	}

	c, err := tiktokauthservice.NewClient(
		constants.AuthServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}

	authClient = c

	fmt.Println("Auth RPC connected!")
}

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
