// Code generated by hertz generator.

package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	user "github.com/ozline/tiktok/api-gateway/biz/model/api/user"
)

// CreateUserResponse .
// @router /v1/user/create [POST]
func CreateUserResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CreateUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.CreateUserResp)

	c.JSON(consts.StatusOK, resp)
}

// QueryUserResponse .
// @router /v1/user/query [POST]
func QueryUserResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.QueryUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.QueryUserResp)

	c.JSON(consts.StatusOK, resp)
}

// UpdateUserResponse .
// @router /v1/user/update/:user_id [POST]
func UpdateUserResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UpdateUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UpdateUserResp)

	c.JSON(consts.StatusOK, resp)
}

// DeleteUserResponse .
// @router /v1/user/delete/:user_id [POST]
func DeleteUserResponse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.DeleteUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.DeleteUserResp)

	c.JSON(consts.StatusOK, resp)
}
