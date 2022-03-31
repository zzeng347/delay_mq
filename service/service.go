package service

import (
	"delay_mq_v2/conf"
	"delay_mq_v2/dao"
	goredis "github.com/go-redis/redis"
)

type Service struct {
	c         *conf.Config
	dao       *dao.Dao
	Redis     *goredis.Client
}

// New new a Service and return.
func New(c *conf.Config) (s *Service) {
	s = &Service{
		c:      c,
		dao:    dao.New(c),
	}
	s.Redis = s.dao.Redis
	return s
}