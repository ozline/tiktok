package main

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

var accessKey = "m5KRX39z1fu9ssut0SFgCWwLxxRiWHB-I2jPalWV"
var secretKey = "CRmeH-AESMTlOr9bCPpDIVtndztgJe_3CHtdVSoK"
var mac = qbox.NewMac(accessKey, secretKey)

// 上传文件 bucket="titok"
func StoragetPutFile(localFileName string, storageFileName string, bucketName string) error {
	bucket := bucketName
	key := storageFileName
	localFile := localFileName

	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := MyPutRet{}
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
	fmt.Println("----- Return Response -----")
	fmt.Println(ret.Bucket, ret.Key, ret.Fsize, ret.Hash, ret.Name)
	return err
}

// 删除文件
func StorageDeleteFile(fileName string, bucketName string) error {
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Region=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)

	bucket := bucketName
	key := fileName

	err := bucketManager.Delete(bucket, key)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

// 获取文件信息
func StorageGetFileInfo(fileName string, bucketName string) {
	bucket := bucketName
	key := fileName

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	fileInfo, sErr := bucketManager.Stat(bucket, key)
	if sErr != nil {
		fmt.Println(sErr)
		return
	}
	fmt.Println(fileInfo.String())
	//可以解析文件的PutTime上传事件，最后一位使用者，文件的hash值，文件大小
	fmt.Println(storage.ParsePutTime(fileInfo.PutTime))
}

// 移动文件
func StorageMoveFile(srcBucketName string, srcFileName string, destBucketName string, destFileName string) {
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)

	srcBucket := srcBucketName
	srcKey := srcFileName
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
func StorageCopyFile(srcBucketName string, srcFileName string, destBucketName string, destFileName string) {
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: true,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)

	srcBucket := srcBucketName
	srcKey := srcFileName
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
func StorageGetFileList(bucketName string, prefix string) {
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
			fmt.Println("list error,", err)
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
func StorageBatchGetFileInfo() {
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
func StorageBatchDeleteFiles() {
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
func StorageBatchCopyFiles() {
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
