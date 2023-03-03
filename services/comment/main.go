package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"

	comment "github.com/ozline/tiktok/kitex_gen/tiktok/comment/tiktokcommentservice"
	"github.com/ozline/tiktok/pkg/constants"
	_ "github.com/ozline/tiktok/services/comment/model"

	tracer "github.com/ozline/tiktok/pkg/tracer"
)

func Init() {
	tracer.InitJaeger(constants.CommentServiceName)
}

func main() {

	// etcd注册
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdEndpoints})

	if err != nil {
		panic(err)
	}

	// 开启服务
	addr, err := net.ResolveTCPAddr("tcp", constants.CommentServiceListenAddress)

	if err != nil {
		panic(err)
	}

	Init()

	svr := comment.NewServer(
		new(TiktokCommentServiceImpl), // 实现服务的结构体
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.CommentServiceName,
		}),
		server.WithServiceAddr(addr),                    // 服务地址
		server.WithRegistry(r),                          // Etcd注册
		server.WithSuite(trace.NewDefaultServerSuite()), // 链路追踪
		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnections,
			MaxQPS:         constants.MaxQPS,
		}), // 限制
		server.WithMuxTransport(), // Multiplex

	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
