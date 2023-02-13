package service

import (
	"context"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils/snowflake"
	"github.com/ozline/tiktok/services/chat/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DataBaseService struct {
	Ctx context.Context
	S   *snowflake.Snowflake
}

func NewDataBaseService(ctx context.Context) *DataBaseService {
	sf, _ := snowflake.NewSnowflake(constants.SnowflakeDatacenterID, constants.SnowflakeWorkerID)
	return &DataBaseService{
		Ctx: ctx,
		S:   sf,
	}
}
func (d *DataBaseService) ReceiveMessageMysqlHandler(message model.Message) {
	db, err := gorm.Open(sqlite.Open("receiveMessage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&message)
}

func (d *DataBaseService) SendMessageMysqlHandler(message model.Message) {
	db, err := gorm.Open(sqlite.Open("sendMessage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&message)
}
