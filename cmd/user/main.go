package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/ozline/tiktok/cmd/user/dal"
	user "github.com/ozline/tiktok/kitex_gen/user/userservice"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/tracer"

	trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	dal.Init()
	tracer.InitJaeger(constants.UserServiceName)
}

func main() {
	Init() // 做一些中间件的初始化

	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdEndpoints})

	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", constants.UserServiceListenAddress)

	if err != nil {
		panic(err)
	}

	svr := user.NewServer(
		new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.UserServiceName,
		}),
		server.WithMuxTransport(),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithSuite(trace.NewDefaultServerSuite()),
		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnections,
			MaxQPS:         constants.MaxQPS,
		}),
	)

	if err = svr.Run(); err != nil {
		panic(err)
	}
}
