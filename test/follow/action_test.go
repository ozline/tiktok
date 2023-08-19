package main

import (
	"testing"
	"time"

	"github.com/ozline/tiktok/kitex_gen/follow"
)

func testAction(t *testing.T) {
	err := followService.Action(&follow.ActionRequest{
		Token:      token,
		ToUserId:   touserid,
		ActionType: actiontype,
	})

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}

func BenchmarkAction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := followService.Action(&follow.ActionRequest{
			Token:      token,
			ToUserId:   touserid,
			ActionType: actiontype,
		})

		if err != nil {
			b.Errorf("err: [%v] \n", err)
		}

		time.Sleep(100 * time.Millisecond) // Add a sleep to simulate some processing time
	}
}
