package model

type VideoFavorite struct {
	UserID  int64 `gorm:"primary_key;autoIncrement:false"`
	VideoID int64 `gorm:"primary_key;autoIncrement:false"`
}
