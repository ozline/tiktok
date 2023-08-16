package main

import (
	"net"
	chat "github.com/ozline/tiktok/kitex_gen/chat/messageservice"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/cloudwego/kitex/server"
	//etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/ozline/tiktok/cmd/chat/dal"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/tracer"

	trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
    dal.Init()
    tracer.InitJaeger(constants.ChatServiceName)
}

func main() {
    Init() // 做一些中间件的初始化

	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdEndpoints})

	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", constants.ChatServiceListenAddress)

	if err != nil {
		panic(err)
	}
    // ...

    svr := chat.NewServer(
        new(MessageServiceImpl),
        server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
            ServiceName: constants.ChatServiceName,
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
    
    // ...
	if err = svr.Run(); err != nil {
		panic(err)
	}
}
