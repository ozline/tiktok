package service

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/kitex_gen/video"
)

// 创建命名管道
func mkfifo(path string) error {
	err := syscall.Mkfifo(path, 0666)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func (s *VideoService) UploadCover(req *video.PutVideoRequest, coverName string) (err error) {
	var imageBuffer bytes.Buffer // 接收图像输出

	// 在临时目录中创建随机命名管道
	pipePath := filepath.Join(os.TempDir(), fmt.Sprintf("input_pipe_%d", time.Now().UnixNano()))

	err = mkfifo(pipePath)

	if err != nil {
		klog.Errorf("error creating named pipe: %v\n", err)
		return err
	}

	defer os.Remove(pipePath)

	// 调用ffmpeg命令
	// cmd := exec.Command("ffmpeg", "-i", pipePath, "-vframes", "1", "tmp_cover.jpg")
	cmd := exec.Command("ffmpeg", "-i", pipePath, "-vframes", "1", "-f", "image2pipe", "-vcodec", "png", "-")
	cmd.Stdout = &imageBuffer
	cmd.Stderr = os.Stderr

	// 给命名管道写入数据
	go func() {
		pipeWriter, err := os.OpenFile(pipePath, os.O_WRONLY, os.ModeNamedPipe)
		if err != nil {
			klog.Errorf("error opening pipe: %v", err)
			return
		}
		defer pipeWriter.Close()

		_, err = pipeWriter.Write(req.VideoFile)

		if err != nil {
			klog.Errorf("error writing to pipe: %v", err)
			return
		}
	}()

	err = cmd.Run()

	if err != nil {
		klog.Errorf("error running FFmpeg: %v", err)
		return err
	}

	// 创建ioReader对象, 上传文件
	imageReader := bytes.NewReader(imageBuffer.Bytes())
	err = s.bucket.PutObject(config.OSS.MainDirectory+"/"+coverName, imageReader)

	if err != nil {
		klog.Errorf("error uploading cover: %v", err)
	}

	return err
}
