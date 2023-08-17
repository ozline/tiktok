package dal

import (
	"github.com/ozline/tiktok/cmd/user/dal/cache"
	"github.com/ozline/tiktok/cmd/user/dal/db"
	"github.com/ozline/tiktok/cmd/user/dal/mq"
)

func Init() {
	db.Init()
	cache.Init()
	mq.Init()
}
