package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (s *TiktokChatServiceImpl) receive_message_mysql_handler(message Message) {
	db, err := gorm.Open(sqlite.Open("receiveMessage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&message)
}

func (s *TiktokChatServiceImpl) send_message_mysql_handler(message Message) {
	db, err := gorm.Open(sqlite.Open("sendMessage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&message)
}
