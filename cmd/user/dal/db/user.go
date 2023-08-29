package db

import (
	"context"
	"errors"
	"time"

	"github.com/ozline/tiktok/pkg/errno"
	"gorm.io/gorm"
)

type User struct {
	Id              int64
	Username        string
	Password        string
	Avatar          string `gorm:"default:https://files.ozline.icu/images/avatar.jpg"`
	BackgroundImage string `gorm:"default:https://files.ozline.icu/images/BannerImg_221116.jpeg"`
	Signature       string `gorm:"default:NOT NULL BUT SEEMS NULL"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

/*

TODO: follow_count, follower_count, is_follow, work_count, favorite_count, total_favorited(string)

很迷的, 我觉得很多东西不需要

*/

func CreateUser(ctx context.Context, user *User) (*User, error) {
	userResp := new(User)

	err := DB.WithContext(ctx).Where("username = ?", user.Username).First(&userResp).Error

	if err == nil {
		return nil, errno.UserExistedError
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err := DB.WithContext(ctx).Create(user).Error; err != nil {
		// add some logs
		return nil, err
	}

	return user, nil
}

func GetUserByUsername(ctx context.Context, username string) (*User, error) {
	userResp := new(User)

	err := DB.WithContext(ctx).Where("username = ?", username).First(&userResp).Error

	if err != nil {
		// add some logs

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return userResp, nil
}

func GetUserByID(ctx context.Context, userid int64) (*User, error) {
	userResp := new(User)

	err := DB.WithContext(ctx).Where("id = ?", userid).First(&userResp).Error

	if err != nil {
		// add some logs

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.UserNotFoundError
		}
		return nil, err
	}

	return userResp, nil
}
