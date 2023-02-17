package main

import (
	follow "github.com/ozline/tiktok/services/follow/kitex_gen/tiktok/follow/tiktokfollowservice"
	"log"
)

func main() {
	svr := follow.NewServer(new(TiktokFollowServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
