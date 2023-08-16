package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/ozline/tiktok/cmd/video/dal"
	video "github.com/ozline/tiktok/cmd/video/kitex_gen/video/videoservice"
	"github.com/ozline/tiktok/cmd/video/rpc"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/tracer"
)

func Init() {
	dal.Init()
	tracer.InitJaeger(constants.UserServiceName)
	rpc.Init()
}
func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdEndpoints})

	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", constants.VideoServiceListenAddress)

	if err != nil {
		panic(err)
	}

	svr := video.NewServer(new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.VideoServiceName,
		}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnections,
			MaxQPS:         constants.MaxQPS,
		}))

	if err = svr.Run(); err != nil {
		panic(err)
	}
}
