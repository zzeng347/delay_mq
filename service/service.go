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