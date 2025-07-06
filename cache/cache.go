package cache

import (
	"tarmac/logger"
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
	// defer logger.Log.TrackTime()()
	data, err := loadJSON[T](cacheService, key)
	if err != nil {
		logger.Log.Log("Failed cache check:", err)
	} else if data == nil {
		logger.Log.Log("No Cache Hit")
	} else {
		logger.Log.Log("Cache Hit!")
	}
	return data
}

func (c *Service) RefreshCache(key string, data any, ttl time.Duration) {
	// defer logger.Log.TrackTime()()
	err := c.storeJSON(key, data, ttl)
	if err != nil {
		logger.Log.Log("Failed cache refresh:", err)
	} else {
		logger.Log.Log("Cache refreshed")
	}
}

func (c *Service) RemoveCache() {
	if err := c.redisClient.FlushAll().Err(); err != nil {
		logger.Log.Log("Failed to flush Redis:", err)
	} else {
		logger.Log.Log("Redis cache cleared")
	}
}
