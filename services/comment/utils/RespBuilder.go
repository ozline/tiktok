package utils

import (
	"errors"

	"github.com/ozline/tiktok/kitex_gen/tiktok/comment"
	"github.com/ozline/tiktok/pkg/errno"
)

func PostRespBuilder(f func() (resp *comment.PostResp, err error)) (*comment.PostResp, error) {
	resp, err := f()
	if err == nil {
		resp.Info = baseResp(errno.Success)
		return resp, nil
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return &comment.PostResp{Info: baseResp(e)}, nil
	}

	s := errno.ServiceError.WithMessage(err.Error())
	return &comment.PostResp{Info: baseResp(s)}, nil
}

func ListRespBuilder(f func() (resp *comment.ListResp, err error)) (*comment.ListResp, error) {
	resp, err := f()
	if err == nil {
		resp.Info = baseResp(errno.Success)
		return resp, nil
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return &comment.ListResp{Info: baseResp(e)}, nil
	}

	s := errno.ServiceError.WithMessage(err.Error())
	return &comment.ListResp{Info: baseResp(s)}, nil
}

func LikeRespBuilder(f func() (resp *comment.LikeResp, err error)) (*comment.LikeResp, error) {
	resp, err := f()
	if err == nil {
		resp.Info = baseResp(errno.Success)
		return resp, nil
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return &comment.LikeResp{Info: baseResp(e)}, nil
	}

	s := errno.ServiceError.WithMessage(err.Error())
	return &comment.LikeResp{Info: baseResp(s)}, nil
}

func FavoriteRespBuilder(f func() (resp *comment.FavoriteResp, err error)) (*comment.FavoriteResp, error) {
	resp, err := f()
	if err == nil {
		resp.Info = baseResp(errno.Success)
		return resp, nil
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return &comment.FavoriteResp{Info: baseResp(e)}, nil
	}

	s := errno.ServiceError.WithMessage(err.Error())
	return &comment.FavoriteResp{Info: baseResp(s)}, nil
}

func FavoriteListRespBuilder(f func() (resp *comment.FavoriteListResp, err error)) (*comment.FavoriteListResp, error) {
	resp, err := f()
	if err == nil {
		resp.Info = baseResp(errno.Success)
		return resp, nil
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return &comment.FavoriteListResp{Info: baseResp(e)}, nil
	}

	s := errno.ServiceError.WithMessage(err.Error())
	return &comment.FavoriteListResp{Info: baseResp(s)}, nil
}

func GetVideoInfoRespBuilder(f func() (resp *comment.GetVideoInfoResp, err error)) (*comment.GetVideoInfoResp, error) {
	resp, err := f()

	if err == nil {
		resp.Info = baseResp(errno.Success)
		return resp, nil
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return &comment.GetVideoInfoResp{Info: baseResp(e)}, nil
	}

	s := errno.ServiceError.WithMessage(err.Error())
	return &comment.GetVideoInfoResp{Info: baseResp(s)}, nil
}

func baseResp(err errno.ErrNo) *comment.BaseInfo {
	return &comment.BaseInfo{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}
