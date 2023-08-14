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

// func GetUserID(ctx context.Context)(int64,error){
// 	var id int64

//		err:=DB.WithContext(ctx).
//	}
//
// TODO:判断uid以及touid是否存在与User表中
func FollowAction(ctx context.Context, follow *Follow) (*Follow, error) {
	followResp := new(Follow)

	//TODO:redis缓存
	//若为空，则写入数据库且成功关注
	err := DB.WithContext(ctx).Model(&Follow{}).
		Where("user_id= ? AND to_user_id = ?", follow.UserId, follow.ToUserId).
		First(&followResp).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := DB.WithContext(ctx).Create(follow).Error; err != nil {
				return nil, err
			}
			return follow, nil
		}
		return nil, err
	}

	//关注操作
	err = DB.WithContext(ctx).Model(&Follow{}).
		Where("user_id= ? AND to_user_id = ?", follow.UserId, follow.ToUserId).
		Update("action_type", follow.ActionType).Error
	if err != nil {
		return nil, err
	}
	return follow, nil
}
