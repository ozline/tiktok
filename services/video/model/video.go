package model

import "sync"

type Video struct {
	M             sync.Mutex //
	ID            int64      // 视频ID
	Author        *User      // 作者信息
	PlayUrl       string     // 播放地址
	CoverUrl      string     // 封面地址
	FavoriteCount int64      // 视频的点赞总数
	CommentCount  int64      // 视频的评论总数
	IsFavorite    bool       // 是否本人已点赞
	Title         string     // 标题
}
type VideoStorageInfo struct {
	VideoID         int64  // 视频id
	VideoPlayUrl    string // 视频标题
	VideoCoverUrl   string
	VideoTitle      string
	VideoCreateTime string // 视频创建时间
	UserName        string // 用户名称
}

type User struct {
	Id            int64  // 用户id
	Name          string // 用户名称
	FollowCount   int64  // 关注总数
	FollowerCount int64  // 粉丝总数
	IsFollow      bool   // true-已关注，false-未关注
}
