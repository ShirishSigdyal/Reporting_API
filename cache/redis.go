package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(redisURL string) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})
	return &RedisCache{client: client}
}

func (c *RedisCache) Get(key string) (string, error) {
	return c.client.Get(context.Background(), key).Result()
}

func (c *RedisCache) Set(key string, value string, ttl time.Duration) error {
	return c.client.Set(context.Background(), key, value, ttl).Err()
}
