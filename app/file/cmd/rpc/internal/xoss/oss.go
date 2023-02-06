package xoss

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

type Xoss struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	PlayPath        string
}

func NewOssClient(xoss Xoss) *oss.Bucket {
	client, _ := oss.New(xoss.Endpoint, xoss.AccessKeyId, xoss.AccessKeySecret)
	bucket, _ := client.Bucket(xoss.BucketName)
	return bucket
}
