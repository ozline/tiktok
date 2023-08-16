package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	Id         int64
	ToUserId   int64
	FromUserId int64
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func GetMessageList(ctx context.Context, to_user_id int64, from_user_id int64) ([]*Message, error) {
	return nil, nil
}
