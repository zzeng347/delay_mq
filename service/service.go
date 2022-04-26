package service

import (
	"context"
	"delay_mq_v2/conf"
	"delay_mq_v2/dao"
	"delay_mq_v2/library/net/http"
	goredis "github.com/go-redis/redis"
	"sync"
)

type Service struct {
	c				*conf.Config
	dao				*dao.Dao
	Redis			*goredis.Client
	HttpClient		*http.Client
	wg				sync.WaitGroup
	ch				chan string
}

var s *Service

// New new a Service and return.
func New(c *conf.Config) *Service {
	s = &Service{
		c:				c,
		dao:			dao.New(c),
		HttpClient:		http.NewHttpClient(c.HTTPCLIENT),
		ch: make(chan string, 1000),
	}
	s.Redis = s.dao.Redis
	return s
}

func (s *Service) Run(ctx context.Context)  {
	// init bucket
	go InitBucket(ctx, s)

	// init queue
	go s.InitQueue(ctx)

	// 取通道job推送给业务方
	go s.InitConsume(ctx)
}

func (s *Service) Close() {
	defer s.wg.Wait()
	//s.dao.Redis.Close()
}