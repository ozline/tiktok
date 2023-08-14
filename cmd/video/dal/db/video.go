package db

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	Id            int64
	UserID        int64
	FavoriteCount int64
	CommentCount  int64
	PlayUrl       string
	CoverUrl      string
	Title         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
