package service

import (
	"fmt"
)

const (
	QueueKey = "dmq:queue"
)

func (s *Service) GetQueueKey() string {
	return fmt.Sprintf("%s", QueueKey)
}

func (s *Service) PushToQueue(queue string, jobId string) (err error) {
	err = s.dao.PushQueue(queue, jobId)
	return
}

func (s *Service) PopFromQueue(queue string) (jobIds []string, err error) {
	jobIds, err = s.dao.PopQueue(queue)
	return
}