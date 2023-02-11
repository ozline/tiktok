package rpc

import (
	"context"

	"github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video"
)

// func VideoGetList(ctx context.Context, req *video.GetVideoListRequest) (*video.GetVideoListResponse, error) {
// 	resp, err := videoClient.GetList(ctx, req)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp, nil
// }

func VideoUpload(ctx context.Context, req *video.PutVideoRequest) (bool, error) {
	resp, err := videoClient.PutVideo(ctx, req)

	if err != nil {
		return false, err
	}

	//TODO: 完善错误

	resp.ErrState = ""

	return true, nil
}
