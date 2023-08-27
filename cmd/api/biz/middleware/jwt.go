package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/ozline/tiktok/cmd/api/biz/pack"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

func AuthToken() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token, ok := c.GetQuery("token")

		if !ok {
			token, ok = c.GetPostForm("token")
		}

		if !ok {
			hlog.CtxInfof(ctx, "token is not exist clientIP: %v\n", c.ClientIP())
			pack.SendFailResponse(c, errno.AuthorizationFailedError)
			c.Abort()
			return
		}
		hlog.CtxInfof(ctx, "token: %v clientIP: %v\n", token, c.ClientIP())

		claims, err := utils.CheckToken(token)

		if err != nil {
			hlog.CtxInfof(ctx, "token is invalid clientIP: %v\n", c.ClientIP())
			pack.SendFailResponse(c, errno.AuthorizationFailedError)
			c.Abort()
			return
		}

		if claims.UserId < constants.StartID {
			hlog.CtxInfof(ctx, "userid invaild")
			pack.SendFailResponse(c, errno.AuthorizationFailedError)
			c.Abort()
		}
		// c.Set("current_user_id", claims.UserId)

		c.Next(ctx)
	}
}
