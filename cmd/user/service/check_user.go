package service

import (
	"github.com/ozline/tiktok/cmd/user/dal/db"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/pkg/errno"
)

// CheckUser check user is exist and it's password
func (s *UserService) CheckUser(req *user.LoginRequest) (*db.User, error) {
	userModel, err := db.GetUserByUsername(s.ctx, req.Username)

	if err != nil {
		return nil, err
	}

	if req.Password != userModel.Password {
		return nil, errno.AuthorizationFailedError
	}

	return userModel, nil
}
