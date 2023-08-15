package constants

import "time"

const (
	MySQLDefaultDSN = "tiktok:tiktok@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=true"

	JwtSecret = "MTAxNTkwMTg1Mw=="

	//redis
	RedisAddr   = "127.0.0.1:6379"
	RedisPWD    = "tiktok"
	RedisDBChat = 1

	// RPC
	MuxConnection  = 1
	RPCTimeout     = 3 * time.Second
	ConnectTimeout = 50 * time.Millisecond

	// 服务名
	GatewayServiceName = "gateway"
	UserServiceName    = "user"
	CommentServiceName = "comment"
	FollowServiceName  = "follow"
	ChatServiceName    = "chat"
	VideoServiceName   = "video"

	// 服务端口
	GatewayListenAddress        = "127.0.0.1:8080"
	UserServiceListenAddress    = "127.0.0.1:8888"
	CommentServiceListenAddress = "127.0.0.1:8889"
	FollowServiceListenAddress  = "127.0.0.1:8890"
	ChatServiceListenAddress    = "127.0.0.1:8891"
	VideoServiceListenAddress   = "127.0.0.1:8892"

	// 表格名
	UserTableName    = "user"
	ChatTableName    = "message"
	CommentTableName = "comment"
	FollowTableName  = "follow"
	VideoTableName   = "video"

	// 雪花
	SnowflakeWorkerID     = 0
	SnowflakeDatacenterID = 0

	// Etcd
	EtcdEndpoints = "127.0.0.1:2379"

	// Limit
	MaxConnections  = 1000
	MaxQPS          = 100
	MaxVideoSize    = 300000
	MaxListLength   = 100
	MaxIdleConns    = 10
	MaxOpenConns    = 100
	ConnMaxLifetime = 10 * time.Second

	// Aliyun SDK
	// AccessKey Expires at 2023-03-03-16-00-00
	OSSEndpoint        = "files.ozline.icu"               // 默认启用域名绑定
	OSSAccessKeyID     = "LTAI5t6gqQgzCzVgdUWw6uip"       // AccessKeyID
	OSSAccessKeySecret = "diAQZbzqI6JJm53IoKSdGAudNZIbol" // AccessKeySecret
	OSSBucketName      = "ozliinex"                       // 桶名
	UplaodRoutines     = 3                                // 并发数量
	MainDirectory      = "tiktok"                         // 主目录
	PartSize           = 100 * 1024

	// Page
	PageNum  = 1
	PageSize = 10

	// ffmpeg
	FrameNum = 1

	//Action Type
	AddComment    = "1"
	DeleteComment = "2"
)
