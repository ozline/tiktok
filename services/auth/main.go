package main

import (
	auth "github.com/ozline/tiktok/services/auth/kitex_gen/tiktok/auth/tiktokauthservice"
	"log"
)

func main() {
	svr := auth.NewServer(new(TiktokAuthServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
