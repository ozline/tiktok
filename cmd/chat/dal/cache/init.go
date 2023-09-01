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
		Password: config.Redis.Password,
		DB:       1,
	})
	if RedisDB == nil {
		panic(errors.New("[redis init error]"))
	}
}
