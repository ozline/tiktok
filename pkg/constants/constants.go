package constants

import "time"

const (
	MySQLDefaultDSN = "tiktok:tiktok@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"

	JwtSecret = "MTAxNTkwMTg1Mw=="

	// RPC
	MuxConnection  = 1
	RPCTimeout     = 3 * time.Second
	ConnectTimeout = 50 * time.Millisecond

	// 服务名
	UserServiceName    = "tiktok-user"
	CommentServiceName = "tiktok-comment"
	FollowServiceName  = "tiktok-follow"
	ChatServiceName    = "tiktok-chat"
	VideoServiceName   = "tiktok-video"

	AuthServiceName = "tiktok-auth"

	// 服务端口
	GatewayListenAddress        = "127.0.0.1:8080"
	UserServiceListenAddress    = "127.0.0.1:8888"
	CommentServiceListenAddress = "127.0.0.1:8889"
	FollowServiceListenAddress  = "127.0.0.1:8890"
	ChatServiceListenAddress    = "127.0.0.1:8891"
	VideoServiceListenAddress   = "127.0.0.1:8892"

	AuthServiceListenAddress = "127.0.0.1:8893"

	// 表格名
	UserTableName    = "user"
	ChatTableName    = "chat"
	CommentTableName = "comment"
	FollowTableName  = "follow"
	VideoTableName   = "video"

	// 雪花
	SnowflakeWorkerID     = 0
	SnowflakeDatacenterID = 0

	// Etcd
	EtcdEndpoints = "127.0.0.1:2379"

	// Limit
	MaxConnections = 1000
	MaxQPS         = 100
)
