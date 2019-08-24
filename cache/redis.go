package cache

import (
  "fmt"
  "time"
  "context"

  "github.com/go-redis/redis"
)

var RidesKeys = map[string]string{
  "userInfo": "user_info", // 用户登录
}

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

// 设置hast 类型缓存
func (c *Cache) HMSet(key string, fields map[string]interface{}) {
  fmt.Println(key)
  fmt.Println(fields)
  c.client.HMSet(key, fields)
}

// 获取hash全部参数
func (c *Cache) HGetAll(hash string) (map[string]string, bool) {
  data, err := c.client.HGetAll(hash).Result()
  if err != nil || len(data) <= 0 {
    return map[string]string{}, false
  }
  return data, true
}