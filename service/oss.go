package service

import (
	"backend-svc-go/global"
	"bytes"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func PushObject(fileContent []byte, filename string) (string, error) {
	bucketName := "littleadds"
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	endpoint := "oss-cn-heyuan.aliyuncs.com"

	// 检查环境变量是否已经设置。
	if endpoint == "" || bucketName == "" {
		log.Fatal("Please set yourEndpoint and bucketName.")
	}

	client, err := oss.New(endpoint, os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		// HandleError(err)
		handleError(err)
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		// HandleError(err)
		handleError(err)
	}
	ext := filepath.Ext(filename)
	filename = global.RandomString(12) +"." + ext
	err = bucket.PutObject(filename, bytes.NewReader(fileContent))
	if err != nil {
		// HandleError(err)
		handleError(err)
	}

	fileURL := fmt.Sprintf("https://%s.%s/%s", bucketName, endpoint, url.PathEscape(filename))
	return fileURL, err
}

func handleError(err error) {
	log.Fatalf("Error: %v", err)
}

// listObjects 用于列举OSS存储空间中的所有对象。
// 参数：
//
//	bucketName - 存储空间名称。
//	endpoint - Bucket对应的Endpoint。
//
// 如果成功，打印所有对象；否则，返回错误。
func listObjects(client oss.Client, bucketName string) error {
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 列举文件。
	marker := ""
	for {
		lsRes, err := bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			return err
		}

		// 打印列举文件，默认情况下一次返回100条记录。
		for _, object := range lsRes.Objects {
			log.Printf("Object: %s", object.Key)
		}

		if !lsRes.IsTruncated {
			break
		}
		marker = lsRes.NextMarker
	}

	return nil
}
