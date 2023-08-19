package constants

import "time"

const (
	JWTValue = "MTAxNTkwMTg1Mw=="

	// redis
	ReidsDB_Chat       = 1
	RedisDBInteraction = 3
	RedisHashName      = "like_count"
	CommentExpiredTime = 1 * time.Hour

	// RPC
	MuxConnection  = 1
	RPCTimeout     = 3 * time.Second
	ConnectTimeout = 50 * time.Millisecond

	// service name
	APIServiceName         = "api"
	UserServiceName        = "user"
	InteractionServiceName = "interaction"
	FollowServiceName      = "follow"
	ChatServiceName        = "chat"
	VideoServiceName       = "video"

	// db table name
	UserTableName     = "user"
	ChatTableName     = "message"
	CommentTableName  = "comment"
	FavoriteTableName = "favorite"
	FollowTableName   = "follow"
	VideoTableName    = "video"

	// snowflake
	SnowflakeWorkerID     = 0
	SnowflakeDatacenterID = 0

	// limit
	MaxConnections  = 1000
	MaxQPS          = 100
	MaxVideoSize    = 300000
	MaxListLength   = 100
	MaxIdleConns    = 10
	MaxOpenConns    = 100
	ConnMaxLifetime = 10 * time.Second

	// Aliyun SDK
	UplaodRoutines = 3 // 并发数量
	PartSize       = 100 * 1024

	// page
	PageNum  = 1
	PageSize = 10

	// ffmpeg
	FrameNum = 1

	// interaction type
	AddComment    = "1"
	DeleteComment = "2"
	Like          = "1"
	Dislike       = "2"

	// follow type
	FollowAction = 1
)
