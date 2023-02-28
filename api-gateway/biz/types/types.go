package types

import "github.com/ozline/tiktok/services/user/model"

type Comment struct {
	Id         int64
	User       *model.User
	Content    string
	CreateDate string
}
