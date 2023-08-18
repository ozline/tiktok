package utils

import (
	"net"
	"strings"

	"github.com/cloudwego/kitex/pkg/klog"
	config "github.com/ozline/tiktok/config"
)

func GetMysqlDSN() string {
	if config.Mysql == nil {
		klog.Fatal("config not found")
	}

	dsn := strings.Join([]string{config.Mysql.Username, ":", config.Mysql.Password, "@tcp(", config.Mysql.Addr, ")/", config.Mysql.Database, "?charset=" + config.Mysql.Charset + "&parseTime=true"}, "")

	return dsn
}

func GetMQUrl() string {
	if config.RabbitMQ == nil {
		klog.Fatal("config not found")
	}

	url := strings.Join([]string{"amqp://", config.RabbitMQ.Username, ":", config.RabbitMQ.Password, "@", config.RabbitMQ.Addr, "/"}, "")

	return url
}

func AddrCheck(addr string) bool {
	l, err := net.Listen("tcp", addr)

	if err != nil {
		return false
	}

	l.Close()

	return true
}
