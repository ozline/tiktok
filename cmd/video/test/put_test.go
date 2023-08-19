package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/pkg/errno"
)

const (
	BlockSize = 1024 * 1024 // 1 MB，每块的大小
)

var blockCount int64
var recvCount int64

func TestPutVideo(t *testing.T) {
	text_video, err := os.Open("test_video.mp4")
	if err != nil {
		panic(err)
	}
	defer text_video.Close()
	text_cover, err := os.ReadFile("test_cover.jpg")
	if err != nil {
		panic(err)
	}
	buffer := make([]byte, BlockSize)
	blockCount = 0
	putStream, err := conn.PutVideo(context.Background())
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	go func() {
		for {
			//读取文件
			n, err := text_video.Read(buffer)
			if err == io.EOF {
				putStream.Send(&video.PutVideoRequest{
					IsFinished: true,
					Cover:      text_cover,
					Token:      token,
					Title:      fmt.Sprintf("%v test_video", time.Now().Unix()),
				})
				break
			}
			if err != nil {
				panic(err)
			}
			//开始发送请求
			req := &video.PutVideoRequest{
				VideoBlock: buffer[:n],
				UserId:     10000,
				BlockId:    blockCount,
				IsFinished: false,
				Token:      token,
			}
			fmt.Printf("正在发送block %v\n", blockCount)
			err = putStream.Send(req)
			if err != nil {
				t.Error(err)
				t.Fail()
			}
			blockCount++
		}
	}()
	recvCount = 0
	for {
		resp, err := putStream.Recv()
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		if resp.Base.Code != errno.SuccessCode {
			t.Error(errno.NewErrNo(resp.Base.Code, resp.Base.Msg))
			t.Fail()
		}
		t.Logf("Resp:\n %+v\n", resp)

		if resp.State == 2 {
			fmt.Println("success")
			break
		}
	}

}
