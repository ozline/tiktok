package dal

import (
	"github.com/ozline/tiktok/cmd/chat/dal/cache"
	"github.com/ozline/tiktok/cmd/chat/dal/db"
	"github.com/ozline/tiktok/cmd/chat/dal/mq"
	"github.com/ozline/tiktok/pkg/ants"
)

func Init() {
	db.Init()
	cache.Init()
	mq.InitRabbitMQ()
	mq.InitMessageMQ()
	mq.InitChatMQ()
	ants.Init()
}
