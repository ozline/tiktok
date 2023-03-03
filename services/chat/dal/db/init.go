package db

import (
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils/snowflake"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB
var Sf *snowflake.Snowflake

func Init() {
	var err error

	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})

	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	DB = DB.Table(constants.ChatTableName)

	// Generate snowflake object
	Sf, err = snowflake.NewSnowflake(constants.SnowflakeDatacenterID, constants.SnowflakeWorkerID)

	if err != nil {
		panic(err)
	}
}
