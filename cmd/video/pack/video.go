package pack

import (
	"fmt"
	"time"

	"github.com/ozline/tiktok/kitex_gen/user"

	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/kitex_gen/video"
)

func Video(data *db.Video, user *user.User, favoriteCount int64, commentCount int64, isFavorite bool) *video.Video {
	if data == nil {
		return nil
	}
	return &video.Video{
		Id: data.Id,
		Author: &video.User{
			Id:              user.Id,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        user.IsFollow,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoritedCount,
		},
		PlayUrl:       data.PlayUrl,
		CoverUrl:      data.CoverUrl,
		FavoriteCount: favoriteCount,
		CommentCount:  commentCount,
		IsFavorite:    isFavorite,
		Title:         data.Title,
	}
}
func VideoList(data []db.Video, userList []*user.User, favoriteCountList []int64, commentCountList []int64, isFavoriteList []bool) []*video.Video {
	videoList := make([]*video.Video, 0, len(data))
	for i := 0; i < len(data); i++ {
		videoList = append(videoList, Video(&data[i], userList[i], favoriteCountList[i], commentCountList[i], isFavoriteList[i]))
	}
	return videoList
}

// 用于获取喜欢列表 IsFavorite保证为true
func VideoLiked(data *db.Video, user *user.User, favoriteCount int64, commentCount int64) *video.Video {
	if data == nil {
		return nil
	}
	return &video.Video{
		Id: data.Id,
		Author: &video.User{
			Id:              user.Id,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        user.IsFollow,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoritedCount,
		},
		PlayUrl:       data.PlayUrl,
		CoverUrl:      data.CoverUrl,
		FavoriteCount: favoriteCount,
		CommentCount:  commentCount,
		IsFavorite:    true,
		Title:         data.Title,
	}
}
func VideoLikedList(data []db.Video, userList []*user.User, favoriteCountList []int64, commentCountList []int64) []*video.Video {
	videoList := make([]*video.Video, 0, len(data))

	for i := 0; i < len(data); i++ {
		videoList = append(videoList, VideoLiked(&data[i], userList[i], favoriteCountList[i], commentCountList[i]))
	}

	return videoList
}
func GenerateVideoName(userID int64) string {
	currentTime := time.Now()
	// 获取年月日和小时分钟
	year, month, day := currentTime.Date()
	hour, minute := currentTime.Hour(), currentTime.Minute()
	return fmt.Sprintf("%v_%d%02d%02d_%02d%02d_video.mp4", userID, year, month, day, hour, minute)
}
func GenerateCoverName(userID int64) string {
	currentTime := time.Now()
	// 获取年月日和小时分钟
	year, month, day := currentTime.Date()
	hour, minute := currentTime.Hour(), currentTime.Minute()
	return fmt.Sprintf("%v_%d%02d%02d_%02d%02d_cover.jpg", userID, year, month, day, hour, minute)
}
