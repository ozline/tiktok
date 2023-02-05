package middleware

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "ozline"

	loginStruct struct {
		Account  string `form:"account" json:"account" query:"account" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
		Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
	}
)

type JWT struct{}

func (*JWT) Init() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "ozline",
		Key:           []byte("ozline"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		TokenLookup:   "header:Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
	})

	if err != nil {
		panic(err)
	}
}

func (*JWT) authenticator(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	if err := c.BindAndValidate(&loginStruct); err != nil {
		return nil, err
	}

}
