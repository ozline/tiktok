package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	follow "github.com/ozline/tiktok/kitex_gen/tiktok/follow/tiktokfollowservice"
	"github.com/ozline/tiktok/pkg/constants"
	tracer "github.com/ozline/tiktok/pkg/tracer"
	"github.com/ozline/tiktok/services/follow/model"
)

func Init() {
	model.Setup()
	tracer.InitJaeger(constants.FollowServiceName)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdEndpoints})

	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", constants.FollowServiceListenAddress)

	if err != nil {
		panic(err)
	}

	Init()

	svr := follow.NewServer(
		new(TiktokFollowServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.FollowServiceName,
		}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithSuite(trace.NewDefaultServerSuite()),
		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnections,
			MaxQPS:         constants.MaxQPS,
		}),
		server.WithMuxTransport(),
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
