package dal

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/cmd/interaction/dal/mq"
	"github.com/ozline/tiktok/cmd/interaction/dal/sensitive_words"
)

func Init(path string) {
	db.Init()
	cache.Init()
	mq.Init()
	sensitive_words.Init(path)
}
