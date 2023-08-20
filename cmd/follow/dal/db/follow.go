package db

import (
	"context"
	"errors"
	"time"

	"github.com/ozline/tiktok/pkg/constants"
	"gorm.io/gorm"
)

type Follow struct {
	Id        int64
	UserId    int64
	ToUserId  int64
	Status    int64 //1-关注, 2-取消关注
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var RecordNotFound = gorm.ErrRecordNotFound

func FollowAction(ctx context.Context, follow *Follow) error {
	followResp := new(Follow)

	//查询db中是否存在记录(判断是否是第一次操作)
	err := DB.WithContext(ctx).Model(&Follow{}).
		Where("user_id= ? AND to_user_id = ?", follow.UserId, follow.ToUserId).
		First(&followResp).Error

	//db中查询不到,创建关注
	if errors.Is(err, gorm.ErrRecordNotFound) {
		follow.Id = SF.NextVal()
		return DB.WithContext(ctx).Create(follow).Error
	} else if err != nil {
		return err
	}

	//db中存在,修改Status
	err = DB.WithContext(ctx).Model(&Follow{}).
		Where("user_id= ? AND to_user_id = ?", follow.UserId, follow.ToUserId).
		Update("status", constants.FollowAction).Error
	if err != nil {
		return err
	}

	return nil
}

func UnFollowAction(ctx context.Context, follow *Follow) error {
	//修改db中的status
	err := DB.WithContext(ctx).Model(&Follow{}).
		Where("user_id= ? AND to_user_id = ?", follow.UserId, follow.ToUserId).
		Update("status", constants.UnFollowAction).Error
	if err != nil {
		return err
	}

	return nil
}

// 关注列表(获取to_user_id的列表)
func FollowListAction(ctx context.Context, uid int64) (*[]int64, error) {
	var followList []int64

	//redis中不存在,查询db并返回结果
	err := DB.WithContext(ctx).Table(constants.FollowTableName).Select("to_user_id").
		Where("user_id = ? AND status = ?", uid, constants.FollowAction).
		Find(&followList).Error
	if err != nil {
		return nil, err
	}

	return &followList, nil
}

// 粉丝列表(获取user_id的列表)
func FollowerListAction(ctx context.Context, uid int64) (*[]int64, error) {
	var followerList []int64

	//redis中不存在,查询db并返回结果
	err := DB.WithContext(ctx).Table(constants.FollowTableName).Select("user_id").
		Where("to_user_id = ? AND status = ?", uid, constants.FollowAction).
		Find(&followerList).Error
	if err != nil {
		return nil, err
	}

	return &followerList, nil
}

// 好友列表(先获取to_user_id的列表)
func FriendListAction(ctx context.Context, uid int64) (*[]int64, error) {
	var friendList []int64

	//先获取本人的关注列表
	tempList, err := FollowListAction(ctx, uid)
	if err != nil {
		return nil, err
	}

	//查询db中的粉丝列表
	for _, id := range *tempList {
		err = DB.WithContext(ctx).Model(&Follow{}).
			Where("user_id = ? AND to_user_id = ? AND status = ?", id, uid, constants.FollowAction).Error
		if errors.Is(err, gorm.ErrRecordNotFound) { //db中也不存在,说明只是单方面关注，不是好友
			continue
		} else if err != nil {
			return nil, err
		}

		//粉丝列表存在就直接添加这个id
		friendList = append(friendList, id)
	}

	return &friendList, nil
}
