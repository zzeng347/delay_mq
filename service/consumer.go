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
	ChanConsumerNum = 2
	RetryCount = 3
	ExecerUrl = "http://127.0.0.1:8112/"
)

func (s *Service) InitConsume(ctx context.Context)  {
	for i := 0; i < ChanConsumerNum; i++ {
		s.wg.Add(1)
		go func(i int) {
			fmt.Printf("chan consumer num%d\n", i)
			defer func() {
				s.wg.Done()
			}()

			for {
				select {
				case jobId := <-s.ch:
					fmt.Printf("consumer%d #%s# <- chan\n", i, jobId)
					jobResp, err := s.GetJob(jobId)
					if err != nil || jobResp == nil {
						continue
					}

					// 判断推送给业务方之后，是否可删除job
					isDelJob := jobResp.RetryCount >= RetryCount || jobResp.TTR <= 0

					if jobResp.RetryCount < RetryCount && jobResp.TTR > 0 {
						// 进ttr bucket
						ttrBucketName := s.GetTtrBucket(jobResp.Id)
						timestamp := jobResp.Delay + time.Now().Unix()
						err = s.PushToBucket(ttrBucketName, timestamp, jobResp.Id)
						if err != nil {
							// TODO 错误处理
							log.Printf("push to bucket err#%s\n", err.Error())
						}

						jobResp.RetryCount++
						err = s.SetJob(jobResp)
						if err != nil {
							// TODO 错误处理
						}
						fmt.Printf("job #%s# retry count+1 = %d\n", jobResp.Id, jobResp.RetryCount)
					}

					// 推送给业务方
					//fmt.Printf("post biz res body#%v\n", string(body))
					body, err := s.HttpClient.Post(ExecerUrl + jobResp.Route, jobResp, "application/json; charset=utf-8")
					if err != nil {
						log.Printf("读取body错误#%s", err.Error())
						s.DeleteJob(isDelJob, jobId)
						continue
					}

					reply := &model.Reply{}
					err = json.Unmarshal(body, reply)
					if err != nil {
						log.Printf("解析json失败#%s", err.Error())
						s.DeleteJob(isDelJob, jobId)
						continue
					}
					fmt.Printf("reply.data #%s\n", string(reply.Data))

					if reply.Code == 1 {
						// 消费成功 删除job
						err = s.DelJob(jobId)
						if err != nil {
							log.Printf("del job err#%s\n", err.Error())
						}
					} else {
						s.DeleteJob(isDelJob, jobId)
					}

				case <-ctx.Done(): // 等待上级通知
					log.Printf("consume Done msg: %#v", ctx.Err())
					return
				}

				time.Sleep(5e8)
			}
		}(i)
	}
}

func (s *Service) DeleteJob(isDel bool, jobId string) (err error) {
	if isDel == true {
		err = s.DelJob(jobId)
		if err != nil {
			log.Printf("del job err#%s\n", err.Error())
		}
	}
	return
}