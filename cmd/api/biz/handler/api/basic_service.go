// Code generated by hertz generator.

package api

import (
	"context"
	"io"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/h2non/filetype"
	api "github.com/ozline/tiktok/cmd/api/biz/model/api"
	"github.com/ozline/tiktok/cmd/api/biz/pack"
	"github.com/ozline/tiktok/cmd/api/biz/rpc"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/kitex_gen/video"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

// Feed .
// @router /douyin/feed [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(api.FeedResponse)

	videoList, nexttime, err := rpc.VideoFeed(ctx, &video.FeedRequest{
		LatestTime: req.LatestTime,
		Token:      req.Token,
	})

	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.VideoList = pack.VideoList(videoList)
	resp.NextTime = &nexttime
	pack.SendResponse(c, resp)
}

// UserRegister .
// @router /douyin/user/register [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(api.UserRegisterResponse)

	resp.UserID, resp.Token, err = rpc.UserRegister(ctx, &user.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	// resp.StatusCode = errno.StatusSuccessCode
	pack.SendResponse(c, resp)
}

// UserLogin .
// @router /douyin/user/login [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(api.UserLoginResponse)

	resp.UserID, resp.Token, err = rpc.UserLogin(ctx, &user.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	pack.SendResponse(c, resp)
}

// UserInfo .
// @router /douyin/user [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(api.UserResponse)

	user, err := rpc.UserInfo(ctx, &user.InfoRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})

	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.User = pack.User(user)
	pack.SendResponse(c, resp)
}

// PublishAction .
// @router /douyin/publish/action [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(api.PublishActionResponse)

	file, err := c.FormFile("data")

	if err != nil {
		pack.SendFailResponse(c, errno.FileUploadError.WithMessage(err.Error()))
		return
	}

	// check header
	if !utils.IsVideoFile(file) {
		pack.SendFailResponse(c, errno.NotVideoFile)
		return
	}

	fileContent, err := file.Open()

	if err != nil {
		pack.SendFailResponse(c, errno.FileUploadError.WithMessage(err.Error()))
		return
	}

	byteContainer, err := io.ReadAll(fileContent)

	if err != nil {
		pack.SendFailResponse(c, errno.FileUploadError.WithMessage(err.Error()))
		return
	}

	if !filetype.IsVideo(byteContainer) {
		pack.SendFailResponse(c, errno.NotVideoFile)
	}

	err = rpc.VideoPublish(ctx, &video.PutVideoRequest{
		VideoFile: byteContainer,
		Title:     req.Title,
		Token:     req.Token,
	})

	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	pack.SendResponse(c, resp)
}

// PublishList .
// @router /douyin/publish/list [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.PublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(api.PublishListResponse)

	videoList, err := rpc.PublishList(ctx, &video.GetPublishListRequest{
		Token:  req.Token,
		UserId: req.UserID,
	})

	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.VideoList = pack.VideoList(videoList)
	pack.SendResponse(c, resp)
}
