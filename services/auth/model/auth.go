package model

import "github.com/golang-jwt/jwt"

type Auth struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// omitempty : 如果信息未被提供，则在序列化时忽略
