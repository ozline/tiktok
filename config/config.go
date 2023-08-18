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
	OSS.AccessKeyID = viper.GetString("oss.accessKey-id")
	OSS.AccessKeySecret = viper.GetString("oss.accessKey-secret")
	OSS.MainDirectory = viper.GetString("oss.main-directory")
	OSS.DomainName = viper.GetString("oss.domain-name")
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
