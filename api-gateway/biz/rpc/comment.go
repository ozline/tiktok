package rpc

import (
	"context"

	"github.com/ozline/tiktok/kitex_gen/tiktok/comment"
	"github.com/ozline/tiktok/pkg/errno"
)

// PostComment returns the result of operation T/F
func PostComment(ctx context.Context, req *comment.PostReq) (*comment.Comment, error) {
	resp, err := commentClient.Post(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Info.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Info.Code, resp.Info.Msg)
	}

	return resp.Comment, nil
}

// GetCommentList returns the list of comments
func GetCommentList(ctx context.Context, req *comment.ListReq) (list []*comment.Comment, err error) {
	resp, err := commentClient.List(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Info.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Info.Code, resp.Info.Msg)
	}

	return resp.Comments, nil
}

// LikeAction returns NOT result of operation but the Like status
func LikeAction(ctx context.Context, req *comment.LikeReq) (bool, error) {
	resp, err := commentClient.SetLike(ctx, req)

	if err != nil {
		return false, err
	}

	if resp.Info.Code != errno.SuccessCode {
		return false, errno.NewErrNo(resp.Info.Code, resp.Info.Msg)
	}

	return resp.IsLike, nil
}

// GetFavoriteList returns the list of users who like the comment
func GetFavoriteList(ctx context.Context, req *comment.FavoriteListReq) (list []int64, err error) {
	resp, err := commentClient.FavoriteList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Info.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Info.Code, resp.Info.Msg)
	}

	return resp.Videos, nil
}

func FavoriteAction(ctx context.Context, req *comment.FavoriteReq) error {
	resp, err := commentClient.SetFavorite(ctx, req)

	if err != nil {
		return err
	}

	if resp.Info.Code != errno.SuccessCode {
		return errno.NewErrNo(resp.Info.Code, resp.Info.Msg)
	}

	return nil
}

func GetVideoCountInfo(ctx context.Context, req *comment.GetVideoInfoReq) (*comment.GetVideoInfoResp, error) {
	resp, err := commentClient.GetVideoInfo(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Info.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Info.Code, resp.Info.Msg)
	}

	return resp, nil
}
