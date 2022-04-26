package dao

// cache存取

import (
	"delay_mq_v2/model"
	"encoding/json"
	"github.com/go-redis/redis"
)

func (dao *Dao) GetJob(key string) (j *model.JobResp, err error) {
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

func (dao *Dao) DelJob(key string) (err error) {
	err = dao.Redis.Del(key).Err()
	return
}

func (dao *Dao) SetJob(key string, job *model.JobResp) (err error) {
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

func (dao *Dao) ZRangeBucket(bucket string, start, stop int64) (ret []redis.Z, err error) {
	ret, err = dao.Redis.ZRangeWithScores(bucket, start, stop).Result()
	return
}

func (dao *Dao) PushQueue(queue string, jobId string) (err error) {
	_, err = dao.Redis.RPush(queue, jobId).Result()
	return
}

func (dao *Dao) RemoveInBucket(bucket string, jobId string) (err error) {
	_, err = dao.Redis.ZRem(bucket, jobId).Result()
	return
}

func (dao *Dao) BLPopQueue(queue string) (jobIds []string, err error) {
	jobIds, err = dao.Redis.BLPop(0, queue).Result()
	return
}

func (dao *Dao) LPopQueue(queue string) (jobId string, err error) {
	jobId, err = dao.Redis.LPop(queue).Result()
	return
}