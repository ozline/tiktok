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

var followCountTests = []Test{
	{10001, 10002, "", 1},
	{001, 002, "", 1},
}

func TestFollowCount(t *testing.T) {
	config.InitForTest()
	dal.Init()
	followService := service.NewFollowService(context.Background())
	for _, test := range followCountTests {
		test.token, _ = utils.CreateToken(test.id)
		_, err := followService.FollowCount(&follow.FollowCountRequest{
			UserId: test.id,
			Token:  test.token,
		})

		if err != nil {
			continue
		}
	}
}
