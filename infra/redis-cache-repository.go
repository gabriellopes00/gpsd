package infra

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheData struct {
	Key   string
	Value string
}

type CacheRepository interface {
	Set(data CacheData) error
	Get(key string) (CacheData, error)
}

var ctx = context.Background()

func (c *RedisCache) Connect() {
	c.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

}

type RedisCache struct {
	client *redis.Client
}

func (c *RedisCache) Set(data CacheData) error {
	err := c.client.Set(ctx, data.Key, data.Value, time.Minute*10).Err()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func (c *RedisCache) Get(key string) (CacheData, error) {
	value, err := c.client.Get(ctx, key).Result()
	if err != nil {
		log.Fatalln(err)
		return CacheData{}, err
	}

	return CacheData{Key: key, Value: value}, err
}
