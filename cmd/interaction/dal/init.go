package dal

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/cmd/interaction/dal/sensitive_words"
)

func Init() {
	db.Init()
	sensitive_words.Init()
}
