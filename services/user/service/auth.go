package service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ozline/tiktok/kitex_gen/tiktok/user"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils/snowflake"
	"github.com/ozline/tiktok/services/user/model"
)

type AuthService struct {
	ctx context.Context
	s   *snowflake.Snowflake
}

// NewAuthService returns a new AuthService.
func NewAuthService(ctx context.Context) *AuthService {
	sf, _ := snowflake.NewSnowflake(constants.SnowflakeDatacenterID, constants.SnowflakeWorkerID)
	return &AuthService{
		ctx: ctx,
		s:   sf,
	}
}

func (a *AuthService) GetToken(req *user.GetTokenRequest) (resp string, err error) {
	fmt.Println(a.s.NextVal())
	claims := model.Auth{
		UserId:   req.UserId,
		Username: req.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()),
			ExpiresAt: int64(time.Now().Unix() + 3600),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	resp, err = token.SignedString([]byte(constants.JwtSecret))

	if err != nil {
		return "", err
	}

	return
}

func (a *AuthService) CheckToken(req *user.CheckTokenRequest) (resp *model.Auth, err error) {
	token, err := jwt.ParseWithClaims(req.Token, &model.Auth{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if resp, ok := token.Claims.(*model.Auth); ok && token.Valid {
		return resp, nil
	}

	return nil, err
}
