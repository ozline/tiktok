package configs

import (
	"fmt"
	"os"
	"time"

	"github.com/ozline/tiktok/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	gormopentracing "gorm.io/plugin/opentracing"
)

var Db *gorm.DB
var err error

// 初始化数据库
func InitDB() {
	dsn := constants.MySQLDefaultDSN
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})

	Db = Db.Table(constants.UserTableName)
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		os.Exit(1)
	}

	fmt.Println("数据库连接成功！")
	sqlDB, _ := Db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	if err = Db.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

}
