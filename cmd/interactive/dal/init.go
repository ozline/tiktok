package dal

import (
	"github.com/ozline/tiktok/cmd/interactive/dal/db"
	"github.com/ozline/tiktok/cmd/interactive/dal/sensitive_words"
)

func Init() {
	db.Init()
	sensitive_words.Init()
}
