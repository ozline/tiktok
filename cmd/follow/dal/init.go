package dal

import (
	"github.com/ozline/tiktok/cmd/follow/dal/cache"
	"github.com/ozline/tiktok/cmd/follow/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
