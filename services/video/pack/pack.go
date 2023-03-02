package pack

import (
	"errors"

	"github.com/ozline/tiktok/kitex_gen/tiktok/video"
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
