package cache

import (
	"fmt"

	"github.com/ozline/tiktok/pkg/constants"
)

var keyPattern = "%s::%d"

func GetVideoKey(videoID int64) string {
	return fmt.Sprintf(keyPattern, "favorited", videoID)
}

func GetUserKey(userID int64) string {
	return fmt.Sprintf(keyPattern, "favorite", userID)
}

func GetCommentNXKey(videoID string) string {
	return fmt.Sprintf("%s:%s", constants.CommentNXKey, videoID)
}

func GetCountKey(videoID string) string {
	return fmt.Sprintf("%s:%s", constants.CountKey, videoID)
}

func GetCommentKey(videoID string) string {
	return fmt.Sprintf("%s:%s", constants.CommentKey, videoID)
}

func GetCountNXKey(videoID string) string {
	return fmt.Sprintf("%s:%s", constants.CountNXKey, videoID)
}
