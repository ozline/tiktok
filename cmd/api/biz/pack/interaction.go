package pack

import (
	"github.com/ozline/tiktok/cmd/api/biz/model/api"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func VideoListFavorited(list []*interaction.Video) []*api.Video {
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
				FavoriteCount:   &data.Author.FavoritedCount,
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

func Comment(data *interaction.Comment) *api.Comment {
	return &api.Comment{
		ID: data.Id,
		User: &api.User{
			ID:              data.User.Id,
			Name:            data.User.Name,
			FollowCount:     &data.User.FollowCount,
			FollowerCount:   &data.User.FollowerCount,
			IsFollow:        data.User.IsFollow,
			Avatar:          &data.User.Avatar,
			BackgroundImage: &data.User.BackgroundImage,
			Signature:       &data.User.Signature,
			TotalFavorited:  &data.User.TotalFavorited,
			WorkCount:       &data.User.WorkCount,
			FavoriteCount:   &data.User.FavoritedCount,
		},
		Content:    data.Content,
		CreateDate: data.CreateDate,
	}
}

func CommentList(list []*interaction.Comment) []*api.Comment {
	resp := make([]*api.Comment, 0)

	for _, data := range list {
		resp = append(resp, Comment(data))
	}

	return resp
}
