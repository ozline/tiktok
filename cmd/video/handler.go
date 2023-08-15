package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/cmd/video/pack"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...
	return
}

func (s *VideoServiceImpl) PutVideo(stream video.VideoService_PutVideoServer) (err error) {
	resp := new(video.PutVideoResponse)
	mergeVideo, err := os.Create("merge_video.mp4")
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		resp.State = 0
		stream.Send(resp)
		return nil
	}

	for {
		req, err := stream.Recv()
		if err != nil {
			resp.Base = pack.BuildBaseResp(err)
			resp.State = 0
			stream.Send(resp)
			return nil
		}
		if req.IsFinished {
			//视频全部传输完成
			//TODO 上传
			fmt.Println("视频全部传输完成")
			resp.Base = pack.BuildBaseResp(nil)
			resp.State = 2
			stream.Send(resp)
			break

		}
		log.Printf("received block %v:", req.GetBlockId())
		_, err = mergeVideo.Write(req.VideoBlock)
		if err != nil {
			resp.Base = pack.BuildBaseResp(err)
			resp.State = 0
			stream.Send(resp)
			return nil
		}
		resp.Base = pack.BuildBaseResp(nil)
		resp.State = 1
		stream.Send(resp)
	}
	stream.Close()
	mergeVideo.Close()
	return
}
