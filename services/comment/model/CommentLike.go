package model

type CommentLike struct {
	UserID    int64 `gorm:"primary_key;autoIncrement:false"`
	CommentId int64 `gorm:"primary_key;autoIncrement:false"`
}
