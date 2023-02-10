package main

import (
	"context"
	"github.com/golang/glog"
	"github.com/ozline/tiktok/pkg/utils/snowflake"
	video "github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video"
	"github.com/ozline/tiktok/services/video/service"
	"strconv"
	"time"
)

// TiktokVideoServiceImpl implements the last service interface defined in the IDL.
type TiktokVideoServiceImpl struct{}

// PutVideo implements the TiktokVideoServiceImpl interface.
func (s *TiktokVideoServiceImpl) PutVideo(ctx context.Context, req *video.PutVideoRequest) (resp *video.PutVideoResponse, err error) {
	user := service.User{
		Name: req.GetOwnerName(),
	}

	videoInfo := service.Video{
		PlayUrl:  req.GetPlayUrl(),
		CoverUrl: req.GetCoverUrl(),
		Title:    req.GetTitle(),
		Author:   &user,
	}

	// 雪花算法
	snow, err := snowflake.NewSnowflake(int64(1), int64(1))
	if err != nil {
		glog.Error(err)
		return
	}

	videoID := snow.NextVal()
	videoInfo.M.Lock()
	videoInfo.ID = videoID
	videoInfo.M.Unlock()

	s.DataBasePutVideo(videoInfo, videoInfo.ID)
	s.StoragPutVideo(req.PlayUrl, videoInfo.ID, "titok")

	response := video.PutVideoResponse{
		State:     true,
		Title:     videoInfo.Title,
		OwnerName: videoInfo.Author.Name,
	}

	return &response, nil
}

// DeleteVideo implements the TiktokVideoServiceImpl interface.
func (s *TiktokVideoServiceImpl) DeleteVideo(ctx context.Context, req *video.DeleteVideoRequest) (resp *video.DeleteVideoResponse, err error) {
	//fmt.Println("----- DeleteVideo -----")
	videoTitle := req.GetTitle()
	deletorName := req.DeletorName
	videoInfo, dataBaseResult := s.DataBaseDeleteVideo(videoTitle, deletorName)

	storageResult := false
	if dataBaseResult == true {
		storageResult = s.StorageDeleteVideo(videoInfo.VideoID, "titok")
	}

	response := video.DeleteVideoResponse{
		State: false,
	}
	if dataBaseResult == true && storageResult == true {
		response = video.DeleteVideoResponse{
			State:           true,
			DeleteVideoName: videoInfo.VideoTitle,
			DeletorName:     videoInfo.UserName,
			VideoOwnerName:  videoInfo.UserName,
		}
	} else {
		if dataBaseResult == false {
			response.ErrState = "DataBase Delete Failed"
		} else {
			response.ErrState = "Storage Delete Failed"
		}
	}
	return &response, nil
}

// GetOneVideoInfo implements the TiktokVideoServiceImpl interface.
func (s *TiktokVideoServiceImpl) GetOneVideoInfo(ctx context.Context, req *video.GetOneVideoInfoRequest) (resp *video.GetOneVideoInfoResponse, err error) {
	//fmt.Println("----- GetVideoInfo -----")
	videoTitle := req.GetVideoName()
	//userId := req.GetUserId()
	videoInfo, findState := s.DataBaseFindVideoIDByTitle(videoTitle)
	response := video.GetOneVideoInfoResponse{
		State: false,
	}
	if findState == true {
		_, videoSize, videoMimeType := s.StorageGetVideoInfo(videoInfo.VideoID, "titok")
		response = video.GetOneVideoInfoResponse{
			State:         true,
			VideoId:       videoInfo.VideoID,
			PlayUrl:       videoInfo.VideoPlayUrl,
			CoverUrl:      videoInfo.VideoCoverUrl,
			VideoTitle:    videoInfo.VideoTitle,
			VideoSize:     videoSize,
			VideoMimeType: videoMimeType,
			OwnerName:     videoInfo.UserName,
		}
	} else {
		response.ErrState = "----- Don't Find the Video -----"
	}
	return &response, nil
}

// DownloadOneVideo implements the TiktokVideoServiceImpl interface.
func (s *TiktokVideoServiceImpl) DownloadOneVideo(ctx context.Context, req *video.DownloadOneVideoRequest) (resp *video.DownloadOneVideoResponse, err error) {
	//fmt.Println("----- DownloadOneVideo -----")
	videoTitle := req.GetVideoName()
	videoInfo, dataBaseResult := s.DataBaseFindVideoIDByTitle(videoTitle)
	response := video.DownloadOneVideoResponse{
		State: false,
	}
	if dataBaseResult == true {
		accessURL := service.StorageDownloadOneVideo(videoInfo.VideoID, "titok")
		response = video.DownloadOneVideoResponse{
			State:      true,
			VideoTitle: videoInfo.VideoTitle,
			VideoUrl:   accessURL,
			OwnerName:  videoInfo.UserName,
		}
	}

	return &response, nil
}

// DownloadMultiVideo implements the TiktokVideoServiceImpl interface.
func (s *TiktokVideoServiceImpl) DownloadMultiVideo(ctx context.Context, req *video.DownloadMultiVideoRequest) (resp *video.DownloadMultiVideoResponse, err error) {
	number := int(req.GetVideoNumber())
	keys, err := RDB.Keys("*").Result()

	videos := service.RandGetNVideo(number)
	videourls := make([]string, number)
	//fmt.Println("len(keys)=", len(keys))
	//fmt.Println("Number=", number)
	if len(keys) >= int(number) {
		//fmt.Println("----- Redis Cache Has Enough Videos -----")
		for i := 0; i < number; i++ {
			videourls[i], err = RDB.Get(keys[i]).Result()
		}
		//for index, videoid := range keys {
		//	videourls[index], err = RDB.Get(videoid).Result()
		//}
	} else {
		//fmt.Println("----- Redic Cache Don't Have Enough Videos -----")
		videourls = service.GetNUrlByVideoID(videos)
	}

	response := video.DownloadMultiVideoResponse{
		VideoNumber: int64(len(videourls)),
		VideoUrls:   videourls,
	}

	if response.VideoNumber == req.GetVideoNumber() {
		response.State = true
	} else {
		response.State = false
		response.ErrState = "Dont't Have Enough Videos"
	}

	return &response, nil
}

func PerioUpdateVideoCache(number int) {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		RDB.FlushDB().Result()
		videos := service.RandGetNVideo(number)
		videourls := service.GetNUrlByVideoID(videos)
		for index, videoInfo := range videos {

			RDB.Set(strconv.FormatInt(videoInfo.VideoID, 10), videourls[index], 0).Result()
		}
	}
}
