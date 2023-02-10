package service

import (
	"context"
	"fmt"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils/snowflake"
	"github.com/ozline/tiktok/services/video/model"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"strconv"
)

type StorageService struct {
	ctx context.Context
	s   *snowflake.Snowflake
}

func NewStorageService(ctx context.Context) *StorageService {
	sf, _ := snowflake.NewSnowflake(constants.SnowflakeDatacenterID, constants.SnowflakeWorkerID)
	return &StorageService{
		ctx: ctx,
		s:   sf,
	}
}

var accessKey = "m5KRX39z1fu9ssut0SFgCWwLxxRiWHB-I2jPalWV"
var secretKey = "CRmeH-AESMTlOr9bCPpDIVtndztgJe_3CHtdVSoK"
var mac = qbox.NewMac(accessKey, secretKey)

// 上传文件 bucket="titok"
func (s *StorageService) StoragPutVideo(localFileName string, storageFileName int64, bucketName string) error {
	bucket := bucketName
	key := strconv.FormatInt(storageFileName, 10)
	localFile := localFileName

	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := model.MyPutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}

	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}
func (s *StorageService) StorageDownloadOneVideo(videoID int64, bucketName string) string {
	domain := "rp9zcsyip.hb-bkt.clouddn.com"
	key := strconv.FormatInt(videoID, 10)
	publicAccessURL := storage.MakePublicURL(domain, key)
	return publicAccessURL
}

// 删除文件
func (s *StorageService) StorageDeleteVideo(videoID int64, bucketName string) bool {
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Region=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)

	bucket := bucketName
	key := strconv.FormatInt(videoID, 10)

	err := bucketManager.Delete(bucket, key)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// 获取文件信息 FileHash,FileSize,FileMimeType
func (s *StorageService) StorageGetVideoInfo(videoID int64, bucketName string) (string, int64, string) {
	bucket := bucketName
	key := strconv.FormatInt(videoID, 10)

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	fileInfo, sErr := bucketManager.Stat(bucket, key)
	if sErr != nil {
		fmt.Println(sErr)
		return "-1", 0, "Nothing"
	}
	//可以解析文件的PutTime上传事件，最后一位使用者，文件的hash值，文件大小
	return fileInfo.Hash, fileInfo.Fsize, fileInfo.MimeType
}

// 移动文件
func (s *StorageService) StorageMoveVideo(srcBucketName string, srcVideoID int64, destBucketName string, destFileName string) {
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)

	srcBucket := srcBucketName
	srcKey := strconv.FormatInt(srcVideoID, 10)
	//目标空间可以和源空间相同，但是不能为跨机房的空间
	destBucket := srcBucket
	//目标文件名可以和源文件名相同，也可以不同
	destKey := destFileName
	//如果目标文件存在，是否强制覆盖，如果不覆盖，默认返回614 file exists
	force := false
	err := bucketManager.Move(srcBucket, srcKey, destBucket, destKey, force)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 复制文件
func (s *StorageService) StorageCopyVideo(srcBucketName string, srcVideoID int64, destBucketName string, destFileName string) {
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)

	srcBucket := srcBucketName
	srcKey := strconv.FormatInt(srcVideoID, 10)
	//目标空间可以和源空间相同，但是不能为跨机房的空间
	destBucket := destBucketName
	//目标文件名可以和源文件名相同，也可以不同
	destKey := destFileName
	//如果目标文件存在，是否强制覆盖，如果不覆盖，默认返回614 file exists
	force := false
	err := bucketManager.Copy(srcBucket, srcKey, destBucket, destKey, force)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 获取指定前缀的文件列表
func (s *StorageService) StorageGetVideoList(bucketName string, prefix string) {
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)

	bucket := bucketName
	limit := 1000
	delimiter := ""
	//初始列举marker为空
	marker := ""
	for {
		entries, _, nextMarker, hasNext, err := bucketManager.ListFiles(bucket, prefix, delimiter, marker, limit)
		if err != nil {
			break
		}
		//print entries
		for _, entry := range entries {
			fmt.Println(entry.Key)
		}
		if hasNext {
			marker = nextMarker
		} else {
			//list end
			break
		}
	}
}

// 批量获取文件信息
func (s *StorageService) StorageBatchGetVideoInfo() {
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)

	//每个batch的操作数量不可以超过1000个，如果总数量超过1000，需要分批发送
	bucket := "titok"
	keys := []string{
		"github1.png",
		"github2.png",
		"github3.png",
		"github4.png",
		"github5.png",
	}
	statOps := make([]string, 0, len(keys))
	for _, key := range keys {
		statOps = append(statOps, storage.URIStat(bucket, key))
	}

	rets, err := bucketManager.Batch(statOps)
	if err != nil {
		// 遇到错误
		if _, ok := err.(*storage.ErrorInfo); ok {
			for _, ret := range rets {
				// 200 为成功
				fmt.Printf("%d\n", ret.Code)
				if ret.Code != 200 {
					fmt.Printf("%s\n", ret.Data.Error)
				} else {
					fmt.Printf("%v\n", ret.Data)
				}
			}
		} else {
			fmt.Printf("batch error, %s", err)
		}
	} else {
		// 完全成功
		for _, ret := range rets {
			// 200 为成功
			fmt.Printf("%d\n", ret.Code)
			fmt.Printf("%v\n", ret.Data)
		}
	}
}

// 批量删除文件
func (s *StorageService) StorageBatchDeleteVideos() {
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	//每个batch的操作数量不可以超过1000个，如果总数量超过1000，需要分批发送
	bucket := "if-pbl"
	keys := []string{
		"github1.png",
		"github2.png",
		"github3.png",
		"github4.png",
		"github5.png",
	}
	deleteOps := make([]string, 0, len(keys))
	for _, key := range keys {
		deleteOps = append(deleteOps, storage.URIDelete(bucket, key))
	}

	rets, err := bucketManager.Batch(deleteOps)
	if err != nil {
		// 遇到错误
		if _, ok := err.(*storage.ErrorInfo); ok {
			for _, ret := range rets {
				// 200 为成功
				fmt.Printf("%d\n", ret.Code)
				if ret.Code != 200 {
					fmt.Printf("%s\n", ret.Data.Error)
				}
			}
		} else {
			fmt.Printf("batch error, %s", err)
		}
	} else {
		// 完全成功
		for _, ret := range rets {
			// 200 为成功
			fmt.Printf("%d\n", ret.Code)
		}
	}
}

// 批量复制文件
func (s *StorageService) StorageBatchCopyVideos() {
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)

	srcBucket := "if-pbl"
	destBucket := "if-pbl"
	force := true
	copyKeys := map[string]string{
		"github1.png": "github1-copy.png",
		"github2.png": "github2-copy.png",
		"github3.png": "github3-copy.png",
		"github4.png": "github4-copy.png",
		"github5.png": "github5-copy.png",
	}
	copyOps := make([]string, 0, len(copyKeys))
	for srcKey, destKey := range copyKeys {
		copyOps = append(copyOps, storage.URICopy(srcBucket, srcKey, destBucket, destKey, force))
	}

	rets, err := bucketManager.Batch(copyOps)
	if err != nil {
		// 遇到错误
		if _, ok := err.(*storage.ErrorInfo); ok {
			for _, ret := range rets {
				// 200 为成功
				fmt.Printf("%d\n", ret.Code)
				if ret.Code != 200 {
					fmt.Printf("%s\n", ret.Data.Error)
				}
			}
		} else {
			fmt.Printf("batch error, %s", err)
		}
	} else {
		// 完全成功
		for _, ret := range rets {
			// 200 为成功
			fmt.Printf("%d\n", ret.Code)
			fmt.Printf("%v\n", ret.Data)
		}
	}
}

func (s *StorageService) GetNUrlByVideoID(videos []model.VideoStorageInfo) []string {
	number := len(videos)
	urls := make([]string, number)
	for index, video := range videos {
		urls[index] = s.StorageDownloadOneVideo(video.VideoID, "titok")
	}

	return urls
}
