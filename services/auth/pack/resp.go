package pack

import (
	"errors"

	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/services/auth/kitex_gen/tiktok/auth"
)

func BuildBaseResp(err error) *auth.BaseResp {
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

func baseResp(err errno.ErrNo) *auth.BaseResp {
	return &auth.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}
