package db

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/ozline/tiktok/pkg/constants"
	"gorm.io/gorm"
)

type Follow struct {
	Id         int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	UserId     int64
	ToUserId   int64
	ActionType int64 `gorm:"default:1"` //1-关注, 2-取消关注
}

const (
	Action       = "followAction"
	FollowList   = "followList"
	FollowerList = "followerList"
	FriendList   = "friendList"
)

func FollowAction(ctx context.Context, follow *Follow) error {
	followResp := new(Follow)

	field := strconv.FormatInt(follow.UserId, 10)
	value := strconv.FormatInt(follow.ToUserId, 10)

	//判断数据是否存在与redis(存在为已关注，不存在为已取关)
	b, err := RedisClient.HExists(ctx, Action, field).Result()
	if err != nil {
		return err
	}

	//数据存在,进行取关
	if b {
		err = DB.WithContext(ctx).Model(&Follow{}).
			Where("user_id= ? AND to_user_id = ?", follow.UserId, follow.ToUserId).
			Update("action_type", follow.ActionType).Error
		if err == nil {
			return RedisClient.HDel(ctx, Action, field).Err()
		}
		return err
	}

	//--------------//
	//TODO:redis限流//
	//-------------//

	//数据不存在，进行关注
	err = RedisClient.HSet(ctx, Action, field, value).Err()
	if err != nil {
		return err
	}

	//查询db
	err = DB.WithContext(ctx).Model(&Follow{}).
		Where("user_id= ? AND to_user_id = ?", follow.UserId, follow.ToUserId).
		First(&followResp).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) { //db中也查询不到,创建关注
		follow.Id = SF.NextVal()
		return DB.WithContext(ctx).Create(follow).Error
	} else if err == nil {
		return DB.WithContext(ctx).Model(&Follow{}).
			Where("user_id= ? AND to_user_id = ?", follow.UserId, follow.ToUserId).
			Update("action_type", follow.ActionType).Error
	}

	return err
}

// 关注列表(获取to_user_id的列表)
func FollowListAction(ctx context.Context, uid int64) (*[]int64, error) {
	var followList []int64
	//TODO:redis缓存

	err := DB.WithContext(ctx).Model(&Follow{}).Select("to_user_id").
		Where("user_id = ? AND action_type = ?", uid, constants.FollowAction).
		Find(&followList).Error

	if err != nil {
		return nil, err
	}
	return &followList, nil
}

// 粉丝列表(获取user_id的列表)
func FollowerListAction(ctx context.Context, uid int64) (*[]int64, error) {
	var followerList []int64
	//TODO:redis缓存
	err := DB.WithContext(ctx).Model(&Follow{}).Select("user_id").
		Where("to_user_id = ? AND action_type = ?", uid, constants.FollowAction).
		Find(&followerList).Error

	if err != nil {
		return nil, err
	}
	return &followerList, nil
}

// 好友列表(获取to_user_id的列表)
func FriendListAction(ctx context.Context, uid int64) (*[]int64, error) {
	var tempList []int64
	var friendList []int64

	//TODO:redis缓存

	//先获取本人关注的列表
	err := DB.WithContext(ctx).Model(&Follow{}).Select("to_user_id").
		Where("user_id = ? AND action_type = ?", uid, constants.FollowAction).
		Find(&tempList).Error

	if err != nil {
		return nil, err
	}

	for _, id := range tempList {
		err = DB.WithContext(ctx).Model(&Follow{}).
			Where("user_id = ? AND to_user_id = ? AND action_type = ?", id, uid, constants.FollowAction).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				continue
			}
			return nil, err
		}
		friendList = append(friendList, id)
	}
	return &friendList, nil
}
