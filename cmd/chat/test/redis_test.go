package test


import (
    "github.com/ozline/tiktok/cmd/chat/dal/db"
    "github.com/ozline/tiktok/cmd/chat/dal"
    "context"
    "testing"
    "fmt"
    "time"
    "github.com/go-redis/redis/v8"
    "strconv"
)


func TestRedis(t *testing.T){
    dal.Init()
    res,_ := db.RedisDB.Exists(context.Background(), "key1").Result()
    fmt.Println(res)
    mem,err :=db.RedisDB.ZRevRangeByScore(context.Background(),"key1",&redis.ZRangeBy{
        Min:strconv.Itoa(0),
        Max:strconv.Itoa(int(time.Now().Unix())),
	}).Result()
    if err!=nil{
        panic(err)
    }
    fmt.Println(mem)
}