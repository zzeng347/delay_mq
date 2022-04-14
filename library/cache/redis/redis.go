package redis

import (
	goredis "github.com/go-redis/redis"
)

type Config struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func NewMyRedis(c *Config) (rdb *goredis.Client) {
	rdb = goredis.NewClient(&goredis.Options{
		Addr:      c.Host+":"+c.Port,
		Password:  c.Password,
		DB:        c.DB,
		PoolSize:  20, // TODO 连接池大小
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	return
}