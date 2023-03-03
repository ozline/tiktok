package model

import (
	"fmt"

	"github.com/ozline/tiktok/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	gormopentracing "gorm.io/plugin/opentracing"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN), &gorm.Config{
		// gorm日志模式：silent
		// Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		// SkipDefaultTransaction: true,
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("数据库连接成功！")
	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(constants.MaxOpenConns)

	// SetConnMaxLifeTime 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime)

	if err = db.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	db = db.Table(constants.CommentTableName)
}

func DB() *gorm.DB {
	return db
}
