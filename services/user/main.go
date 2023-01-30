package main

import (
	kitex_gen "github.com/ozline/tiktok/services/user/kitex_gen/kitex_gen/kitexprotobuf"
	"log"
)

func main() {
	svr := kitex_gen.NewServer(new(KitexProtoBufImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
