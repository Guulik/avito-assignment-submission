package redis

import (
	"Banner_Infrastructure/internal/configure"
	"github.com/redis/go-redis/v9"
)

func InitRedis(c *configure.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Address,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})
	return client
}
