package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/pkg/errno"
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
			n, err := inputFile.Read(buffer)
			if err == io.EOF {
				putStream.Send(&video.PutVideoRequest{
					IsFinished: true,
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
			}
			fmt.Printf("正在发送block %v\n", blockCount)
			err = putStream.Send(req)
			if err != nil {
				panic(err)
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
		fmt.Printf("Resp:\n %v\n", resp)

		if resp.State == 2 {
			fmt.Println("success")
			break
		}
	}

}
