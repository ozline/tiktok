package dal

import (
    "github.com/ozline/tiktok/cmd/chat/dal/db"
    "github.com/ozline/tiktok/cmd/chat/dal/mq"
)

func Init() {
	db.Init()
    mq.InitRabbitMQ()
}
