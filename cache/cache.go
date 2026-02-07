package cache

import (
	"log/slog"
	"time"

	"github.com/go-redis/redis"
)

type Client struct {
	Addr            string
	ShortCacheTime  time.Duration
	MediumCacheTime time.Duration
	LongCacheTime   time.Duration
}

type CacheTimes struct {
	ShortCacheTime  time.Duration
	MediumCacheTime time.Duration
	LongCacheTime   time.Duration
}

type Service struct {
	CacheTimes  *CacheTimes
	redisClient *redis.Client
}

func (c *Client) NewCacheService() *Service {
	r := redis.NewClient(&redis.Options{
		Addr: c.Addr,
	})
	return &Service{
		CacheTimes: &CacheTimes{
			ShortCacheTime:  c.ShortCacheTime,
			MediumCacheTime: c.MediumCacheTime,
			LongCacheTime:   c.LongCacheTime,
		},
		redisClient: r,
	}
}

func CheckCacheHit[T any](cacheService *Service, key string) *T {
	data, err := loadJSON[T](cacheService, key)
	if err != nil {
		slog.Warn("Failed cache check:", "Error", err)
	} else if data == nil {
		slog.Debug("No Cache Hit")
	} else {
		slog.Debug("Cache Hit!")
	}
	return data
}

func (c *Service) RefreshCache(key string, data any, ttl time.Duration) {
	err := c.storeJSON(key, data, ttl)
	if err != nil {
		slog.Warn("Failed cache refresh:", "Error", err)
	} else {
		slog.Debug("Cache refreshed")
	}
}

func (c *Service) RemoveCache() {
	if err := c.redisClient.FlushAll().Err(); err != nil {
		slog.Warn("Failed to flush Redis:", "Error", err)
	} else {
		slog.Debug("Redis cache cleared")
	}
}
