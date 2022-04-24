package service

import (
	"context"
	"delay_mq_v2/model"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

const (
	ChanConsumerNum = 100
	ConsumerUrl = "http://127.0.0.1:8112/biz"
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
					fmt.Printf("#%s# <- chan\n", jobId)
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
							log.Printf("push to bucket err#%s\n", err.Error())
						}
					}

					// 推送给业务方
					body, err := s.HttpClient.Post(ConsumerUrl, jobInfo, "application/json; charset=utf-8")
					if err != nil {
						log.Printf("读取body错误#%s", err.Error())
						return
					}

					reply := &model.Reply{}
					err = json.Unmarshal(body, reply)
					if err != nil {
						log.Printf("解析json失败#%s", err.Error())
						return
					}

					fmt.Printf("post biz res body#%v\n", string(body))
					fmt.Printf("reply#%d\n", reply.Code)

					if reply.Code == 1 {
						// 消费成功 删除job
						err = s.DelJob(jobId)
						if err != nil {
							log.Printf("del job err#%s\n", err.Error())
						}
						fmt.Println("del job success")
					}

					return
				case <-ctx.Done(): // 等待上级通知
					log.Printf("consume Done msg: %#v", ctx.Err())
					return
				}

				time.Sleep(5e8)
			}
		}()
	}
}