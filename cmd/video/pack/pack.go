package pack

import (
	"errors"

	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/pkg/errno"
)

func BuildBaseResp(err error) *video.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceError.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *video.BaseResp {
	return &video.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}

// func BuildPutVideoResp(err error, state int64) *video.PutVideoResponse {
// 	if err == nil {
// 		return putVideoResp(errno.Success, state)
// 	}
// 	e := errno.ErrNo{}

// 	if errors.As(err, &e) {
// 		return putVideoResp(e, state)
// 	}
// 	s := errno.ServiceError.WithMessage(err.Error())
// 	return putVideoResp(s, state)
// }

// func putVideoResp(err errno.ErrNo, state int64) *video.PutVideoResponse {
// 	return &video.PutVideoResponse{
// 		Base:  BuildBaseResp(err),
// 		State: state,
// 	}
// }
