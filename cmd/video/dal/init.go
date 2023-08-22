package dal

import (
	"github.com/ozline/tiktok/cmd/video/dal/cache"
	"github.com/ozline/tiktok/cmd/video/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
