package db

import (
	"time"

	"gorm.io/gorm"
)

type Follow struct {
	Id         int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	UserId     int64
	ToUserId   int64
	ActionType int64 `gorm:"default:1"` //1-关注, 2-取消关注
}
