package cache

import "fmt"

var keyPattern = "%s::%d"

func GetVideoKey(videoID int64) string {
	return fmt.Sprintf(keyPattern, "like", videoID)
}

func GetUserKey(userID int64) string {
	return fmt.Sprintf(keyPattern, "like", userID)
}
