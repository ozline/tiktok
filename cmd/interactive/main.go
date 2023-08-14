package main

import (
	interactive "github.com/ozline/tiktok/kitex_gen/interactive/interactiveservice"
	"log"
)

func main() {
	svr := interactive.NewServer(new(InteractiveServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
