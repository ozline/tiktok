package main

import (
	"testing"

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
