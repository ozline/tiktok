package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/elastic/go-elasticsearch"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/ozline/tiktok/cmd/user/dal"
	"github.com/ozline/tiktok/cmd/user/rpc"
	"github.com/ozline/tiktok/config"
	user "github.com/ozline/tiktok/kitex_gen/user/userservice"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/eslogrus"
	"github.com/ozline/tiktok/pkg/tracer"
	"github.com/ozline/tiktok/pkg/utils"
	"github.com/sirupsen/logrus"
)

var (
	path       *string
	listenAddr string // listen port

	EsClient *elasticsearch.Client
)

func Init() {
	// config init
	path = flag.String("config", "./config", "config path")
	flag.Parse()
	config.Init(*path, constants.UserServiceName)

	// others
	dal.Init()
	tracer.InitJaeger(constants.UserServiceName)

	rpc.Init()

	EsInit()
	klog.SetLevel(klog.LevelDebug)
	klog.SetLogger(kitexlogrus.NewLogger(kitexlogrus.WithHook(EsHookLog())))
}

func EsHookLog() *eslogrus.ElasticHook {
	hook, err := eslogrus.NewElasticHook(EsClient, config.Elasticsearch.Host, logrus.DebugLevel, constants.UserServiceName)
	if err != nil {
		panic(err)
	}

	return hook
}

func EsInit() {
	esConn := fmt.Sprintf("http://%s", config.Elasticsearch.Addr)
	cfg := elasticsearch.Config{
		Addresses: []string{esConn},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	EsClient = client
}

func main() {
	Init() // 做一些中间件的初始化

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
