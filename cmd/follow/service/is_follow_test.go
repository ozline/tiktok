package service_test

import (
	"context"
	"testing"

	"github.com/ozline/tiktok/cmd/follow/dal"
	"github.com/ozline/tiktok/cmd/follow/service"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/pkg/utils"
)

var isFollowTests = []Test{
	{10001, 10002, "", 1},
	{11001, 10002, "", 1},
}

func TestIsFollow(t *testing.T) {
	config.InitForTest()
	dal.Init()
	followService := service.NewFollowService(context.Background())
	for _, test := range isFollowTests {
		test.token, _ = utils.CreateToken(test.id)
		_, err := followService.IsFollow(&follow.IsFollowRequest{
			UserId:   test.id,
			ToUserId: test.touserid,
			Token:    test.token,
		})

		if err != nil {
			continue
		}
	}
}
