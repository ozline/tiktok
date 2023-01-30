package main

import (
	user "github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user/tiktokuserservice"
	"log"
)

func main() {
	svr := user.NewServer(new(TiktokUserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
