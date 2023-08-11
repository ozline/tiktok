package service

import (
	"github.com/ozline/tiktok/cmd/user/dal/db"
	"github.com/ozline/tiktok/kitex_gen/user"
)

// GetUser check token and get user's info
func (s *UserService) GetUser(req *user.InfoRequest) (*db.User, error) {

	// 获取用户信息
	userModel, err := db.GetUserByID(s.ctx, req.UserId)

	if err != nil {
		return nil, err
	}

	return userModel, nil
}
