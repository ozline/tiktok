package constants

import "time"

const (
	// auth
	JWTValue = "MTAxNTkwMTg1Mw=="
	StartID  = 10000

	// Viper config
	// etcdAddress = "http://127.0.0.1:2379"

	// redis
	ReidsDB_Chat       = 1
	RedisDBFollow      = 2
	RedisDBInteraction = 3
	RedisDBVideo       = 4
	CommentExpiredTime = 1 * time.Hour
	LikeExpiredTime    = 1 * time.Hour
	NoDataExpiredTime  = 1 * time.Minute
	LockTime           = 1 * time.Second
	LockWaitTime       = 5 * time.Millisecond
	MaxRetryTimes      = 3
	UserLikeKey        = "user:like"
	VideoLikeCountKey  = "video:like:count"
	CountKey           = "count"
	CommentKey         = "comment"
	CommentNXKey       = "commentNX"
	CountNXKey         = "countNX"
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
	MaxGoroutines   = 10
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
	AddComment    = 1
	DeleteComment = 2
	Like          = 1
	Dislike       = 2

	// follow type
	FollowAction   = 1
	UnFollowAction = 2

	// follow limit
	Interval         = 1 * time.Second
	ActionRate       = 100
	FollowListRate   = 200
	FollowerListRate = 200
	FriendListRate   = 200
)
