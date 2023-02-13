package main

import (
	"fmt"
	"os"
)

func main() {
	fi, err := os.Stat("video1.mp4")
	if err == nil {
		fmt.Println("name:", fi.Name())
		fmt.Println("size:", fi.Size())
		fmt.Println("is dir:", fi.IsDir())
		fmt.Println("mode::", fi.Mode())
		fmt.Println("modTime:", fi.ModTime())
	}
}
