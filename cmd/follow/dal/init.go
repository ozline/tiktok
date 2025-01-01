package dal

import (
	"github.com/ozline/tiktok/cmd/follow/dal/cache"
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/cmd/follow/dal/mq"
)

func Init() {
	db.Init()
	cache.Init()
	mq.Init()
}
