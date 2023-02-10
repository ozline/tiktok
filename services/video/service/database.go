package service

import (
	"context"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils/snowflake"
	"github.com/ozline/tiktok/services/video/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type DataBaseService struct {
	Ctx context.Context
	S   *snowflake.Snowflake
}

func NewDataBaseService(ctx context.Context) *DataBaseService {
	sf, _ := snowflake.NewSnowflake(constants.SnowflakeDatacenterID, constants.SnowflakeWorkerID)
	return &DataBaseService{
		Ctx: ctx,
		S:   sf,
	}
}

func (d *DataBaseService) CreateDataBaseTable(dataBaseName string) {
	db, err := gorm.Open(sqlite.Open(dataBaseName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&model.VideoStorageInfo{})
}

func (d *DataBaseService) DataBasePutVideo(video model.Video, videoID int64) {
	db, err := gorm.Open(sqlite.Open("videoStorage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var videostorageInfo = model.VideoStorageInfo{
		VideoID:         videoID,
		VideoPlayUrl:    video.PlayUrl,
		VideoCoverUrl:   video.CoverUrl,
		VideoTitle:      video.Title,
		VideoCreateTime: time.Now().Format("2006-01-02 15:04:05"), //当前时间的字符串
		UserName:        video.Author.Name,
	}

	db.Create(&videostorageInfo)
}

func (d *DataBaseService) DataBaseDeleteVideo(videoTitle string, userName string) (model.VideoStorageInfo, bool) {
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

func (d *DataBaseService) DataBaseFindVideoIDByTitle(videoTitle string) (model.VideoStorageInfo, bool) {
	db, err := gorm.Open(sqlite.Open("videoStorage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var videoStorage = model.VideoStorageInfo{
		VideoID: -1,
	}
	db.First(&videoStorage, "video_title=?", videoTitle)
	if videoStorage.VideoID == -1 {
		return videoStorage, false
	}
	return videoStorage, true
}

func RandGetNVideo(number int) []model.VideoStorageInfo {
	rand.Seed(time.Now().Unix())
	videos := make([]model.VideoStorageInfo, number)

	db, err := gorm.Open(sqlite.Open("videoStorage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var rowsNumber int64
	db.Model(&model.VideoStorageInfo{}).Count(&rowsNumber)

	var videoStorage model.VideoStorageInfo
	for i := 0; i < number; i++ {
		index := rand.Intn(int(rowsNumber))
		db.Offset(index).Take(&videoStorage)
		videos[i] = videoStorage
	}
	return videos
}
