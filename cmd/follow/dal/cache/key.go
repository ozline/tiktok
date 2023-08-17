package cache

import "strconv"

const (
	FollowKey   = "follow:"
	FollowerKey = "follower:"
)

func FollowListKey(uid int64) string {
	return FollowKey + strconv.FormatInt(uid, 10)
}

func FollowerListKey(uid int64) string {
	return FollowerKey + strconv.FormatInt(uid, 10)
}
