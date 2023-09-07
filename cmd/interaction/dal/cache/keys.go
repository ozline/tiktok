package cache

import (
	"fmt"
	"strings"

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
	var builder strings.Builder
	builder.Grow(len(constants.CommentNXKey) + 1 + len(videoID))
	builder.WriteString(constants.CommentNXKey)
	builder.WriteString(":")
	builder.WriteString(videoID)
	// return fmt.Sprintf("%s:%s", constants.CommentNXKey, videoID)
	return builder.String()
}

func GetCountKey(videoID string) string {
	var builder strings.Builder
	builder.Grow(len(constants.CountKey) + 1 + len(videoID))
	builder.WriteString(constants.CountKey)
	builder.WriteString(":")
	builder.WriteString(videoID)
	// return fmt.Sprintf("%s:%s", constants.CountKey, videoID)
	return builder.String()
}

func GetCommentKey(videoID string) string {
	var builder strings.Builder
	builder.Grow(len(constants.CommentKey) + 1 + len(videoID))
	builder.WriteString(constants.CommentKey)
	builder.WriteString(":")
	builder.WriteString(videoID)
	// return fmt.Sprintf("%s:%s", constants.CommentKey, videoID)
	return builder.String()
}

func GetCountNXKey(videoID string) string {
	var builder strings.Builder
	builder.Grow(len(constants.CountNXKey) + 1 + len(videoID))
	builder.WriteString(constants.CountNXKey)
	builder.WriteString(":")
	builder.WriteString(videoID)
	// return fmt.Sprintf("%s:%s", constants.CountNXKey, videoID)
	return builder.String()
}
