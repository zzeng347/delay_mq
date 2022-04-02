package service

import (
	"context"
	"fmt"
)

const (
	BUCKET_KEY = "dmq:bucket:%d"
	TTR_BUCKET_KEY = "dmq:ttr_bucket:%d"
	BUCKET_NUM = 5
	TTR_BUCKET_NUM = 5
)

type Bucket struct {

}

func InitBucket(ctx context.Context, s *Service)  {
	var bucketName string
	for i := 0; i < BUCKET_NUM; i++ {
		bucketName = fmt.Sprintf(BUCKET_KEY, i+1)
		// Init ticker
		go InitTicker(ctx, bucketName, s)
	}
}