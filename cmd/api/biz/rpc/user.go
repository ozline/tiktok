package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/kitex_gen/user/userservice"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/middleware"

	trace "github.com/kitex-contrib/tracer-opentracing"
)

func InitUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})

	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer()),
	)

	if err != nil {
		panic(err)
	}

	userClient = c
}

func UserRegister(ctx context.Context, req *user.RegisterRequest) (int64, string, error) {
	resp, err := userClient.Register(ctx, req)

	if err != nil {
		return -1, "", err
	}

	if resp.Base.Code != errno.SuccessCode {
		return -1, "", errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp.UserId, resp.Token, nil
}

func UserLogin(ctx context.Context, req *user.LoginRequest) (int64, string, error) {
	resp, err := userClient.Login(ctx, req)

	if err != nil {
		return -1, "", err
	}

	if resp.Base.Code != errno.SuccessCode {
		return -1, "", errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp.User.Id, resp.Token, nil
}

func UserInfo(ctx context.Context, req *user.InfoRequest) (*user.User, error) {
	resp, err := userClient.Info(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp.User, nil
}
