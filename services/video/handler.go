package main

import (
	"context"
	"fmt"
	video "github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video"
)

// TiktokVideoServiceImpl implements the last service interface defined in the IDL.
type TiktokVideoServiceImpl struct{}

// PingPong implements the TiktokVideoServiceImpl interface.
func (s *TiktokVideoServiceImpl) PingPong(ctx context.Context, req *video.Request1) (resp *video.Response, err error) {
	// TODO: Your code here...
	resp = &video.Response{}
	resp.Message = req.Message
	fmt.Println("---- Server PingPong  ---")
	return
}
