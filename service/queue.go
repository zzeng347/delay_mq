package service

import (
	"context"
	"fmt"
	"log"
	"time"
)

type QueueConfig struct {
	ExecerUrl string
}

const (
	QueuePrefix = "dmq:queue:"
)

var (
	QueueContainer = map[string]string{"container1":"container1", "container2":"container2"}
	queueDefaultDuration = 1 * time.Second
)

func (s *Service) InitQueue(ctx context.Context) {
	for container, _ := range QueueContainer {
		s.wg.Add(1)
		go func(container string) {
			defer func() {
				s.wg.Done()
			}()
			ticker := time.NewTicker(queueDefaultDuration)

			for {
				select {
				case <-ticker.C:
					s.queueHandle(container)
				case <-ctx.Done(): // 等待上级通知
					log.Printf("queue Done msg: %#v\n", ctx.Err())
					return
				}
			}
		}(container)
	}
}

func (s *Service) queueHandle(queueContainer string) {
	queueKey := s.GetQueueKey(queueContainer)
	//fmt.Printf("queue handle container#%s\n", queueKey)

	// 使用BLPop会影响程序安全退出，暂时用LPop
	// jobIds = [queueKey, jobId]
	//jobIds, err := s.BLPopFromQueue(queueKey)
	//fmt.Println(err)
	//if err != nil {
	//	return
	//}
	//if len(jobIds) < 1 {
	//	return
	//}
	//
	//jobId := jobIds[1]

	// 使用LPop
	jobId, err := s.LPopFromQueue(queueKey)
	if err != nil {
		return
	}

	s.ch <- jobId
	fmt.Printf("%s chan <- #%s#\n", queueKey, jobId)
	return
}

func (s *Service) GetQueueKey(container string) string {
	return fmt.Sprintf("%s%s", QueuePrefix, QueueContainer[container])
}

func (s *Service) PushToQueue(queue string, jobId string) (err error) {
	err = s.dao.PushQueue(queue, jobId)
	return
}

func (s *Service) BLPopFromQueue(queue string) (jobIds []string, err error) {
	jobIds, err = s.dao.BLPopQueue(queue)
	return
}

func (s *Service) LPopFromQueue(queue string) (jobId string, err error) {
	jobId, err = s.dao.LPopQueue(queue)
	return
}