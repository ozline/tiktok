package snapshot

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/ozline/tiktok/pkg/constants"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// GetSnapShot Get a snapshot from a video, and save it to the same path as the video
// videoPath: The path of the video. NOTE: include the video filename
// savename: The name of the snapshot. NOTE: NOT include the suffix
func GetSnapShot(videoPath, savename string) (imgfile string, err error) {
	snapshotPath, _ := filepath.Split(videoPath)

	buf := bytes.NewBuffer(nil)

	err = ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", constants.FrameNum)}).Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).WithOutput(buf, os.Stdout).Run()

	if err != nil {
		return "", err
	}

	img, err := imaging.Decode(buf)

	if err != nil {
		return "", err
	}

	err = imaging.Save(img, snapshotPath+savename+".png")

	if err != nil {
		return "", err
	}

	return snapshotPath + savename + ".png", nil
}
