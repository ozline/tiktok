package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	user "github.com/ozline/tiktok/kitex_gen/tiktok/user/tiktokuserservice"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/services/user/configs"

	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"

	trace "github.com/kitex-contrib/tracer-opentracing"
	tracer "github.com/ozline/tiktok/pkg/tracer"
)

func main() {

	// Jaeger
	tracer.InitJaeger(constants.UserServiceName)

	// 数据库
	configs.InitDB()

	// 日志记录
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelDebug)

	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdEndpoints})

	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", constants.UserServiceListenAddress)

	if err != nil {
		panic(err)
	}

	svr := user.NewServer(
		new(TiktokUserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.UserServiceName,
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
