package db

import (
	"github.com/go-redis/redis/v8"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB      *gorm.DB
	SF      *utils.Snowflake
	RedisDB *redis.Client
)

func Init() {
	var err error

	DB, err = gorm.Open(mysql.Open(config.Etcd.Addr),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,                                // 禁用默认事务
			Logger:                 logger.Default.LogMode(logger.Info), // 设置日志模式
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单数表名
			},
		})

	// TODO: 加入一些其他特性

	if err != nil {
		panic(err)
	}

	sqlDB, err := DB.DB()

	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)       // 最大闲置连接数
	sqlDB.SetMaxOpenConns(constants.MaxConnections)     // 最大连接数
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime) // 最大可复用时间

	DB = DB.Table(constants.ChatTableName)

	if SF, err = utils.NewSnowflake(constants.SnowflakeDatacenterID, constants.SnowflakeWorkerID); err != nil {
		panic(err)
	}
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,  // no password set
		DB:       constants.ReidsDB_Chat, // use default DB
	})
	//docker run -d --privileged=true -p 6379:6379 -v /usr/local/redis/conf/redis.conf:/etc/redis/redis.conf -v /usr/local/redis/data:/data --name redis-1 redis:latest redis-server /etc/redis/redis.conf --appendonly yes
	if RedisDB == nil {
		panic(errors.New("[redis init error]"))
	}
}
