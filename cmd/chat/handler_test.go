package main

import (
	"context"
	"testing"
	"time"

	"github.com/ozline/tiktok/cmd/chat/dal"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/pkg/utils"
)

func TestHadnlerGet(t *testing.T) {
	t.Log("result===>")
	config.InitForTest()
	dal.Init()
	msi := new(MessageServiceImpl)
	token, err := utils.CreateToken(3)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	mlr, err := msi.MessageList(context.Background(), &chat.MessageListRequest{
		Token:    token,
		ToUserId: 2,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log("result===>", mlr.MessageList)
	t.Log("result===>", mlr.Total)
	time.Sleep(2 * time.Second)
}
