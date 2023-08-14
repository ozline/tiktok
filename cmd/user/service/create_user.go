package service

import (
	"github.com/ozline/tiktok/cmd/user/dal/db"
	"github.com/ozline/tiktok/kitex_gen/user"
)

// CreateUser create user info
func (s *UserService) CreateUser(req *user.RegisterRequest) (*db.User, error) {
	userModel := &db.User{
		Username: req.Username,
		Password: req.Password,
	}

	return db.CreateUser(s.ctx, userModel)
}
