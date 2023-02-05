package main

import (
	"context"
	auth "github.com/ozline/tiktok/services/auth/kitex_gen/tiktok/auth"
)

// TiktokAuthServiceImpl implements the last service interface defined in the IDL.
type TiktokAuthServiceImpl struct{}

// Ping implements the TiktokAuthServiceImpl interface.
func (s *TiktokAuthServiceImpl) Ping(ctx context.Context, req *auth.PingRequest) (resp *auth.PingResponse, err error) {
	// TODO: Your code here...
	return
}

// GetToken implements the TiktokAuthServiceImpl interface.
func (s *TiktokAuthServiceImpl) GetToken(ctx context.Context, req *auth.GetTokenRequest) (resp *auth.GetTokenResponse, err error) {
	// TODO: Your code here...
	return
}

// CheckToken implements the TiktokAuthServiceImpl interface.
func (s *TiktokAuthServiceImpl) CheckToken(ctx context.Context, req *auth.CheckTokenRequest) (resp *auth.CheckTokenResponse, err error) {
	// TODO: Your code here...
	return
}
