package dao

// cache存取

import (
	"delay_mq_v2/model"
	"encoding/json"
	"github.com/go-redis/redis"
)

func (dao *Dao) GetJob(key string) (j *model.PushJobReq, err error) {
	val, err := dao.Redis.Get(key).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(val), &j)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (dao *Dao) SetJob(key string, job *model.PushJobReq) (err error) {
	jobJson, err := json.Marshal(job)
	if err != nil {
		return err
	}
	err = dao.Redis.Set(key, jobJson, 0).Err()
	return
}

func (dao *Dao) PushBucket(bucket string, timestamp float64, jobId string) (err error) {
	value := redis.Z{Score: timestamp, Member: jobId}
	_, err = dao.Redis.ZAdd(bucket, value).Result()
	return
}

