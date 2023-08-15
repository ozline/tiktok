package db

import (
	"context"
	"errors"
	"time"

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

func FollowAction(ctx context.Context, follow *Follow) error {
	followResp := new(Follow)

	//TODO:redis缓存

	//若查无此数据，则写入数据库且成功关注
	err := DB.WithContext(ctx).Model(&Follow{}).
		Where("user_id= ? AND to_user_id = ?", follow.UserId, follow.ToUserId).
		First(&followResp).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			follow.Id = SF.NextVal()
			return DB.WithContext(ctx).Create(follow).Error
		}
		return err
	}

	//关注/取关操作
	err = DB.WithContext(ctx).Model(&Follow{}).
		Where("user_id= ? AND to_user_id = ?", follow.UserId, follow.ToUserId).
		Update("action_type", follow.ActionType).Error
	if err != nil {
		return err
	}
	return nil
}

// 关注列表(获取to_user_id的列表)
func FollowListAction(ctx context.Context, uid int64) (*[]int64, error) {
	var followList []int64
	//TODO:redis缓存
	err := DB.WithContext(ctx).Model(&Follow{}).Select("to_user_id").
		Where("user_id = ? AND action_type = ?", uid, 1).
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
		Where("to_user_id = ? AND action_type = ?", uid, 1).
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
		Where("user_id = ? AND action_type = ?", uid, 1).
		Find(&tempList).Error

	if err != nil {
		return nil, err
	}

	for _, id := range tempList {
		err = DB.WithContext(ctx).Model(&Follow{}).
			Where("user_id = ? AND to_user_id = ? AND action_type = ?", id, uid, 1).Error
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
