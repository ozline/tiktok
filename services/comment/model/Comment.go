package model

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID             int64
	UserID         int64
	VideoID        int64
	Content        string
	IsUploaderLike bool
	LikeCount      int32

	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
