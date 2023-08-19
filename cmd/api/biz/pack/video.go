package pack

import (
	"github.com/ozline/tiktok/cmd/api/biz/model/api"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

func VideoList(list []*video.Video) []*api.Video {
	resp := make([]*api.Video, 0)

	for _, data := range list {
		resp = append(resp, &api.Video{
			ID: data.Id,
			Author: &api.User{
				ID:              data.Anthor.Id,
				Name:            data.Anthor.Name,
				FollowCount:     &data.Anthor.FollowCount,
				FollowerCount:   &data.Anthor.FollowerCount,
				IsFollow:        data.Anthor.IsFollow,
				Avatar:          &data.Anthor.Avatar,
				BackgroundImage: &data.Anthor.BackgroundImage,
				Signature:       &data.Anthor.Signature,
				// TotalFavorited:  data.Anthor.TotalFavorited,
				// TODO: fix type error
				WorkCount:     &data.Anthor.WorkCount,
				FavoriteCount: &data.Anthor.FavoriteCount,
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
