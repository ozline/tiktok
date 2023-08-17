package main

import (
	"context"
	"github.com/ozline/tiktok/cmd/interactive/dal/db"
	"github.com/ozline/tiktok/cmd/interactive/pack"
	"github.com/ozline/tiktok/cmd/interactive/rpc"
	"github.com/ozline/tiktok/cmd/interactive/service"
	interactive "github.com/ozline/tiktok/kitex_gen/interactive"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/pkg/constants"
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

	commentResp := new(db.Comment)
	commentService := service.NewCommentService(ctx)

	switch req.ActionType {
	//1-发布评论
	case constants.AddComment:
		if len(*req.CommentText) == 0 {
			resp.Base = pack.BuildBaseResp(errno.ParamError)
			return resp, nil
		}

		ok, err := commentService.MatchSensitiveWords(*req.CommentText)
		if err != nil {
			resp.Base = pack.BuildBaseResp(errno.SensitiveWordsHTTPError)
			return resp, nil
		}
		if ok {
			resp.Base = pack.BuildBaseResp(errno.SensitiveWordsError)
			return resp, nil
		}

		commentResp, err = commentService.CreateComment(req)

		if err != nil {
			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		}
	//2-删除评论
	case constants.DeleteComment:
		commentResp, err = commentService.DeleteComment(req)

		if err != nil {
			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		}

	default:
		resp.Base = pack.BuildBaseResp(errno.UnexpectedTypeError)
		return resp, nil
	}

	userInfo, err := rpc.UserInfo(ctx, &user.InfoRequest{
		UserId: commentResp.UserId,
		Token:  req.Token,
	})
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.Comment = pack.Comment(commentResp)
	resp.Comment.User = userInfo

	return
}

// CommentList implements the InteractiveServiceImpl interface.
func (s *InteractiveServiceImpl) CommentList(ctx context.Context, req *interactive.CommentListRequest) (resp *interactive.CommentListResponse, err error) {
	resp = new(interactive.CommentListResponse)

	// 校验token
	if _, err = utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}

	commentsResp, err := service.NewCommentService(ctx).GetComments(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	users := make(map[int64]int) // 利用map避免重复查询
	commentList := make([]*interactive.Comment, 0, len(*commentsResp))
	for commentIndex, comment := range *commentsResp {
		rComment := new(interactive.Comment)
		rComment = pack.Comment(&comment)

		index, ok := users[comment.UserId]
		if !ok {
			userInfo, err := rpc.UserInfo(ctx, &user.InfoRequest{
				UserId: comment.UserId,
				Token:  req.Token,
			})
			if err != nil {
				resp.Base = pack.BuildBaseResp(err)
				return resp, nil
			}
			rComment.User = userInfo
			users[comment.UserId] = commentIndex
		} else {
			rComment.User = commentList[index].User
		}

		commentList = append(commentList, rComment)
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.CommentList = commentList
	return
}
