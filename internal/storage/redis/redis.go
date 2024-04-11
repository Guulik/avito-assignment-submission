package redis

import (
	"Avito_trainee_assignment/internal/config"
	sl "Avito_trainee_assignment/internal/lib/logger/slog"
	"Avito_trainee_assignment/internal/storage"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"time"
)

var _ storage.BannerCache = (*Cache)(nil)

type Cache struct {
	log   *slog.Logger
	redis *redis.Client
	cfg   *config.Config
}

func New(log *slog.Logger, redis *redis.Client, cfg *config.Config) *Cache {
	return &Cache{
		log:   log,
		redis: redis,
		cfg:   cfg,
	}
}

func (c Cache) GetBannerCached(
	featureId int64,
	tagId int64,
) ([]byte, error) {
	const op = "Cache.GetBannerCached"
	log := c.log.With(
		slog.String("op", op),
	)
	ctx := context.Background()
	key := fmt.Sprintf("%v:%v", featureId, tagId)

	banner, err := c.redis.HGet(ctx, key, "content").Bytes()
	if err != nil {
		log.Warn("failed to get cached banner", sl.Err(err))
		return nil, err
	}

	return banner, err
}

func (c Cache) SetBannerCache(featureId int64,
	tagId int64, content []byte) error {
	const op = "Cache.SetBannerCache"
	log := c.log.With(
		slog.String("op", op),
	)

	ctx := context.Background()
	key := fmt.Sprintf("%v:%v", featureId, tagId)

	err := c.redis.HSet(ctx, key, "content", content).Err()
	if err != nil {
		log.Error("failed to save banner to cache", sl.Err(err))
		return err
	}
	_, err = c.redis.Expire(ctx, fmt.Sprintf("banner:%d:%d", featureId, tagId),
		time.Duration(c.cfg.Redis.TTLMinutes)).Result()
	if err != nil {
		log.Error("Failed to set expiration time for banner: %v", err)
	}
	return nil
}

func (c Cache) DeleteBannerCache(bannerId int64) error {
	const op = "Cache.DeleteBannerCache"
	_ = c.log.With(
		slog.String("op", op),
	)
	//TODO: implement me
	//ctx := context.Background()
	return nil
}
