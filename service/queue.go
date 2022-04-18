package service

import (
	"fmt"
)

const (
	QueueKey = "dmq:queue:"
)

func (s *Service) GetQueueKey(jobId string) string {
	return fmt.Sprintf("%s%s", QueueKey, jobId)
}

func (s *Service) PushToQueue(queue string, jobId string) (err error) {
	err = s.dao.PushQueue(queue, jobId)
	return
}