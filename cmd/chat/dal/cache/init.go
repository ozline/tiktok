package cache

import (
	"errors"
	"time"

	"github.com/ozline/tiktok/config"
	redis "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	RedisDB *redis.Client
)

type Message struct {
	Id         int64
	ToUserId   int64
	FromUserId int64
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func Init() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password, // no password set
		DB:       1,                     // use default DB
	})
	//docker run -d --privileged=true -p 6379:6379 -v /usr/local/redis/conf/redis.conf:/etc/redis/redis.conf -v /usr/local/redis/data:/data --name redis-1 redis:latest redis-server /etc/redis/redis.conf --appendonly yes

	if RedisDB == nil {
		panic(errors.New("[redis init error]"))
	}
}
