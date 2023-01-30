package main

import (
	"context"
	user "github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user"
)

// TiktokUserServiceImpl implements the last service interface defined in the IDL.
type TiktokUserServiceImpl struct{}

// PingPong implements the TiktokUserServiceImpl interface.
func (s *TiktokUserServiceImpl) PingPong(ctx context.Context, req *user.Request1) (resp *user.Response, err error) {
	// TODO: Your code here...
	return
}
