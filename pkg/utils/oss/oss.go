package oss

import (
	"bytes"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/ozline/tiktok/pkg/constants"
)

type OSS struct {
	bucket        *oss.Bucket
	mainDirectory string
	endpoint      string
}

func NewOSS(endpoint, accessKeyID, accessKeySecret, MainDirectory string) (*OSS, error) {
	res := &OSS{}

	client, err := oss.New(endpoint, accessKeyID, accessKeySecret, oss.UseCname(true))

	if err != nil {
		return nil, err
	}

	res.bucket, err = client.Bucket(constants.OSSBucketName)
	res.mainDirectory = MainDirectory
	res.endpoint = endpoint

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Upload files to OSS
func (o *OSS) UploadObject(filename string, data []byte) (string, error) {
	err := o.bucket.PutObject(o.mainDirectory+"/"+filename, bytes.NewReader(data), oss.Routines(constants.UplaodRoutines))
	if err != nil {
		return "", err
	}

	return o.BuildSourceURL(filename), nil
}

func (o *OSS) BuildSourceURL(filename string) string {
	return "https://" + o.endpoint + "/" + o.mainDirectory + "/" + filename
}

// func DeleteObject(bucket *oss.Bucket, deletePath string) error {
// 	return nil
// }
