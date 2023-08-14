package main

import (
	"context"
	"github.com/ozline/tiktok/cmd/interactive/dal/db"
	"github.com/ozline/tiktok/cmd/interactive/pack"
	"github.com/ozline/tiktok/cmd/interactive/rpc"
	"github.com/ozline/tiktok/cmd/interactive/service"
	interactive "github.com/ozline/tiktok/kitex_gen/interactive"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

// InteractiveServiceImpl implements the last service interface defined in the IDL.
type InteractiveServiceImpl struct{}

// FavoriteAction implements the InteractiveServiceImpl interface.
func (s *InteractiveServiceImpl) FavoriteAction(ctx context.Context, req *interactive.FavoriteActionRequest) (resp *interactive.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the InteractiveServiceImpl interface.
func (s *InteractiveServiceImpl) FavoriteList(ctx context.Context, req *interactive.FavoriteListRequest) (resp *interactive.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentAction implements the InteractiveServiceImpl interface.
func (s *InteractiveServiceImpl) CommentAction(ctx context.Context, req *interactive.CommentActionRequest) (resp *interactive.CommentActionResponse, err error) {
	resp = new(interactive.CommentActionResponse)

	if (req.ActionType != "1" && req.ActionType != "2") || (req.ActionType == "1" && len(*req.CommentText) == 0) {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}

	commentResp := new(db.Comment)

	//1-发布评论
	if req.ActionType == "1" {
		commentResp, err = service.NewCommentService(ctx).CreateComment(req)

		if err != nil {
			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		}

		//2-删除评论
	} else if req.ActionType == "2" {
		commentResp, err = service.NewCommentService(ctx).DeleteComment(req)

		if err != nil {
			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		}
	}

	user, err := rpc.UserInfo(ctx, &user.InfoRequest{
		UserId: commentResp.UserId,
		Token:  req.Token,
	})
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.Comment = pack.Comment(commentResp)
	resp.Comment.User = user

	return
}

// CommentList implements the InteractiveServiceImpl interface.
func (s *InteractiveServiceImpl) CommentList(ctx context.Context, req *interactive.CommentListRequest) (resp *interactive.CommentListResponse, err error) {
	resp = new(interactive.CommentListResponse)

	// 校验token
	if _, err := utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}

	commentResp, err := service.NewCommentService(ctx).GetComments(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.CommentList = pack.Comments(commentResp)
	for index, comment := range resp.CommentList {
		user, err := rpc.UserInfo(ctx, &user.InfoRequest{
			UserId: (*commentResp)[index].UserId,
			Token:  req.Token,
		})
		if err != nil {
			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		}
		comment.User = user
	}

	return
}
