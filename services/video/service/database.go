package service

import (
	"github.com/ozline/tiktok/services/video/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type DataBaseService struct {
}

func CreateDataBaseTable(dataBaseName string) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&VideoStorageInfo{})
}

func (d *DataBaseService) DataBasePutVideo(video Video, videoID int64) {
	db, err := gorm.Open(sqlite.Open("videoStorage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var videostorageInfo = VideoStorageInfo{
		VideoID:         videoID,
		VideoPlayUrl:    video.PlayUrl,
		VideoCoverUrl:   video.CoverUrl,
		VideoTitle:      video.Title,
		VideoCreateTime: time.Now().Format("2006-01-02 15:04:05"), //当前时间的字符串
		UserName:        video.Author.Name,
	}

	db.Create(&videostorageInfo)
}

func (d *DataBaseService) DataBaseDeleteVideo(videoTitle string, userName string) (VideoStorageInfo, bool) {
	db, err := gorm.Open(sqlite.Open("videoStorage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	videoInfo, findResult := d.DataBaseFindVideoIDByTitle(videoTitle)
	if findResult == false {
		return videoInfo, false
	}

	var deleteState bool

	if videoInfo.UserName == userName {
		db.Delete(&videoInfo, videoInfo.VideoID)
		deleteState = true
	} else {
		deleteState = false
	}

	return videoInfo, deleteState
}

func (d *DataBaseService) DataBaseFindVideoIDByTitle(videoTitle string) (VideoStorageInfo, bool) {
	db, err := gorm.Open(sqlite.Open("videoStorage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var videoStorage = VideoStorageInfo{
		VideoID: -1,
	}
	db.First(&videoStorage, "video_title=?", videoTitle)
	if videoStorage.VideoID == -1 {
		return videoStorage, false
	}
	return videoStorage, true
}

func RandGetNVideo(number int) []VideoStorageInfo {
	rand.Seed(time.Now().Unix())
	videos := make([]VideoStorageInfo, number)

	db, err := gorm.Open(sqlite.Open("videoStorage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var rowsNumber int64
	db.Model(&VideoStorageInfo{}).Count(&rowsNumber)

	var videoStorage VideoStorageInfo
	for i := 0; i < number; i++ {
		index := rand.Intn(int(rowsNumber))
		db.Offset(index).Take(&videoStorage)
		videos[i] = videoStorage
	}
	return videos
}
