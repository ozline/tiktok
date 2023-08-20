package cache

import (
	"context"
	"errors"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
)

// 判断关注数据是否存在于redis
func IsFollow(ctx context.Context, uid, followId int64) (bool, error) {
	tid := strconv.FormatInt(followId, 10)
	b, err := RedisClient.SIsMember(ctx, FollowListKey(uid), tid).Result()
	if err != nil {
		klog.Infof("err: %v", err)
		return b, err
	}

	return b, nil
}

func FollowAction(ctx context.Context, uid, tid int64) error {
	b, err := IsFollow(ctx, uid, tid) //查询数据是否存在于redis中
	if err != nil {
		return err
	} else if b { //存在说明已关注
		return errors.New("you already follow this user")
	}

	//不存在，进行关注操作
	err = RedisClient.SAdd(ctx, FollowListKey(uid), strconv.FormatInt(tid, 10)).Err() //自己的关注列表
	if err != nil {
		klog.Infof("err: %v", err)
		return err
	}

	err = RedisClient.SAdd(ctx, FollowerListKey(tid), strconv.FormatInt(uid, 10)).Err() //对方的粉丝列表
	if err != nil {
		klog.Infof("err: %v", err)
		return err
	}
	return nil
}

func UnFollowAction(ctx context.Context, uid, tid int64) error {
	b, err := IsFollow(ctx, uid, tid) //查询数据是否存在于redis中
	if err != nil {
		return err
	} else if !b { //redis中不存在,说明并未关注
		return errors.New("you are not following this user")
	}

	//从对方的粉丝列表移除
	err = RedisClient.SRem(ctx, FollowerListKey(tid), strconv.FormatInt(uid, 10)).Err()
	if err != nil {
		klog.Infof("err: %v", err)
		return err
	}

	//从自己的关注列表移除
	err = RedisClient.SRem(ctx, FollowListKey(uid), strconv.FormatInt(tid, 10)).Err()
	if err != nil {
		klog.Infof("err: %v", err)
		return err
	}
	return nil
}

func FollowListAction(ctx context.Context, uid int64) (*[]int64, error) {
	var followList []int64

	key := FollowListKey(uid)

	//查询redis
	idList, err := RedisClient.SMembers(ctx, key).Result()
	if err == redis.Nil {
		klog.Info("Not found followList")
	} else if err != nil {
		klog.Infof("err: %v", err)
		return nil, err
	}

	for _, id := range idList {
		followId, _ := strconv.ParseInt(id, 10, 64)
		followList = append(followList, followId)
	}
	return &followList, nil
}

func FollowerListAction(ctx context.Context, uid int64) (*[]int64, error) {
	var followerList []int64

	key := FollowerListKey(uid)

	//查询redis
	idList, err := RedisClient.SMembers(ctx, key).Result()
	if err == redis.Nil {
		klog.Info("Not found followerList")
	} else if err != nil {
		klog.Infof("err: %v", err)
		return nil, err
	}

	for _, id := range idList {
		followerId, _ := strconv.ParseInt(id, 10, 64)
		followerList = append(followerList, followerId)
	}
	return &followerList, nil
}

func FriendListAction(ctx context.Context, uid int64) (*[]int64, error) {
	var friendList []int64
	//先获取本人的关注列表
	tempList, err := FollowListAction(ctx, uid)
	if err != nil {
		klog.Infof("err: %v", err)
		return nil, err
	}

	//查询redis中的粉丝列表
	for _, id := range *tempList {
		b, err := RedisClient.SIsMember(ctx, FollowerListKey(uid), id).Result()
		if err != nil {
			klog.Infof("err: %v", err)
			return nil, err
		} else if !b { //粉丝列表不存在，说明只是单方面关注，不是好友
			continue
		}
		friendList = append(friendList, id)
	}
	return &friendList, nil
}
