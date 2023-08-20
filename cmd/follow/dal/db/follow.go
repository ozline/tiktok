package db

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/follow/dal/cache"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Follow struct {
	Id       int64
	Status   int64 //1-关注, 2-取消关注
	UserId   int64
	ToUserId int64
	// ActionType int64 `gorm:"default:1"` //1-关注, 2-取消关注
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FollowAction(ctx context.Context, follow *Follow) error {
	if err := cache.Limit(ctx); err != nil {
		return err
	}

	r := cache.RedisClient
	followResp := new(Follow)

	tid := strconv.FormatInt(follow.ToUserId, 10)
	uid := strconv.FormatInt(follow.UserId, 10)

	//判断数据是否存在于redis(存在说明已关注)
	b, err := r.SIsMember(ctx, cache.FollowListKey(follow.UserId), tid).Result()
	if err != nil {
		return err
	} else if b { //redis中存在
		return errors.New("you already follow this user")
	}

	//不存在，进行关注操作
	err = r.SAdd(ctx, cache.FollowListKey(follow.UserId), tid).Err() //自己的关注列表
	if err != nil {
		return err
	}

	err = r.SAdd(ctx, cache.FollowerListKey(follow.ToUserId), uid).Err() //对方的粉丝列表
	if err != nil {
		return err
	}

	//查询db中是否存在记录(判断是否是第一次操作)
	err = DB.WithContext(ctx).Model(&Follow{}).
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
	if err := cache.Limit(ctx); err != nil {
		return err
	}

	r := cache.RedisClient

	tid := strconv.FormatInt(follow.ToUserId, 10)
	uid := strconv.FormatInt(follow.UserId, 10)

	//判断数据是否存在于redis(不存在说明并未关注)
	b, err := r.SIsMember(ctx, cache.FollowListKey(follow.UserId), tid).Result()
	if err != nil {
		return err
	} else if !b { //redis中不存在
		return errors.New("you are not following this user")
	}

	//从对方的粉丝列表移除
	err = r.SRem(ctx, cache.FollowerListKey(follow.ToUserId), uid).Err()
	if err != nil {
		return err
	}

	//从自己的关注列表移除
	err = r.SRem(ctx, cache.FollowListKey(follow.UserId), tid).Err()
	if err != nil {
		return err
	}

	//修改db中的status
	err = DB.WithContext(ctx).Model(&Follow{}).
		Where("user_id= ? AND to_user_id = ?", follow.UserId, follow.ToUserId).
		Update("status", constants.UnFollowAction).Error
	if err != nil {
		return err
	}

	return nil
}

// 关注列表(获取to_user_id的列表)
func FollowListAction(ctx context.Context, uid int64) (*[]int64, error) {
	if err := cache.Limit(ctx); err != nil {
		return nil, err
	}

	r := cache.RedisClient
	var followList []int64

	key := cache.FollowListKey(uid)

	//查询redis
	idList, err := r.SMembers(ctx, key).Result()
	if err == redis.Nil {
		klog.Info("Not found followList")
	} else if err != nil {
		return nil, err
	}

	if len(idList) == 0 { //redis中不存在,查询db并返回结果
		err := DB.WithContext(ctx).Table(constants.FollowTableName).Select("to_user_id").
			Where("user_id = ? AND status = ?", uid, constants.FollowAction).
			Find(&followList).Error
		if err != nil {
			return nil, err
		}
	} else { //redis中存在,处理redis的数据
		for _, id := range idList {
			followId, _ := strconv.ParseInt(id, 10, 64)
			followList = append(followList, followId)
		}
	}

	return &followList, nil
}

// 粉丝列表(获取user_id的列表)
func FollowerListAction(ctx context.Context, uid int64) (*[]int64, error) {
	if err := cache.Limit(ctx); err != nil {
		return nil, err
	}

	r := cache.RedisClient
	var followerList []int64

	key := cache.FollowerListKey(uid)

	//查询redis
	idList, err := r.SMembers(ctx, key).Result()
	if err == redis.Nil {
		klog.Info("Not found followerList")
	} else if err != nil {
		return nil, err
	}

	if len(idList) == 0 { //redis中不存在,查询db并返回结果
		err := DB.WithContext(ctx).Table(constants.FollowTableName).Select("user_id").
			Where("to_user_id = ? AND status = ?", uid, constants.FollowAction).
			Find(&followerList).Error
		if err != nil {
			return nil, err
		}
	} else { //redis中存在,处理redis的数据
		for _, id := range idList {
			followerId, _ := strconv.ParseInt(id, 10, 64)
			followerList = append(followerList, followerId)
		}
	}

	return &followerList, nil
}

// 好友列表(先获取to_user_id的列表)
func FriendListAction(ctx context.Context, uid int64) (*[]int64, error) {
	if err := cache.Limit(ctx); err != nil {
		return nil, err
	}

	r := cache.RedisClient
	// var tempList []int64
	var friendList []int64

	//先获取本人的关注列表
	tempList, err := FollowListAction(ctx, uid)
	if err != nil {
		return nil, err
	}

	//查询redis中的粉丝列表

	for _, id := range *tempList {
		b, err := r.SIsMember(ctx, cache.FollowerListKey(uid), id).Result()
		if err != nil {
			return nil, err
		} else if !b { //redis中粉丝列表不存在，查询db记录
			err = DB.WithContext(ctx).Model(&Follow{}).
				Where("user_id = ? AND to_user_id = ? AND status = ?", id, uid, constants.FollowAction).Error
			if errors.Is(err, gorm.ErrRecordNotFound) { //db中也不存在,说明只是单方面关注，不是好友
				continue
			} else if err != nil {
				return nil, err
			}
		}
		//粉丝列表存在就直接添加这个id
		friendList = append(friendList, id)
	}

	return &friendList, nil
}
