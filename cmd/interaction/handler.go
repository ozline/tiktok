package main

import (
	"context"

	"github.com/ozline/tiktok/cmd/interaction/pack"
	"github.com/ozline/tiktok/cmd/interaction/rpc"
	"github.com/ozline/tiktok/cmd/interaction/service"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	interaction "github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

// InteractionServiceImpl implements the last service interface defined in the IDL.
type InteractionServiceImpl struct{}

// FavoriteAction implements the interactionServiceImpl interface.
func (s *InteractionServiceImpl) FavoriteAction(ctx context.Context, req *interaction.FavoriteActionRequest) (resp *interaction.FavoriteActionResponse, err error) {
	resp = new(interaction.FavoriteActionResponse)

	claims, err := utils.CheckToken(req.Token)
	if err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, err
	}

	if req.ActionType != 1 && req.ActionType != 2 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}

	switch req.ActionType {
	// 1 like
	case constants.Like:
		if err := service.NewInteractionService(ctx).Like(req, claims.UserId); err != nil {
			resp.Base = pack.BuildBaseResp(err)
			return resp, err
		}
	// 2 dislike
	case constants.Dislike:
		if err := service.NewInteractionService(ctx).Dislike(req, claims.UserId); err != nil {
			resp.Base = pack.BuildBaseResp(err)
			return resp, err
		}
	}

	resp.Base = pack.BuildBaseResp(nil)
	return resp, nil
}

// FavoriteList implements the interactionServiceImpl interface.
func (s *InteractionServiceImpl) FavoriteList(ctx context.Context, req *interaction.FavoriteListRequest) (resp *interaction.FavoriteListResponse, err error) {
	resp = new(interaction.FavoriteListResponse)

	if _, err := utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, err
	}

	videoIdList, err := service.NewInteractionService(ctx).FavoriteList(req)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, err
	}

	videos, err := rpc.GetFavoriteVideoList(ctx, &video.GetFavoriteVideoInfoRequest{
		VideoId: videoIdList,
		Token:   req.Token,
	})

	resp.Base = pack.BuildBaseResp(nil)
	resp.VideoList = pack.BuildVideos(videos)
	return
}

// CommentAction implements the interactionServiceImpl interface.
func (s *InteractionServiceImpl) CommentAction(ctx context.Context, req *interaction.CommentActionRequest) (resp *interaction.CommentActionResponse, err error) {
	resp = new(interaction.CommentActionResponse)

	commentService := service.NewInteractionService(ctx)

	claim, err := utils.CheckToken(req.Token)
	if err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}
	userId := claim.UserId

	switch req.ActionType {
	//1-发布评论
	case constants.AddComment:

		if req.CommentText == nil || len(*req.CommentText) == 0 || len(*req.CommentText) > 255 {
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

		commentResp, err := commentService.CreateComment(req, userId)

		if err != nil {
			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		}

		resp.Comment = commentResp

	//2-删除评论
	case constants.DeleteComment:

		if req.CommentId == nil {
			resp.Base = pack.BuildBaseResp(errno.ParamError)
			return resp, nil
		}

		commentResp, err := commentService.DeleteComment(req, userId)

		if err != nil {
			resp.Base = pack.BuildBaseResp(err)
			return resp, nil
		}
		resp.Comment = commentResp

	default:
		resp.Base = pack.BuildBaseResp(errno.UnexpectedTypeError)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)

	return resp, nil
}

// CommentList implements the interactionServiceImpl interface.
func (s *InteractionServiceImpl) CommentList(ctx context.Context, req *interaction.CommentListRequest) (resp *interaction.CommentListResponse, err error) {
	resp = new(interaction.CommentListResponse)

	// 校验token
	if _, err = utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}

	commentsResp, err := service.NewInteractionService(ctx).GetComments(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.CommentList = commentsResp
	return
}

// FavoriteCount implements the interactionServiceImpl interface.
func (s *InteractionServiceImpl) FavoriteCount(ctx context.Context, req *interaction.FavoriteCountRequest) (resp *interaction.FavoriteCountResponse, err error) {
	resp = new(interaction.FavoriteCountResponse)

	if _, err := utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, err
	}

	likeCount, err := service.NewInteractionService(ctx).GetLikeCount(req)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, err
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.LikeCount = likeCount
	return
}

// CommentCount implements the interactionServiceImpl interface.
func (s *InteractionServiceImpl) CommentCount(ctx context.Context, req *interaction.CommentCountRequest) (resp *interaction.CommentCountResponse, err error) {
	resp = new(interaction.CommentCountResponse)

	// 校验token
	if req.Token != nil {
		if _, err = utils.CheckToken(*req.Token); err != nil {
			resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
			return resp, nil
		}
	}

	count, err := service.NewInteractionService(ctx).CountComments(req)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.CommentCount = count
	return
}
