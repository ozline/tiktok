package pack

import (
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

// 用于获取喜欢列表 IsFavorite保证为true
func VideoLiked(data *db.Video) *video.Video {
	if data == nil {
		return nil
	}
	return &video.Video{
		Id:            data.Id,
		Anthor:        nil, //TODO
		PlayUrl:       data.PlayUrl,
		CoverUrl:      data.CoverUrl,
		FavoriteCount: data.FavoriteCount,
		CommentCount:  data.FavoriteCount,
		IsFavorite:    true,
		Title:         data.Title,
	}
}
func VideoLikedList(data []db.Video) []*video.Video {
	videoList := make([]*video.Video, 0, len(data))
	for _, v := range data {
		videoList = append(videoList, VideoLiked(&v))
	}
	return videoList
}
