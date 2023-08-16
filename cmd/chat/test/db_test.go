package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ozline/tiktok/cmd/chat/dal"
	"github.com/ozline/tiktok/cmd/chat/dal/db"
)

func TestGetMessageList(t *testing.T) {
	dal.Init()
	list, err := db.GetMessageList(context.Background(), 2, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("num:", len(list))
	for i := 0; i < len(list); i++ {
		fmt.Println(*list[i])
	}
}
