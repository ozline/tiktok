package model

import (
	"context"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
)

// 添加token进缓存
func AddToken(token string, user *User) error {
	data := map[string]interface{}{
		"UserId":        user.UserId,
		"Username":      user.Username,
		"Password":      user.Password,
		"FollowCount":   user.FollowCount,
		"FollowerCount": user.FollowerCount,
		"CreateDate":    user.CreateDate,
	}
	log.Info("AddToken: ", RedisDb, token, user, data)
	err := RedisDb.HMSet(context.Background(), token, data).Err()
	log.Info(err)
	return err
}
