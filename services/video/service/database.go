package service

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"sync"
	"time"
)

type TiktokVideoServiceImpl struct{}
type User struct {
	Id            int64  // 用户id
	Name          string // 用户名称
	FollowCount   int64  // 关注总数
	FollowerCount int64  // 粉丝总数
	IsFollow      bool   // true-已关注，false-未关注
}

type Video struct {
	M             sync.Mutex //
	ID            int64      // 视频ID
	Author        *User      // 作者信息
	PlayUrl       string     // 播放地址
	CoverUrl      string     // 封面地址
	FavoriteCount int64      // 视频的点赞总数
	CommentCount  int64      // 视频的评论总数
	IsFavorite    bool       // 是否本人已点赞
	Title         string     // 标题
}

type VideoStorageInfo struct {
	VideoID         int64  // 视频id
	VideoPlayUrl    string // 视频标题
	VideoCoverUrl   string
	VideoTitle      string
	VideoCreateTime string // 视频创建时间
	UserName        string // 用户名称
}

func CreateDataBaseTable(dataBaseName string) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&VideoStorageInfo{})
}

func (s *TiktokVideoServiceImpl) DataBasePutVideo(video Video, videoID int64) {
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

func (s *TiktokVideoServiceImpl) DataBaseDeleteVideo(videoTitle string, userName string) (VideoStorageInfo, bool) {
	db, err := gorm.Open(sqlite.Open("videoStorage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	videoInfo, findResult := s.DataBaseFindVideoIDByTitle(videoTitle)
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

func (s *TiktokVideoServiceImpl) DataBaseFindVideoIDByTitle(videoTitle string) (VideoStorageInfo, bool) {
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
