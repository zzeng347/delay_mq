package service

import (
	"fmt"
)

const (
	QueuePrefix = "dmq:queue:"
)

var (
	QueueContainer = map[string]string{"container1":"container1", "container2":"container2"}
)

func (s *Service) GetQueueKey(container string) string {
	return fmt.Sprintf("%s%s", QueuePrefix, container)
}

func (s *Service) PushToQueue(queue string, jobId string) (err error) {
	err = s.dao.PushQueue(queue, jobId)
	return
}

func (s *Service) PopFromQueue(queue string) (jobIds []string, err error) {
	jobIds, err = s.dao.PopQueue(queue)
	return
}