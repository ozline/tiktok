package rpc

import (
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/services/auth/kitex_gen/tiktok/auth/tiktokauthservice"
	"github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user/tiktokuserservice"
	"github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video/tiktokvideoservice"
)

var (
	userClient  tiktokuserservice.Client
	authClient  tiktokauthservice.Client
	videoClient tiktokvideoservice.Client
)

func Init() {
	initAuthRPC()
	initUserRPC()
	initVideoRPC()
}

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
}

func initUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdEndpoints})

	if err != nil {
		panic(err)
	}

	c, err := tiktokuserservice.NewClient(
		constants.UserServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}

	userClient = c
}

func initVideoRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdEndpoints})

	if err != nil {
		panic(err)
	}

	c, err := tiktokvideoservice.NewClient(
		constants.VideoServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}

	videoClient = c
}
