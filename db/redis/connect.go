package redis

import (
	"github.com/go-redis/redis"
	"../../config"
	"time"
)

var redisClient *redis.Client

func Init() {
	var err error
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.GetString("redis.addr"),
		Password: config.AppConfig.GetString("redis.password"),
		DB:       0,
	})
	if _, err = redisClient.Ping().Result(); err != nil {
		panic(err)
	}
	return
}

func GetValue(key string) (interface{}, error) {
	var (
		val interface{}
		err error
	)

	if val, err = redisClient.Get(key).Result(); err != nil {
		return val, err
	}

	return val, nil
}

func SetValue(key string, value interface{}) error {
	var err error
	if err = redisClient.Set(key, value, 0).Err(); err != nil {
		return err
	}

	return nil
}

func SetValueExpire(key string, value interface{}, ex time.Duration) error {
	var err error
	if err = redisClient.Set(key, value, ex).Err(); err != nil {
		return err
	}

	return nil
}
func DelKey(key string) error {
	var err error
	if _, err = redisClient.Del(key).Result(); err != nil {
		return err
	}
	return nil
}
