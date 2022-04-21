package service

import (
	"context"
	"fmt"
	"log"
	"time"
)

const (
	ChanConsumerNum = 100
)

var ch chan string

func (s *Service) InitConsumer(ctx context.Context)  {
	defer func() {
		s.wg.Done()
	}()
	s.wg.Add(1)

	ch = make(chan string, 1000)

	// 取通道job推送给业务方
	go s.consume(ctx)

	// 取队列消息进通道
	for {
		queueKey := s.GetQueueKey()
		// jobIds = [queueKey, jobId]
		jobIds, err := s.PopFromQueue(queueKey)
		if err != nil {
			return
		}
		if len(jobIds) < 1 {
			return
		}

		jobId := jobIds[1]
		
		select {
		case ch <- jobId:
			fmt.Printf("chan <- #%s#\n", jobId)
			time.Sleep(5e8)
		case <-ctx.Done(): // 等待上级通知
			log.Printf("pop from queue Done msg: %#v", ctx.Err())
			return
		}
	}
}

func (s *Service) consume(ctx context.Context)  {
	for i := 0; i < ChanConsumerNum; i++ {

		go func() {
			defer func() {
				s.wg.Done()
			}()
			s.wg.Add(1)

			for {
				select {
				case jobId := <-ch:
					jobInfo, err := s.GetJob(jobId)
					if err != nil || jobInfo == nil {
						return
					}

					if jobInfo.TTR > 0 {
						// 进ttr bucket
						ttrBucketName := s.GetTtrBucket(jobInfo.Id)
						timestamp := jobInfo.Delay + time.Now().Unix()
						err = s.PushToBucket(ttrBucketName, timestamp, jobInfo.Id)
						if err != nil {
							// TODO 错误处理
						}
					}

					// 推送给业务方
					fmt.Printf("#%s# <- chan\n", jobId)
				case <-ctx.Done(): // 等待上级通知
					log.Printf("consume Done msg: %#v", ctx.Err())
					return
				}

				time.Sleep(5e8)
			}
		}()
	}
}