package utils

import (
	"mime/multipart"
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

func IsVideoFile(header *multipart.FileHeader) bool {
	contentType := header.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "video/") {
		return true
	}

	filename := header.Filename
	extensions := []string{".mp4", ".avi", ".mkv", ".mov"} // Add more video extensions if needed
	for _, ext := range extensions {
		if strings.HasSuffix(strings.ToLower(filename), ext) {
			return true
		}
	}

	return false
}
