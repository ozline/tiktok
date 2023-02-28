package model

import (
	"errors"
	"log"

	"github.com/ozline/tiktok/pkg/constants"
	"gorm.io/gorm"
)

type (
	Follow struct {
		gorm.Model
		UserId   int64 `json:"user_id,omitempty" gorm:"column:user_id;index"`
		ToUserId int64 `json:"to_user_id,omitempty" gorm:"column:to_user_id;index"`
	}

	// User 查询到的用户信息，微服务应该怎么拆分？
	// 个人理解的话，follow 的这几个接口应该都做到用户服务里，即便拆分两个服务的数据也过度耦合
	User struct {
		ID            int64  `json:"id,omitempty"`
		Username      string `json:"username,omitempty"`
		FollowCount   int64  `json:"follow_count,omitempty"`
		FollowerCount int64  `json:"follower_count,omitempty"`
		IsFollow      bool   `json:"is_follow,omitempty"`
	}
)

func AddRelation(userId, toUserId int64) error {
	var (
		relation Follow
	)

	if err := db.Model(Follow{}).
		Where("user_id = ? AND to_user_id", userId, toUserId).
		First(&relation).Error;
	// if there is no record for this relationship, create;
	// or find record, return error;
	// or other error, return error.
	err == gorm.ErrRecordNotFound {
		// insert new relation to database.
		if err := db.Create(&relation).Error; err != nil {
			// TODO: import a global log system
			log.Println(err)
			return errors.New("database error")
		} else {
			return nil
		}
	} else if err == nil {
		return errors.New("the relation have been existed")
	} else {
		return errors.New("database error")
	}
}

func DeleteRelation(userId, toUserId int64) error {
	var (
		relation Follow
	)

	if err := db.Model(Follow{}).
		Where("user_id = ? AND to_user_id", userId, toUserId).
		First(&relation).Error;
	// if there is no record for this relationship, error no record;
	// or find record, deleter;
	// or other error, return error.
	err != nil {
		// insert new relation to database.
		if err := db.Delete(&relation).Error; err != nil {
			// TODO: import a global log system
			log.Println(err)
			return errors.New("database error")
		} else {
			return nil
		}
	} else if err == gorm.ErrRecordNotFound {
		return errors.New("this relation is not exist")
	} else {
		return errors.New("database error")
	}
}

func IsFollowRelation(userId, toUserId int64) bool {
	var (
		count int64
	)

	if err := db.Table(constants.FollowTableName).Where("user_id = ? AND to_user_id = ?", userId, toUserId).Count(&count).Error; err != nil {
		log.Println("[FollowService]", err)
		return false
	} else {
		return count > 0
	}
}

func IsFollowerRelation(userId, toUserId int64) bool {
	var (
		count int64
	)

	if err := db.Table(constants.FollowTableName).Where("user_id = ? AND to_user_id = ?", toUserId, userId).Count(&count).Error; err != nil {
		log.Println("[FollowService]", err)
		return false
	} else {
		return count > 0
	}
}

func QueryFollowsList(userId int64, pageNum int32, pageSize int32) ([]User, error) {
	var (
		data         []User
		followIdList []int64
	)

	if err := db.Table(constants.FollowTableName).
		Select("to_user_id").
		Where("user_id = ?", userId).
		Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).
		Find(&followIdList).Error; err != nil {
		return nil, err
	} else {
		db.Table(constants.UserTableName).Where("id = ?", followIdList).Find(&data)
		return data, nil
	}
}

func QueryFollowersList(userId int64, pageNum int32, pageSize int32) ([]User, error) {
	var (
		data           []User
		followerIdList []int64
	)

	if err := db.Table(constants.FollowTableName).
		Select("user_id").
		Where("to_user_id = ?", userId).
		Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).
		Find(&followerIdList).Error; err != nil {
		return nil, err
	} else {
		db.Table(constants.UserTableName).Where("user_id = ?", followerIdList).Find(&data)
		return data, nil
	}
}

func QueryFriendsList(userId int64, pageNum int32, pageSize int32) ([]User, error) {
	var (
		data         []User
		friendIdList []int64
		tempList     []int64
	)

	if err := db.Table(constants.FollowTableName).
		Select("to_user_id").
		Where("user_id = ?", userId).
		Find(&tempList).Error; err != nil {
		return nil, err
	}

	for i := 0; i < len(tempList); i++ {
		if IsFollowerRelation(userId, tempList[i]) {
			friendIdList = append(friendIdList, tempList[i])
		}
	}

	if err := db.Table(constants.UserTableName).
		Where("id = ?", userId).
		Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
