package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	Id            int64  // 用户id
	Name          string // 用户名称
	FollowCount   int64  // 关注总数
	FollowerCount int64  // 粉丝总数
	IsFollow      bool   // true-已关注，false-未关注
}

type VideoStorageInfo struct {
	VideoID         string // 视频id
	VideoTitle      string // 视频标题
	VideoCreateTime string // 视频创建时间
	UserId          int64  // 用户id
	UserName        string // 用户名称
}

func (s *TiktokVideoServiceImpl) CreateDataBaseTable(dataBaseName string) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&VideoStorageInfo{})
}

func (s *TiktokVideoServiceImpl) DataBasePutFile(video Video, privateKey string) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var videostorageInfo = VideoStorageInfo{
		VideoTitle: video.Title,
		UserId:     video.Author.Id,
		UserName:   video.Author.Name,
	}
	videostorageInfo.VideoID = privateKey

	db.Create(&videostorageInfo)
}

func (s *TiktokVideoServiceImpl) DataBaseDeleteFile(fileName string) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var video VideoStorageInfo
	db.Where("VideoTitle=?", fileName).First(&video)
	db.Delete(&video)
}
