package pack

import (
	"github.com/ozline/tiktok/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/ozline/tiktok/cmd/api/biz/model/api"
	"github.com/ozline/tiktok/pkg/errno"
)

type Response struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

func SendFailResponse(c *app.RequestContext, err error) {
	c.JSON(consts.StatusOK, Response{
		Code: -1,
		Msg:  errno.ConvertErr(err).Error(),
	})
}

func SendResponse(c *app.RequestContext, data interface{}) {
	c.JSON(consts.StatusOK, data)
}

func User(data *user.User) *api.User {
	return &api.User{
		ID:              data.Id,
		Name:            data.Name,
		FollowCount:     &data.FollowCount,
		FollowerCount:   &data.FollowerCount,
		IsFollow:        data.IsFollow,
		Avatar:          &data.Avatar,
		BackgroundImage: &data.BackgroundImage,
		Signature:       &data.Signature,
		TotalFavorited:  &data.TotalFavorited,
		WorkCount:       &data.WorkCount,
		FavoriteCount:   &data.FavoritedCount,
	}
}
