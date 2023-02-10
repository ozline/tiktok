package main

import (
	"context"
	video "github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video"
	"github.com/ozline/tiktok/services/video/model"
	"github.com/ozline/tiktok/services/video/service"
	"strconv"
	"time"
)

// TiktokVideoServiceImpl implements the last service interface defined in the IDL.
type TiktokVideoServiceImpl struct{}

// PutVideo implements the TiktokVideoServiceImpl interface.
func (s *TiktokVideoServiceImpl) PutVideo(ctx context.Context, req *video.PutVideoRequest) (resp *video.PutVideoResponse, err error) {
	user := model.User{
		Name: req.GetOwnerName(),
	}

	videoInfo := model.Video{
		PlayUrl:  req.GetPlayUrl(),
		CoverUrl: req.GetCoverUrl(),
		Title:    req.GetTitle(),
		Author:   &user,
	}
	dataSerivce := service.NewDataBaseService(ctx)
	videoID := dataSerivce.S.NextVal()
	videoInfo.M.Lock()
	videoInfo.ID = videoID
	videoInfo.M.Unlock()

	go dataSerivce.DataBasePutVideo(videoInfo, videoInfo.ID)
	go service.NewStorageService(ctx).StoragPutVideo(req.PlayUrl, videoInfo.ID, "titok")

	response := video.PutVideoResponse{
		State:     true,
		Title:     videoInfo.Title,
		OwnerName: videoInfo.Author.Name,
	}
	return &response, nil
}

// DeleteVideo implements the TiktokVideoServiceImpl interface.
func (s *TiktokVideoServiceImpl) DeleteVideo(ctx context.Context, req *video.DeleteVideoRequest) (resp *video.DeleteVideoResponse, err error) {
	videoTitle := req.GetTitle()
	deletorName := req.DeletorName
	videoInfo, dataBaseResult := service.NewDataBaseService(ctx).DataBaseDeleteVideo(videoTitle, deletorName)

	storageResult := false
	if dataBaseResult == true {
		storageResult = service.NewStorageService(ctx).StorageDeleteVideo(videoInfo.VideoID, "titok")
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
	videoInfo, findState := service.NewDataBaseService(ctx).DataBaseFindVideoIDByTitle(videoTitle)
	response := video.GetOneVideoInfoResponse{
		State: false,
	}
	if findState == true {
		_, videoSize, videoMimeType := service.NewStorageService(ctx).StorageGetVideoInfo(videoInfo.VideoID, "titok")
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
	videoInfo, dataBaseResult := service.NewDataBaseService(ctx).DataBaseFindVideoIDByTitle(videoTitle)
	response := video.DownloadOneVideoResponse{
		State: false,
	}
	if dataBaseResult == true {
		accessURL := service.NewStorageService(ctx).StorageDownloadOneVideo(videoInfo.VideoID, "titok")
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
		videourls = service.NewStorageService(ctx).GetNUrlByVideoID(videos)
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

func (s *TiktokVideoServiceImpl) PerioUpdateVideoCache(ctx context.Context, number int) {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		RDB.FlushDB().Result()
		videos := service.RandGetNVideo(number)
		videourls := service.NewStorageService(ctx).GetNUrlByVideoID(videos)
		for index, videoInfo := range videos {

			RDB.Set(strconv.FormatInt(videoInfo.VideoID, 10), videourls[index], 0).Result()
		}
	}
}
