package service

import (
	"context"
	"delay_mq_v2/conf"
	"delay_mq_v2/dao"
	goredis "github.com/go-redis/redis"
	"sync"
)

type Service struct {
	c         *conf.Config
	dao       *dao.Dao
	Redis     *goredis.Client
	wg sync.WaitGroup
}

var s *Service

// New new a Service and return.
func New(c *conf.Config) *Service {
	s = &Service{
		c:      c,
		dao:    dao.New(c),
	}
	s.Redis = s.dao.Redis
	return s
}

func (s *Service) Run(ctx context.Context)  {
	// init bucket
	go InitBucket(ctx, s)

	// init consumer
	go s.InitConsumer(ctx)
}