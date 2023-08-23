package cache

import "fmt"

var keyPattern = "%s::%d"

func GetVideoKey(videoID int64) string {
	return fmt.Sprintf(keyPattern, "favorited", videoID)
}

func GetUserKey(userID int64) string {
	return fmt.Sprintf(keyPattern, "favorite", userID)
}
