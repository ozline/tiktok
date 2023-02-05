package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	id         int64
	FromUserId int64
	ToUserID   int64
	content    string
}

func (table *Message) TableName() string {
	return "Message"
}
