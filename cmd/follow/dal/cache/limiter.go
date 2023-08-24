package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
)

var (
	rate     = 100             // 每秒钟最多处理的请求数
	interval = 1 * time.Second // 限流器的时间间隔
)

func Limit(ctx context.Context) error {
	key := LimiterKey
	now := time.Now().UnixNano()
	// 移除时间戳小于(now - interval)的令牌数
	RedisClient.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%d", now-int64(interval)))
	// 获取当前集合长度，即令牌数
	size, err := RedisClient.ZCard(ctx, key).Result()
	if err != nil {
		return err
	}
	// 如果令牌数大于等于最大请求数拒绝该请求
	if size >= int64(rate) {
		klog.Info("limit!")
		return errors.New("too many request")
	}
	// 否则添加一个新的时间戳成为成员
	RedisClient.ZAdd(ctx, key, redis.Z{Score: float64(now), Member: float64(now)})
	return nil
}
