package service

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/kitex_gen/video"
)

func (s *VideoService) UploadCover(req *video.PutVideoRequest, coverName string) (err error) {
	// 将接收到的数据保存为临时文件
	tmpFile, err := os.CreateTemp("", "tmp_video_*.mp4")
	if err != nil {
		return err
	}
	_, err = tmpFile.Write(req.VideoFile)
	if err != nil {
		return err
	}
	// 使用FFmpeg截取第一帧
	ffmpegCmd := exec.Command("ffmpeg", "-i", tmpFile.Name(), "-vframes", "1", "tmp_cover.jpg")
	// 设置标准输入，在发现重名文件时自动回答 y
	ffmpegCmd.Stdin = bytes.NewBufferString("y\n")
	err = ffmpegCmd.Run()
	if err != nil {
		return err
	}
	// 上传封面
	err = s.bucket.PutObjectFromFile(config.OSS.MainDirectory+"/"+coverName, "tmp_cover.jpg")
	if err != nil {
		return err
	}
	// 关闭临时文件并删除
	err = tmpFile.Close()
	if err != nil {
		return err
	}
	err = os.Remove(tmpFile.Name())
	if err != nil {
		return err
	}
	err = os.Remove("tmp_cover.jpg")
	if err != nil {
		return err
	}
	return nil
}
