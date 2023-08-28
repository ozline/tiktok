package service_test

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/follow/dal"
	"github.com/ozline/tiktok/cmd/follow/rpc"
	"github.com/ozline/tiktok/cmd/follow/service"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/pkg/utils"
)

type Test struct {
	id         int64
	touserid   int64
	token      string
	actiontype int64
}

var actionTests = []Test{
	{10001, 10002, "", 2},
	{10002, 10001, "", 2},
	{10001, 10002, "", 1},
	{10001, 10002, "", 1},
	{10001, 10002, "", 2},
	{10001, 10002, "", 2},
	{10001, 10002, "", 1},
	{10001, 10002, "hhh", 1},
	{10001, 10002, "", 3},
	{10002, 10086, "", 1},
	{10001, 10001, "", 1},
	{10001, 10002, "", 1},
	{10002, 10001, "", 1},
}

func TestAction(t *testing.T) {
	config.InitForTest()
	dal.Init()
	rpc.Init()
	followService := service.NewFollowService(context.Background())
	for i, test := range actionTests {
		test.token, _ = utils.CreateToken(test.id)
		err := followService.Action(&follow.ActionRequest{
			Token:      test.token,
			ToUserId:   test.touserid,
			ActionType: test.actiontype,
		})

		if err != nil {
			klog.Infof("test num %v,err:%v", i, err)
			continue
		}
	}
}
