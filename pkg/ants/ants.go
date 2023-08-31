package ants

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ozline/tiktok/cmd/chat/dal/db"
	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/panjf2000/ants"
)

var (
	AntsPool *ants.PoolWithFunc
	Wg       sync.WaitGroup
)

func Init() {
	ants_Pool, err := ants.NewPoolWithFunc(500, func(payload interface{}) {
		defer Wg.Done()
		val := payload.(*db.MessageBuild)
		create_time := fmt.Sprintf("%v", val.MessageElem.CreatedAt.UnixMilli())
		message := &chat.Message{
			Id:         val.MessageElem.Id,
			ToUserId:   val.MessageElem.ToUserId,
			FromUserId: val.MessageElem.FromUserId,
			Content:    val.MessageElem.Content,
			CreateTime: &create_time,
		}
		//将传入的message转为middlemessage
		val.MessageList = append(val.MessageList, message)
	})
	if err != nil {
		panic(errors.New("[ants goroutine init error]"))
	}
	AntsPool = ants_Pool
}
