package main

import (
	"flag"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/ozline/tiktok/cmd/video/dal"
	"github.com/ozline/tiktok/cmd/video/dal/cache"
	video "github.com/ozline/tiktok/cmd/video/kitex_gen/video/videoservice"
	"github.com/ozline/tiktok/cmd/video/rpc"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/tracer"
	"github.com/ozline/tiktok/pkg/utils"
)

var (
	path       *string
	listenAddr string // listen port
)

func Init() {
	// config init
	path = flag.String("config", "./config", "config path")
	flag.Parse()
	config.Init(*path, constants.VideoServiceName)
	cache.Init()
	dal.Init()
	tracer.InitJaeger(constants.VideoServiceName)
	rpc.Init()
}
func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})

	if err != nil {
		panic(err)
	}

	// get available port from config set
	for index, addr := range config.Service.AddrList {
		if ok := utils.AddrCheck(addr); ok {
			listenAddr = addr
			break
		}

		if index == len(config.Service.AddrList)-1 {
			klog.Fatal("not available port from config")
		}
	}

	addr, err := net.ResolveTCPAddr("tcp", listenAddr)

	if err != nil {
		panic(err)
	}

	svr := video.NewServer(
		new(VideoServiceImpl),
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
