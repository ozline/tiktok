package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

const (
	BlockSize = 1024 * 1024 // 1 MB，每块的大小
)

var blockCount int64
var recvCount int64

func TestPutVideo(t *testing.T) {
	inputFile, err := os.Open("test_video.mp4")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	var videoData []byte
	buffer := make([]byte, 1024) // 使用一个缓冲区来逐步读取文件内容
	for {
		n, err := inputFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("读取文件失败：", err)
		}
		videoData = append(videoData, buffer[:n]...)
	}
	resp, err := conn.PutVideo(context.Background(), &video.PutVideoRequest{
		Title:     "test_title",
		Token:     token,
		VideoFile: videoData,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Printf("Resp:\n%+v\n", resp)
}
