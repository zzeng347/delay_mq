package dao

import (
	goredis "github.com/go-redis/redis"
)


type Dao struct {
	redis *goredis.Client
}

