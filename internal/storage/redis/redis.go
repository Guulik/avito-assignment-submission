package redis

import (
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
}

func New(log *slog.Logger, redis *redis.Client) *Cache {
	return &Cache{
		log:   log,
		redis: redis,
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

	banner, err := c.redis.Get(ctx, key).Bytes()
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

	err := c.redis.Set(ctx, key, content, time.Minute).Err()
	if err != nil {
		log.Error("failed to save banner to cache", sl.Err(err))
		return err
	}
	return nil
}
