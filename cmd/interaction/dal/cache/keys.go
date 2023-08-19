package cache

import "fmt"

var keyPattern = "%s::%s"

func GetVideoKey(videoId string) string {
	return fmt.Sprintf(keyPattern, "video", videoId)
}
