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

func (s *Service) GetJob(jobId string) (j *model.JobResp, err error) {
	poolKey := s.GetJobPoolKey(jobId)
	j, err = s.dao.GetJob(poolKey)
	return
}

func (s *Service) SetJob(j *model.JobResp) (err error) {
	poolKey := s.GetJobPoolKey(j.Id)
	err = s.dao.SetJob(poolKey, j)
	return
}

func (s *Service) DelJob(jobId string) (err error) {
	// 删除job pool
	poolKey := s.GetJobPoolKey(jobId)

	jobResp, err := s.dao.GetJob(poolKey)
	if err != nil {
		return
	}

	err = s.dao.DelJob(poolKey)
	if err != nil {
		// TODO 删除失败
	}

	if jobResp.TTR > 0 {
		// 删除ttr bucket
		ttrBucketName := s.GetTtrBucket(jobResp.Id)
		err = s.RemoveBucketJob(ttrBucketName, jobResp.Id)
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

	jobResp := &model.JobResp{
		PushJobReq: *job,
		RetryCount: 0,
	}

	// hash job_id get bucket name
	bucketName := s.GetBucket(jobResp.Id)
	fmt.Printf("push to bucket name#%s\n", bucketName)

	// set job pool
	poolKey := s.GetJobPoolKey(jobResp.Id)
	err = s.dao.SetJob(poolKey, jobResp)
	if err != nil {
		return err
	}

	// push bucket
	timestamp := jobResp.Delay + time.Now().Unix()
	err = s.PushToBucket(bucketName, timestamp, jobResp.Id)
	return err
}