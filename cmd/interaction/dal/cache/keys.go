package cache

import (
	"fmt"

	"github.com/ozline/tiktok/pkg/constants"
)

var keyPattern = "%s:%d"

func GetVideoLikeCountKey(videoID int64) string {
	return fmt.Sprintf(keyPattern, constants.VideoLikeCountKey, videoID)
}

func GetUserLikeKey(userID int64) string {
	return fmt.Sprintf(keyPattern, constants.UserLikeKey, userID)
}

func GetCommentNXKey(videoID string) string {
	return constants.CommentNXKey + ":" + videoID
}

func GetCountKey(videoID string) string {
	return constants.CountKey + ":" + videoID
}

func GetCommentKey(videoID string) string {
	return constants.CommentKey + ":" + videoID
}

func GetCountNXKey(videoID string) string {
	return constants.CountNXKey + ":" + videoID
}
