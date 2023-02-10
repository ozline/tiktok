package service

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils/snowflake"
	"github.com/ozline/tiktok/services/auth/kitex_gen/tiktok/auth"
	"github.com/ozline/tiktok/services/auth/model"
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

func (a *AuthService) GetToken(req *auth.GetTokenRequest) (resp string, err error) {
	fmt.Println(a.s.NextVal())
	claims := model.Auth{
		UserId:   req.UserId,
		Username: req.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 3600,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	resp, err = token.SignedString([]byte(constants.JwtSecret))

	if err != nil {
		return "", err
	}

	return
}

func (a *AuthService) CheckToken(req *auth.CheckTokenRequest) (resp *model.Auth, err error) {
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
