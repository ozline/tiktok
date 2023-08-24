package pack

import (
	"github.com/ozline/tiktok/cmd/api/biz/model/api"
	"github.com/ozline/tiktok/kitex_gen/follow"
)

// serve for follow list
func UserList(list []*follow.User) []*api.User {
	resp := make([]*api.User, 0)

	for _, data := range list {
		resp = append(resp, &api.User{
			ID:              data.Id,
			Name:            data.Name,
			FollowCount:     data.FollowCount,
			FollowerCount:   data.FollowerCount,
			IsFollow:        data.IsFollow,
			Avatar:          data.Avatar,
			BackgroundImage: data.BackgroundImage,
			Signature:       data.Signature,
			TotalFavorited:  data.TotalFavorited,
			WorkCount:       data.WorkCount,
			FavoriteCount:   data.FavoriteCount,
		})
	}

	return resp
}

func FriendList(list []*follow.FriendUser) []*api.FriendUser {
	resp := make([]*api.FriendUser, 0)

	for _, data := range list {
		resp = append(resp, &api.FriendUser{
			ID:              data.User.Id,
			Name:            data.User.Name,
			FollowCount:     data.User.FollowCount,
			FollowerCount:   data.User.FollowerCount,
			IsFollow:        data.User.IsFollow,
			Avatar:          data.User.Avatar,
			BackgroundImage: data.User.BackgroundImage,
			Signature:       data.User.Signature,
			TotalFavorited:  data.User.TotalFavorited,
			WorkCount:       data.User.WorkCount,
			FavoriteCount:   data.User.FavoriteCount,
			Message:         data.Message,
			MsgType:         data.MsgType,
		})
	}

	return resp
}
