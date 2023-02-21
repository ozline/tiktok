package service

import (
	"github.com/ozline/tiktok/kitex_gen/tiktok/follow"
	"github.com/ozline/tiktok/services/follow/model"
)

// 由于本项目较小，所以并没有区分 service 下的函数
// 但是如果本项目要进行拓展，请务必将 service 层下的服务进行分类拆分

func AddFollowRelation(userId, toUserId int64) int64 {
	err := model.AddRelation(userId, toUserId)
	if err != nil {
		return Success
	} else if err.Error() == "the relation have been existed" {
		return InvalidParams
	} else if err.Error() == "database error" {
		return ErrorDBRecord
	} else {
		return Error
	}
}

func RemoveFollowRelation(userId, toUserId int64) int64 {
	err := model.DeleteRelation(userId, toUserId)
	if err != nil {
		return Success
	} else if err.Error() == "the relation have been existed" {
		return InvalidParams
	} else if err.Error() == "database error" {
		return ErrorDBRecord
	} else {
		return Error
	}
}

func QueryRelation(userId, toUserId int64) int32 {
	// 如果是关注关系，则可能互关可能关注
	// 如果不是关注关系，则可能粉丝可能未关注
	if model.IsFollowRelation(userId, toUserId) {
		if model.IsFollowerRelation(userId, toUserId) {
			return 4
		}
		return 2
	} else {
		if model.IsFollowerRelation(userId, toUserId) {
			return 3
		}
		return 1
	}
}

// ListFollows 获取关注清单
func ListFollows(userId int64, pageNum int32, pageSize int32) ([]*follow.User, int64) {
	data, err := model.QueryFollowsList(userId, pageNum, pageSize)
	if err != nil {
		return nil, Error
	}
	return buildUserList(data), Success
}

// ListFollowers 获取粉丝清单
func ListFollowers(userId int64, pageNum int32, pageSize int32) ([]*follow.User, int64) {
	data, err := model.QueryFollowersList(userId, pageNum, pageSize)
	if err != nil {
		return nil, Error
	}
	return buildUserList(data), Success
}

// ListFriends 获取好友清单
func ListFriends(userId int64, pageNum int32, pageSize int32) ([]*follow.User, int64) {
	data, err := model.QueryFriendsList(userId, pageNum, pageSize)
	if err != nil {
		return nil, Error
	}
	return buildUserList(data), Success
}

func buildUser(item model.User) follow.User {
	return follow.User{
		Id:            item.ID,
		Name:          item.Username,
		FollowCount:   &item.FollowCount,
		FollowerCount: &item.FollowerCount,
		IsFollow:      true,
	}
}

func buildUserList(items []model.User) (data []*follow.User) {
	for _, item := range items {
		user := buildUser(item)
		data = append(data, &user)
	}
	return
}
