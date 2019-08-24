package cache

import (
  // "fmt"
  "time"
  "context"

  "github.com/go-redis/redis"
)

type Cache struct {
  client redis.UniversalClient
  ttl    time.Duration
}

func NewCache(redisAddress string, password string, ttl time.Duration) (*Cache, error) {
  client := redis.NewClient(&redis.Options{
    Addr: redisAddress,
  })

  err := client.Ping().Err()
  if err != nil {
    return nil, err
  }

  return &Cache{client: client, ttl: ttl}, nil
}

func (c *Cache) Add(ctx context.Context, hash string, query string) {
  c.client.Set(hash, query, c.ttl)
}

func (c *Cache) Get(ctx context.Context, hash string) (string, bool) {
  s, err := c.client.Get(hash).Result()
  if err != nil {
    return "", false
  }
  return s, true
}

func (c *Cache) HGetAll(hash string) (map[string]string, bool) {
  data, err := c.client.HGetAll(hash).Result()
  if err != nil || len(data) <= 0 {
    return map[string]string{}, false
  }
  return data, true
}