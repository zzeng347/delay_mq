package service

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Ticker struct {
	
}

var (
	tickerDefaultDuration = 1 * time.Second
)

// InitTicker 根据bucket数量启动相应ticker进行扫描
func InitTicker(ctx context.Context, bucketName string, s *Service)  {
	ticker := time.NewTicker(tickerDefaultDuration)
	defer func() {
		s.wg.Done()
		ticker.Stop()
	}()
	s.wg.Add(1)

	for {
		select {
		case <-ticker.C:
			tickHandler(bucketName)
		case <-ctx.Done(): // 等待上级通知
			log.Printf("ticker Done msg: %#v\n", ctx.Err())
			return
		}
	}
}

func tickHandler(bucketName string)  {
	// 扫描bucket
	bItem, err := s.GetLatestJobFromBucket(bucketName)
	if err != nil {
		return
	}

	if bItem == nil {
		return
	}

	if bItem.Timestamp > time.Now().Unix() {
		fmt.Printf("next job, bucket#%s, job id#%s, time#%s\n", bucketName, bItem.JobId, time.Unix(bItem.Timestamp, 0).Format("2006-01-02 15:04:05"))
		return
	}

	// 读取job
	jobInfo, err := s.GetJob(bItem.JobId)
	if err != nil {
		return
	}
	if jobInfo == nil {
		// 从bucket删除
		err = s.RemoveBucketJob(bucketName, bItem.JobId)
		if err != nil {
			//
		}
		return
	}
	//fmt.Printf("正在消费job#%v\n", jobInfo)

	// 进queue
	queueKey := s.GetQueueKey(jobInfo.Container)
	err = s.PushToQueue(queueKey, bItem.JobId)
	if err != nil {
		return
	}

	// 出bucket
	err = s.RemoveBucketJob(bucketName, bItem.JobId)
	if err != nil {
		return
	}
	return
}