package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

type TokenService struct {
	rateLimit *rate.Limiter
}

func NewTokenService(ctx context.Context) *TokenService {
	return &TokenService{
		rateLimit: &rate.Limiter{},
	}
}
func (t *TokenService) tokenBucket() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := t.rateLimit.WaitN(c, 10)
		if err != nil {
			c.String(http.StatusOK, "rate limit,Drop")
			c.Abort()
			return
		}
		c.Next()
	}
}

func (t *TokenService) TokenLimits() {
	//limit := rate.Every(100 * time.Millisecond)
	t.rateLimit = rate.NewLimiter(10, 100)
	r := gin.Default()
	r.GET("/ping", t.tokenBucket(), func(c *gin.Context) {
		c.JSON(200, true)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
