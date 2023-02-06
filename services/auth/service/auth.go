package service

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/services/auth/kitex_gen/tiktok/auth"
	"github.com/ozline/tiktok/services/auth/model"
)

var (
	secret = constants.JwtSecret
)

type AuthService struct {
	ctx context.Context
}

// NewAuthService returns a new AuthService.
func NewAuthService(ctx context.Context) *AuthService {
	return &AuthService{ctx: ctx}
}

func (a *AuthService) GetToken(req *auth.GetTokenRequest) (resp string, err error) {
	claims := model.Auth{
		UserId:   req.UserId,
		Username: req.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()),
			ExpiresAt: int64(time.Now().Unix() + 3600),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	resp, err = token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return
}

func (a *AuthService) CheckToken(req *auth.CheckTokenRequest) (resp *model.Auth, err error) {
	token, err := jwt.ParseWithClaims(req.Token, &model.Auth{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if resp, ok := token.Claims.(*model.Auth); ok && token.Valid {
		return resp, nil
	}

	return nil, err
}
