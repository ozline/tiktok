package pack

import (
	"github.com/ozline/tiktok/kitex_gen/user"

	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

func Video(data *db.Video, user *user.User) *video.Video {
	if data == nil {
		return nil
	}
	return &video.Video{
		Id: data.Id,
		Anthor: &video.User{
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
		FavoriteCount: 0,    //TODO
		CommentCount:  0,    //TODO
		IsFavorite:    true, //TODO
		Title:         data.Title,
	}
}
func VideoList(data []db.Video, userList []*user.User) []*video.Video {
	videoList := make([]*video.Video, 0, len(data))
	for i := 0; i < len(data); i++ {
		videoList = append(videoList, Video(&data[i], userList[i]))
	}
	return videoList
}

// 用于获取喜欢列表 IsFavorite保证为true
func VideoLiked(data *db.Video, user *user.User) *video.Video {
	if data == nil {
		return nil
	}
	return &video.Video{
		Id: data.Id,
		Anthor: &video.User{
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
		FavoriteCount: 0, //TODO
		CommentCount:  0, //TODO
		IsFavorite:    true,
		Title:         data.Title,
	}
}
func VideoLikedList(data []db.Video, userList []*user.User) []*video.Video {
	videoList := make([]*video.Video, 0, len(data))
	for i := 0; i < len(data); i++ {
		videoList = append(videoList, VideoLiked(&data[i], userList[i]))
	}
	return videoList
}
