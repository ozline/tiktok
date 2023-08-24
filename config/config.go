package config

import (
	"log"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
)

var (
	Server    *server
	Mysql     *mySQL
	Snowflake *snowflake
	Service   *service
	Etcd      *etcd
	RabbitMQ  *rabbitMQ
	Redis     *redis
	OSS       *oss
)

func Init(path string, service string) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	// use etcd for config save
	// viper.AddRemoteProvider("etcd", "http://127.0.0.1:2379", "/config/config.yaml")

	klog.Infof("config path: %v\n", path)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			klog.Fatal("could not find config files")
		} else {
			klog.Fatal("read config error")
		}
		klog.Fatal(err)
	}

	configMapping(service)
}

func configMapping(srv string) {
	c := new(config)
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
	Snowflake = &c.Snowflake

	Server = &c.Server
	Server.Secret = []byte(viper.GetString("server.jwt-secret"))

	Etcd = &c.Etcd
	Mysql = &c.MySQL
	RabbitMQ = &c.RabbitMQ
	Redis = &c.Redis
	OSS = &c.OSS
	Service = GetService(srv)
}

func GetService(srvname string) *service {
	addrlist := viper.GetStringSlice("services." + srvname + ".addr")

	return &service{
		Name:     viper.GetString("services." + srvname + ".name"),
		AddrList: addrlist,
		LB:       viper.GetBool("services." + srvname + ".load-balance"),
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
