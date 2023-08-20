package cache

import "fmt"

var keyPattern = "%s::%d"

func GetVideoKey(videoId int64) string {
	return fmt.Sprintf(keyPattern, "like", videoId)
}

func GetUserKey(userId int64) string {
	return fmt.Sprintf(keyPattern, "like", userId)
}
