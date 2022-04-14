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

func (s *Service) PushJob(job *model.PushJobReq) error {
	// 验证job_id唯一性
	poolKey := s.GetJobPoolKey(job.Id)
	jobInfo, err := s.dao.GetJob(poolKey)
	if err == redis.Nil {

	} else if err != nil {
		return err
	} else if jobInfo != nil {
		return errors.New("job id exist")
	}

	// hash job_id get bucket name
	bucketName := s.GetBucket(job.Id)
	fmt.Printf("push to ttr bucket name#%s\n", bucketName)

	// set job pool
	err = s.dao.SetJob(poolKey, job)
	if err != nil {
		return err
	}

	// push bucket
	timestamp := job.Delay + time.Now().Unix()
	err = s.PushToBucket(bucketName, timestamp, job.Id)
	return err
}