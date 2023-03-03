package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	chat "github.com/ozline/tiktok/kitex_gen/tiktok/chat/tiktokchatservice"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/tracer"
	"github.com/ozline/tiktok/services/video/dal"
)

func Init() {
	dal.Init()
	tracer.InitJaeger(constants.CommentServiceName)
}

func main() {
	// Etcd Register
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdEndpoints})

	if err != nil {
		panic(err)
	}

	// Start Service

	addr, _ := net.ResolveTCPAddr("tcp", constants.ChatServiceListenAddress)
	svr := chat.NewServer(
		new(TiktokChatServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.ChatServiceName,
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
		log.Println(err.Error())
	}
}
