package db

import (
	"github.com/ozline/tiktok/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	gormopentracing "gorm.io/plugin/opentracing"

	"github.com/ozline/tiktok/pkg/utils/snowflake"
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

	DB = DB.Table(constants.VideoTableName) // choose a table

	// Generate a snowflake object
	Sf, err = snowflake.NewSnowflake(constants.SnowflakeDatacenterID, constants.SnowflakeWorkerID)

	if err != nil {
		panic(err)
	}
}
