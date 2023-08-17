package test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/chat/dal"
	"github.com/ozline/tiktok/cmd/chat/dal/cache"
	"github.com/ozline/tiktok/cmd/chat/dal/db"
	redis "github.com/redis/go-redis/v9"
)

func TestGetMessageList(t *testing.T) {
	dal.Init()
	list, err := db.GetMessageList(context.Background(), 2, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range list {
		res, _ := json.Marshal(v)
		err := cache.RedisDB.ZAdd(context.Background(), "2-3", redis.Z{
			Member: res,
			Score:  float64(v.CreatedAt.Unix()),
		}).Err()
		if err != nil {
			klog.Info(err)

		}
	}
	listTwo, err := db.GetMessageList(context.Background(), 2, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("num:", len(listTwo))
	for i := 0; i < len(listTwo); i++ {
		fmt.Println(*listTwo[i])
	}
}
