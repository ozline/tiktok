package db

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/ozline/tiktok/cmd/follow/dal/cache"
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

var r = cache.RedisClient

func FollowAction(ctx context.Context, follow *Follow) error {
	followResp := new(Follow)

	//TODO:redis限流
	//TODO:mysql的定时写入

	tid := strconv.FormatInt(follow.ToUserId, 10)
	uid := strconv.FormatInt(follow.UserId, 10)
	//判断数据是否存在于redis(存在为已关注，不存在为已取关)
	b, err := r.SIsMember(ctx, cache.FollowListKey(follow.UserId), tid).Result()
	if err != nil {
		return err
	}
	//数据存在,进行取关
	if b {
		err = DB.WithContext(ctx).Model(&Follow{}).
			Where("user_id= ? AND to_user_id = ?", follow.UserId, follow.ToUserId).
			Update("action_type", follow.ActionType).Error
		if err != nil {
			return err
		}
		//从对方的粉丝列表移除
		r.SRem(ctx, cache.FollowerListKey(follow.ToUserId), uid).Err()
		if err != nil {
			return err
		}
		return r.SRem(ctx, cache.FollowListKey(follow.UserId), tid).Err()
	}

	//数据不存在，添加数据进行关注
	err = r.SAdd(ctx, cache.FollowerListKey(follow.UserId), tid).Err() //自己的关注列表
	if err != nil {
		return err
	}

	err = r.SAdd(ctx, cache.FollowerListKey(follow.ToUserId), uid).Err() //对方的粉丝列表
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
	} else if err == nil { //db中存在,修改ActionType
		return DB.WithContext(ctx).Model(&Follow{}).
			Where("user_id= ? AND to_user_id = ?", follow.UserId, follow.ToUserId).
			Update("action_type", follow.ActionType).Error
	}
	return err
}

// 关注列表(获取to_user_id的列表)
func FollowListAction(ctx context.Context, uid int64) (*[]int64, error) {
	var followList []int64
	//TODO:redis限流
	//TODO:mysql的定时写入
	key := cache.FollowListKey(uid)

	idList, err := r.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if len(idList) == 0 {
		err := DB.WithContext(ctx).Model(&Follow{}).Select("to_user_id").
			Where("user_id = ? AND action_type = ?", uid, constants.FollowAction).
			Find(&followList).Error
		if err != nil {
			return nil, err
		}
	} else {
		for i, id := range idList {
			followList[i], _ = strconv.ParseInt(id, 10, 64)
		}
	}

	return &followList, nil
}

// 粉丝列表(获取user_id的列表)
func FollowerListAction(ctx context.Context, uid int64) (*[]int64, error) {
	var followerList []int64
	//TODO:redis限流
	//TODO:mysql的定时写入

	key := cache.FollowerListKey(uid)

	idList, err := r.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if len(idList) == 0 {
		err := DB.WithContext(ctx).Model(&Follow{}).Select("user_id").
			Where("to_user_id = ? AND action_type = ?", uid, constants.FollowAction).
			Find(&followerList).Error
		if err != nil {
			return nil, err
		}
	} else {
		for i, id := range idList {
			followerList[i], _ = strconv.ParseInt(id, 10, 64)
		}
	}

	return &followerList, nil
}

// 好友列表(获取to_user_id的列表)
func FriendListAction(ctx context.Context, uid int64) (*[]int64, error) {
	var tempList []int64
	var friendList []int64

	//TODO:redis限流
	//TODO:mysql的定时写入

	//先获取本人关注的列表
	key := cache.FollowListKey(uid)
	idList, err := r.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if len(idList) != 0 {
		for _, id := range idList {
			//查询粉丝列表
			b, err := r.SIsMember(ctx, cache.FollowerListKey(uid), id).Result()
			if err != nil {
				return nil, err
			}
			if !b {
				continue
			}
			fid, _ := strconv.ParseInt(id, 10, 64)
			friendList = append(friendList, fid)
		}
		return &friendList, nil
	}

	//从db中获取
	err = DB.WithContext(ctx).Model(&Follow{}).Select("to_user_id").
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
