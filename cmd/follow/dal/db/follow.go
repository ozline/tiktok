package db

import (
	"context"
	"errors"
	"time"

	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
	"gorm.io/gorm"
)

type Follow struct {
	Id        int64
	UserID    int64
	ToUserID  int64
	Status    int64 // 1-关注, 2-取消关注
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var RecordNotFound = gorm.ErrRecordNotFound

// 关注
func FollowAction(ctx context.Context, follow *Follow) error {
	followResp := new(Follow)

	// 查询db中是否存在记录(判断是否是第一次操作)
	err := DB.WithContext(ctx).Model(&Follow{}).
		Where("user_id= ? AND to_user_id = ?", follow.UserID, follow.ToUserID).
		First(&followResp).Error

	// db中查询不到,创建关注
	if errors.Is(err, gorm.ErrRecordNotFound) {
		follow.Id = SF.NextVal()
		follow.Status = 1
		return DB.WithContext(ctx).Create(follow).Error
	} else if err != nil {
		return err
	}

	// db中存在,修改Status
	err = DB.WithContext(ctx).Model(&Follow{}).
		Where("user_id= ? AND to_user_id = ?", follow.UserID, follow.ToUserID).
		Update("status", constants.FollowAction).Error
	if err != nil {
		return err
	}

	return nil
}

// 取消关注
func UnFollowAction(ctx context.Context, follow *Follow) error {
	_, err := IsFollow(ctx, follow.UserID, follow.ToUserID)
	if errors.Is(err, RecordNotFound) {
		return errno.NotFollowError
	} else if err != nil {
		return err
	}
	// 修改db中的status
	err = DB.WithContext(ctx).Model(&Follow{}).
		Where("user_id= ? AND to_user_id = ?", follow.UserID, follow.ToUserID).
		Update("status", constants.UnFollowAction).Error
	if err != nil {
		return err
	}

	return nil
}

// 关注列表(获取to_user_id的列表)
func FollowListAction(ctx context.Context, uid int64) (*[]int64, error) {
	var followList []int64

	// redis中不存在,查询db并返回结果
	err := DB.WithContext(ctx).Table(constants.FollowTableName).Select("to_user_id").
		Where("user_id = ? AND status = ?", uid, constants.FollowAction).
		Find(&followList).Error
	if err != nil {
		return nil, err
	}

	if len(followList) == 0 { // db中也查不到
		return nil, RecordNotFound
	}

	return &followList, nil
}

// 粉丝列表(获取user_id的列表)
func FollowerListAction(ctx context.Context, uid int64) (*[]int64, error) {
	var followerList []int64

	// redis中不存在,查询db并返回结果
	err := DB.WithContext(ctx).Table(constants.FollowTableName).Select("user_id").
		Where("to_user_id = ? AND status = ?", uid, constants.FollowAction).
		Find(&followerList).Error
	if err != nil {
		return nil, err
	}

	if len(followerList) == 0 { // db中也查不到
		return nil, RecordNotFound
	}

	return &followerList, nil
}

// 好友列表(先获取to_user_id的列表)
func FriendListAction(ctx context.Context, uid int64) (*[]int64, error) {
	friendList := make([]int64, 0, 10)
	// 先获取本人的关注列表
	tempList, err := FollowListAction(ctx, uid)
	if err != nil {
		return nil, err
	}

	// 查询db中的粉丝列表
	for _, followID := range *tempList {
		var count int64
		err = DB.WithContext(ctx).Model(&Follow{}).
			Where("user_id = ? AND to_user_id = ? AND status = ?", followID, uid, constants.FollowAction).
			Count(&count).Error
		if err != nil {
			return nil, err
		} else if count == 0 { // 查无此纪录，说明不是好友，只是单方面关注
			continue
		}

		// 粉丝列表存在就直接添加这个id
		friendList = append(friendList, followID)
	}

	if len(friendList) == 0 { // db中也查不到
		return nil, RecordNotFound
	}

	return &friendList, nil
}

func FollowCount(ctx context.Context, uid int64) (int64, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&Follow{}).
		Where("user_id = ? AND status = ?", uid, constants.FollowAction).
		Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}

func FollowerCount(ctx context.Context, uid int64) (int64, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&Follow{}).
		Where("to_user_id = ? AND status = ?", uid, constants.FollowAction).
		Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}

func IsFollow(ctx context.Context, uid, tid int64) (bool, error) {
	var followModel *Follow

	err := DB.WithContext(ctx).
		Model(&Follow{}).
		Select("status").
		Where("user_id = ? AND to_user_id = ?", uid, tid).
		First(&followModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // db中也查不到
			return false, RecordNotFound
		}
		return false, err
	}

	return followModel.Status == int64(constants.FollowAction), nil
}
