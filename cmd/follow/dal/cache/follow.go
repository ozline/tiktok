package cache

import (
	"context"
	"errors"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/redis/go-redis/v9"
)

// 判断关注数据是否存在于redis
func IsFollow(ctx context.Context, uid, tid int64) (bool, error) {
	exist, err := RedisClient.SIsMember(ctx, FollowListKey(uid), strconv.FormatInt(tid, 10)).Result()

	if err != nil {
		klog.Infof("err: %v", err)
		return exist, err
	}

	return exist, nil
}

func FollowAction(ctx context.Context, uid, tid int64) error {
	exist, err := IsFollow(ctx, uid, tid) // 查询数据是否存在于redis中
	if err != nil {
		return err
	} else if exist { // 存在说明已关注
		return errno.AlreadyFollowError
	}

	// 不存在，进行关注操作
	err = RedisClient.SAdd(ctx, FollowListKey(uid), strconv.FormatInt(tid, 10)).Err() // 自己的关注列表
	if err != nil {
		klog.Infof("err: %v", err)
		return err
	}

	err = RedisClient.SAdd(ctx, FollowerListKey(tid), strconv.FormatInt(uid, 10)).Err() // 对方的粉丝列表
	if err != nil {
		klog.Infof("err: %v", err)
		return err
	}
	return nil
}

func UnFollowAction(ctx context.Context, uid, tid int64) error {
	exist, err := IsFollow(ctx, uid, tid) // 查询数据是否存在于redis中
	if err != nil {
		return err
	} else if !exist { // redis中不存在,说明并未关注
		return errno.NotFollowError
	}

	// 从对方的粉丝列表移除
	err = RedisClient.SRem(ctx, FollowerListKey(tid), strconv.FormatInt(uid, 10)).Err()
	if err != nil {
		klog.Infof("err: %v", err)
		return err
	}

	// 从自己的关注列表移除
	err = RedisClient.SRem(ctx, FollowListKey(uid), strconv.FormatInt(tid, 10)).Err()
	if err != nil {
		klog.Infof("err: %v", err)
		return err
	}
	return nil
}

func FollowListAction(ctx context.Context, uid int64) (*[]int64, error) {
	followList := make([]int64, 0, 10)

	key := FollowListKey(uid)

	// 查询redis
	idList, err := RedisClient.SMembers(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
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

func UpdateFollowList(ctx context.Context, uid int64, followList *[]int64) error {
	key := FollowListKey(uid)
	idList := make([]interface{}, len(*followList))
	for i, v := range *followList {
		idList[i] = v
	}
	if err := RedisClient.SAdd(ctx, key, idList).Err(); err != nil {
		klog.Infof("err: %v", err)
		return err
	}
	return nil
}

func FollowerListAction(ctx context.Context, uid int64) (*[]int64, error) {
	followerList := make([]int64, 0, 10)

	key := FollowerListKey(uid)

	// 查询redis
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

func UpdateFollowerList(ctx context.Context, uid int64, followerList *[]int64) error {
	key := FollowerListKey(uid)
	idList := make([]interface{}, len(*followerList))
	for i, v := range *followerList {
		idList[i] = v
	}
	if err := RedisClient.SAdd(ctx, key, idList).Err(); err != nil {
		klog.Infof("err: %v", err)
		return err
	}
	return nil
}

func FriendListAction(ctx context.Context, uid int64) (*[]int64, error) {
	friendList := make([]int64, 0, 10)
	// 先获取本人的关注列表
	tempList, err := FollowListAction(ctx, uid)
	if err != nil {
		klog.Infof("err: %v", err)
		return nil, err
	}

	// 查询redis中的粉丝列表
	for _, followID := range *tempList {
		b, err := RedisClient.SIsMember(ctx, FollowerListKey(uid), followID).Result()
		if err != nil {
			klog.Infof("err: %v", err)
			return nil, err
		} else if !b { // 粉丝列表不存在，说明只是单方面关注，不是好友
			continue
		}
		friendList = append(friendList, followID)
	}

	return &friendList, nil
}

func UpdateFriendList(ctx context.Context, uid int64, followList, followerList *[]int64) error {
	userList, _ := FollowListAction(ctx, uid)
	if len(*userList) == 0 {
		err := UpdateFollowList(ctx, uid, followList)
		if err != nil {
			return err
		}
	}

	err := UpdateFollowerList(ctx, uid, followerList)
	if err != nil {
		return err
	}

	return nil
}

func FollowCount(ctx context.Context, uid int64) (int64, error) {
	key := FollowListKey(uid)
	size, err := RedisClient.SCard(ctx, key).Result()
	if err != nil {
		return -1, err
	}
	return size, nil
}

func FollowerCount(ctx context.Context, uid int64) (int64, error) {
	key := FollowerListKey(uid)
	size, err := RedisClient.SCard(ctx, key).Result()
	if err != nil {
		return -1, err
	}
	return size, nil
}
