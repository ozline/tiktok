package db

import (
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	gormopentracing "gorm.io/plugin/opentracing"
)

type DBAction struct {
	DB *gorm.DB
}

var (
	DB *gorm.DB
	SF *utils.Snowflake
)

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(utils.GetMysqlDSN()),
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

	if err = DB.Use(gormopentracing.New()); err != nil {
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
}
func NewDBAction() *DBAction {
	return &DBAction{
		DB: DB,
	}
}
