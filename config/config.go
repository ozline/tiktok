package config

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
	"time"
)

var (
	Server        *server
	Mysql         *mySQL
	Snowflake     *snowflake
	Service       *service
	Etcd          *etcd
	RabbitMQ      *rabbitMQ
	Redis         *redis
	OSS           *oss
	Elasticsearch *elasticsearch

	runtime_viper = viper.New()
)

func Init(path string, service string) {
	runtime_viper.SetConfigType("yaml")
	runtime_viper.AddConfigPath(path)
	klog.Infof("config path: %v\n", path)

	if err := runtime_viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			klog.Fatal("could not find config files")
		} else {
			klog.Fatal("read config error")
		}
		klog.Fatal(err)
	}

	configMapping(service)

	// 持续监听配置
	runtime_viper.OnConfigChange(func(e fsnotify.Event) {
		klog.Infof("config file changed: %v\n", e.String())
	})
	runtime_viper.WatchConfig()
}

func InitRemote(endpoint string, path string, service string) {
	runtime_viper.AddRemoteProvider("etcd", endpoint, path)
	runtime_viper.SetConfigType("yaml")

	klog.Infof("config endpoint: %v\n", endpoint)

	if err := runtime_viper.ReadRemoteConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			klog.Fatal("could not find config files")
		} else {
			klog.Fatal("read config error")
		}
		klog.Fatal(err)
	}

	configMapping(service)

	// 持续监听配置
	go func() {
		for {
			time.Sleep(time.Second * 5)
			err := runtime_viper.WatchRemoteConfig()
			if err != nil {
				klog.Fatal("unable to read the remote config: %v", err)
				continue
			}
			configMapping(service)
		}
	}()
	runtime_viper.OnConfigChange(func(e fsnotify.Event) {
		klog.Infof("config file changed: %v\n", e.String())
	})
}

func configMapping(srv string) {
	c := new(config)
	if err := runtime_viper.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
	Snowflake = &c.Snowflake

	Server = &c.Server
	Server.Secret = []byte(runtime_viper.GetString("server.jwt-secret"))

	Etcd = &c.Etcd
	Mysql = &c.MySQL
	RabbitMQ = &c.RabbitMQ
	Redis = &c.Redis
	OSS = &c.OSS
	Elasticsearch = &c.Elasticsearch
	Service = GetService(srv)
}

func GetService(srvname string) *service {
	addrlist := runtime_viper.GetStringSlice("services." + srvname + ".addr")

	return &service{
		Name:     runtime_viper.GetString("services." + srvname + ".name"),
		AddrList: addrlist,
		LB:       runtime_viper.GetBool("services." + srvname + ".load-balance"),
	}
}

// Init any essentials for ci testing
func InitForTest() {
	Snowflake = &snowflake{
		WorkerID:      0,
		DatancenterID: 0,
	}

	Server = &server{
		Version: "1.0",
		Name:    "tiktok",
		Secret:  []byte("MTAxNTkwMTg1Mw=="),
	}

	Etcd = &etcd{
		Addr: "127.0.0.1:2379",
	}

	Mysql = &mySQL{
		Addr:     "127.0.0.1:3306",
		Database: "tiktok",
		Username: "tiktok",
		Password: "tiktok",
		Charset:  "utf8mb4",
	}

	RabbitMQ = &rabbitMQ{
		Addr:     "127.0.0.1:5672",
		Username: "tiktok",
		Password: "tiktok",
	}

	Redis = &redis{
		Addr:     "127.0.0.1:6379",
		Password: "tiktok",
	}
}
