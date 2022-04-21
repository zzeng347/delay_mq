package service

import (
	"context"
	"delay_mq_v2/model"
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
		go InitTicker(ctx, bucketName, s)
	}

	for i := 0; i < TtrBucketNum; i++ {
		bucketName = GetBucketName(TtrBucketKey, i+1)
		// Init ticker
		go InitTicker(ctx, bucketName, s)
	}
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

func (s *Service) GetTtrBucket(jobId string) (bucketName string) {
	hashId := FnvHash32(jobId)
	modulo := hashId%TtrBucketNum
	if modulo == 0 {
		modulo = TtrBucketNum
	}
	bucketName = GetBucketName(TtrBucketKey, int(modulo))
	return
}

func (s *Service) PushToBucket(bucket string, timestamp int64, jobId string) (err error) {
	return s.dao.PushBucket(bucket, float64(timestamp), jobId)
}

// GetLatestJobFromBucket TODO 批量获取可执行job
func (s *Service) GetLatestJobFromBucket(bucket string) (bItem *model.ZRangeBucketItem, err error) {
	bItem = &model.ZRangeBucketItem{}

	zRet, err := s.dao.ZRangeBucket(bucket, 0, 1)
	if err != nil {
		return
	}

	if len(zRet) < 1 {
		return nil, nil
	}

	for _, z := range zRet {
		bItem.JobId = z.Member.(string)
		bItem.Timestamp = int64(z.Score)
	}
	return
}

func (s *Service) RemoveBucketJob(bucket string, jobId string) (err error) {
	return s.dao.RemoveInBucket(bucket, jobId)
}