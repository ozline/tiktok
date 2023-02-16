package model

import (
	"gorm.io/gorm"
)

// Follow 每个关系为一条记录，性能很拉，升级方案考虑使用 redis 进行储存，或者用 mq 维护一个队列
type (
	Follow struct {
		gorm.Model
		UserID   int64 `json:"user_id" gorm:"column:user_id;index"`
		TargetID int64 `json:"target_id" gorm:"column:user_id;index"`
	}
)

func AddRelation(userId, targetId int64) int {
	return 0
}

func RemoveRelation(userId, targetId int64) int {
	return 0
}

func GetFollowList(userId int64) error {

	// model.Db.Model(Follow{}).Where("user_id = ?", userId).Find()

	return nil
}

func GetFollowerList(userId int64) int {
	return 0
}

func GetFriendList(userId int64) int {
	return 0
}
