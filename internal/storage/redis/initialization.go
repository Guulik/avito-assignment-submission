package redis

import (
	"Avito_trainee_assignment/internal/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis(c *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Address,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})
	return client
}
