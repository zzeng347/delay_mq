package service

import (
	"context"
	"fmt"
)

const (
	BucketKey    = "dmq:bucket:"
	TtrBucketKey = "dmq:ttr_bucket:"
	BucketNum    = 5
	TtrBucketNum = 5
)

type Bucket struct {

}

func InitBucket(ctx context.Context, s *Service)  {
	var bucketName string
	for i := 0; i < BucketNum; i++ {
		bucketName = GetBucketName(BucketKey, i+1)
		// Init ticker
		//go InitTicker(ctx, bucketName, s)
	}

	for i := 0; i < TtrBucketNum; i++ {
		bucketName = GetBucketName(TtrBucketKey, i+1)
		// Init ticker
		//go InitTicker(ctx, bucketName, s)
	}
	fmt.Println(bucketName)
}

func GetBucketName(key string, num int) string {
	return fmt.Sprintf("%s%d", key, num)
}

func (s *Service) GetBucket(jobId string) (bucketName string) {
	hashId := FnvHash32(jobId)
	modulo := hashId%BucketNum
	if modulo == 0 {
		modulo = BucketNum
	}
	bucketName = GetBucketName(BucketKey, int(modulo))
	return
}

func (s *Service) PushToBucket(bucket string, timestamp int64, jobId string) (err error) {
	return s.dao.PushBucket(bucket, float64(timestamp), jobId)
}