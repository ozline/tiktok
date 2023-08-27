package pack

import (
	"github.com/ozline/tiktok/cmd/api/biz/model/api"
	"github.com/ozline/tiktok/kitex_gen/video"
)

func VideoList(list []*video.Video) []*api.Video {
	resp := make([]*api.Video, 0)

	for _, data := range list {
		resp = append(resp, &api.Video{
			ID: data.Id,
			Author: &api.User{
				ID:              data.Author.Id,
				Name:            data.Author.Name,
				FollowCount:     &data.Author.FollowCount,
				FollowerCount:   &data.Author.FollowerCount,
				IsFollow:        data.Author.IsFollow,
				Avatar:          &data.Author.Avatar,
				BackgroundImage: &data.Author.BackgroundImage,
				Signature:       &data.Author.Signature,
				TotalFavorited:  &data.Author.TotalFavorited,
				WorkCount:       &data.Author.WorkCount,
				FavoriteCount:   &data.Author.FavoriteCount,
			},
			PlayURL:       data.PlayUrl,
			CoverURL:      data.CoverUrl,
			FavoriteCount: data.FavoriteCount,
			CommentCount:  data.CommentCount,
			IsFavorite:    data.IsFavorite,
			Title:         data.Title,
		})
	}

	return resp
}
