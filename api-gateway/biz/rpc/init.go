package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/ozline/tiktok/kitex_gen/tiktok/chat/tiktokchatservice"
	"github.com/ozline/tiktok/kitex_gen/tiktok/comment/tiktokcommentservice"
	"github.com/ozline/tiktok/kitex_gen/tiktok/follow/tiktokfollowservice"
	"github.com/ozline/tiktok/kitex_gen/tiktok/user/tiktokuserservice"
	"github.com/ozline/tiktok/kitex_gen/tiktok/video/tiktokvideoservice"
	"github.com/ozline/tiktok/pkg/constants"

	trace "github.com/kitex-contrib/tracer-opentracing"
)

var (
	userClient    tiktokuserservice.Client
	videoClient   tiktokvideoservice.Client
	followClient  tiktokfollowservice.Client
	commentClient tiktokcommentservice.Client
	chatClient    tiktokchatservice.Client
)

func Init() {
	initUserRPC()
	initVideoRPC()
	initChatRPC()
	initCommentRPC()
	initFollowRPC()
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
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithFailureRetry(retry.NewFailurePolicy()),
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
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}

	videoClient = c
}

func initFollowRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdEndpoints})

	if err != nil {
		panic(err)
	}

	c, err := tiktokfollowservice.NewClient(
		constants.FollowServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}

	followClient = c
}

func initCommentRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdEndpoints})

	if err != nil {
		panic(err)
	}

	c, err := tiktokcommentservice.NewClient(
		constants.CommentServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}

	commentClient = c
}

func initChatRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdEndpoints})

	if err != nil {
		panic(err)
	}

	c, err := tiktokchatservice.NewClient(
		constants.ChatServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}

	chatClient = c
}
