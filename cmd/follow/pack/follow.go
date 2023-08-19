package pack

import (
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/user"
)

func User(user *user.User) *follow.User {
	return &follow.User{
		Id:              user.Id,
		Name:            user.Name,
		FollowCount:     &user.FollowCount,
		FollowerCount:   &user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          &user.Avatar,
		BackgroundImage: &user.BackgroundImage,
		Signature:       &user.Signature,
		TotalFavorited:  &user.TotalFavorited,
		WorkCount:       &user.WorkCount,
		FavoriteCount:   &user.FavoritedCount,
	}
}

// func FriendUser(user *follow.User, msg string, msgType int64) *follow.FriendUser {
// 	return &follow.FriendUser{
// 		Id:              user.Id,
// 		Name:            user.Name,
// 		FollowCount:     user.FollowCount,
// 		FollowerCount:   user.FollowerCount,
// 		IsFollow:        user.IsFollow,
// 		Avatar:          user.Avatar,
// 		BackgroundImage: user.BackgroundImage,
// 		Signature:       user.Signature,
// 		TotalFavorited:  user.TotalFavorited,
// 		WorkCount:       user.WorkCount,
// 		FavoriteCount:   user.FavoriteCount,
// 		Message:         &msg,
// 		MsgType:         msgType,
// 	}
// }
