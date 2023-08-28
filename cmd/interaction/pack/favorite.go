package pack

import (
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/kitex_gen/video"
)

// build *user.User
func BuildUser(data *video.User) *user.User {
	if data == nil {
		return nil
	}

	return &user.User{
		Id:              data.Id,
		Name:            data.Name,
		Avatar:          data.Avatar,
		FollowCount:     data.FollowCount,
		FollowerCount:   data.FollowerCount,
		IsFollow:        data.IsFollow,
		BackgroundImage: data.BackgroundImage,
		Signature:       data.Signature,
		WorkCount:       data.WorkCount,
		FavoritedCount:  data.FavoriteCount,
		TotalFavorited:  data.TotalFavorited,
	}
}

// convert video.Video To interaction.Video
func BuildVideo(data *video.Video) *interaction.Video {
	if data == nil {
		return nil
	}

	return &interaction.Video{
		Id:            data.Id,
		Author:        BuildUser(data.Author),
		PlayUrl:       data.PlayUrl,
		CoverUrl:      data.CoverUrl,
		FavoriteCount: data.FavoriteCount,
		CommentCount:  data.CommentCount,
		IsFavorite:    data.IsFavorite,
		Title:         data.Title,
	}
}

// build []*interaction.Video
func BuildVideos(datas []*video.Video) []*interaction.Video {
	if datas == nil {
		return nil
	}

	videos := make([]*interaction.Video, 0, len(datas))

	for _, data := range datas {
		videos = append(videos, BuildVideo(data))
	}
	return videos
}
