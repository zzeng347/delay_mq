package service

import (
	"delay_mq_v2/model"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type Job struct {

}

const (
	JobPoolKey = "dmq:jobpool:"
)

func (s *Service) GetJobPoolKey(key string) string {
	return fmt.Sprintf("%s%s", JobPoolKey, key)
}

func (s *Service) GetJob(jobId string) (j *model.PushJobReq, err error) {
	poolKey := s.GetJobPoolKey(jobId)
	j, err = s.dao.GetJob(poolKey)
	return
}

func (s *Service) DelJob(jobId string) (err error) {
	// 删除job pool
	poolKey := s.GetJobPoolKey(jobId)

	jobInfo, err := s.dao.GetJob(poolKey)
	if err != nil {
		return
	}

	err = s.dao.DelJob(poolKey)
	if err != nil {
		// TODO 删除失败
	}

	if jobInfo.TTR > 0 {
		// 删除ttr bucket
		ttrBucketName := s.GetTtrBucket(jobInfo.Id)
		err = s.RemoveBucketJob(ttrBucketName, jobInfo.Id)
		if err != nil {
			// TODO 删除失败
		}
	}

	return
}

func (s *Service) PushJob(job *model.PushJobReq) error {
	// 验证container
	if _, ok := QueueContainer[job.Container]; !ok {
		return errors.New("container error")
	}
	
	// 验证job_id唯一性
	jobInfo, err := s.GetJob(job.Id)
	if err == redis.Nil {

	} else if err != nil {
		return err
	} else if jobInfo != nil {
		return errors.New("job id exist")
	}

	// hash job_id get bucket name
	bucketName := s.GetBucket(job.Id)
	fmt.Printf("push to bucket name#%s\n", bucketName)

	// set job pool
	poolKey := s.GetJobPoolKey(job.Id)
	err = s.dao.SetJob(poolKey, job)
	if err != nil {
		return err
	}

	// push bucket
	timestamp := job.Delay + time.Now().Unix()
	err = s.PushToBucket(bucketName, timestamp, job.Id)
	return err
}