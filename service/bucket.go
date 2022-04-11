package service

import (
	"context"
	"fmt"
)

const (
	BucketKey    = "dmq:bucket:%d"
	TtrBucketKey = "dmq:ttr_bucket:%d"
	BucketNum    = 5
	TtrBucketNum = 5
)

type Bucket struct {

}

func InitBucket(ctx context.Context, s *Service)  {
	var bucketName string
	for i := 0; i < BucketNum; i++ {
		bucketName = fmt.Sprintf(BucketKey, i+1)
		// Init ticker
		go InitTicker(ctx, bucketName, s)
	}

	for i := 0; i < TtrBucketNum; i++ {
		bucketName = fmt.Sprintf(TtrBucketKey, i+1)
		// Init ticker
		go InitTicker(ctx, bucketName, s)
	}
}