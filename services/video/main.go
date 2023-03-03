package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	video "github.com/ozline/tiktok/kitex_gen/tiktok/video/tiktokvideoservice"
	"github.com/ozline/tiktok/pkg/constants"

	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/ozline/tiktok/services/video/dal"

	tracer "github.com/ozline/tiktok/pkg/tracer"
)

func Init() {
	dal.Init()
	tracer.InitJaeger(constants.VideoServiceName)
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdEndpoints})

	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", constants.VideoServiceListenAddress)

	if err != nil {
		panic(err)
	}

	Init()

	svr := video.NewServer(
		new(TiktokVideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.VideoServiceName,
		}),
		server.WithServiceAddr(addr),                    // 监听地址
		server.WithRegistry(r),                          // ETCD注册发现
		server.WithSuite(trace.NewDefaultServerSuite()), // 链路追踪
		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnections,
			MaxQPS:         constants.MaxQPS,
		}),
		// server.WithMuxTransport(), // Multiplex
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
