package dao

import (
	"delay_mq_v2/conf"
	"delay_mq_v2/library/cache/redis"
	goredis "github.com/go-redis/redis"
)


type Dao struct {
	c		*conf.Config
	Redis	*goredis.Client
}

func New(c *conf.Config) (d *Dao) {
	d = &Dao{
		c:			c,
		Redis:		redis.NewMyRedis(c.REDIS),
	}
	return
}