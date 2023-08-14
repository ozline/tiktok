package main

import (
	"net"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/ozline/tiktok/cmd/follow/dal"
	follow "github.com/ozline/tiktok/kitex_gen/follow/followservice"
	"github.com/ozline/tiktok/pkg/constants"
)

func main() {
	dal.Init()

	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdEndpoints})

	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", constants.FollowServiceListenAddress)

	if err != nil {
		panic(err)
	}

	svr := follow.NewServer(
		new(FollowServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.FollowServiceName,
		}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnections,
			MaxQPS:         constants.MaxQPS,
		}),
	)

	if err := svr.Run(); err != nil {
		panic(err)
	}
}
