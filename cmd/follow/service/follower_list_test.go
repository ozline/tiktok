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

var followerListTests = []Test{
	{10001, 10002, "", 1},
	{11001, 10002, "", 1},
}

func TestFollowerList(t *testing.T) {
	config.InitForTest()
	dal.Init()
	rpc.Init()
	followService := service.NewFollowService(context.Background())
	for i, test := range followerListTests {
		test.token, _ = utils.CreateToken(test.id)
		_, err := followService.FollowerList(&follow.FollowerListRequest{
			UserId: test.id,
			Token:  test.token,
		})

		if err != nil {
			klog.Infof("test num %v,err:%v", i, err)
			continue
		}
	}
}
