package cache

import (
	"tarmac/logger"
	"time"

	"github.com/go-redis/redis"
)

type Client struct {
	Addr string
	// ...
}

func Start(cache *Client) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: cache.Addr,
	})
	return client
}

func CheckCacheHit[T any](key string, cache *redis.Client) *T {
	defer logger.Log.TrackTime()()
	data, err := loadJSON[T](key, cache)
	if err != nil {
		logger.Log.Log("Failed cache check:", err)
	} else if data == nil {
		logger.Log.Log("No Cache Hit")
	} else {
		logger.Log.Log("Cache Hit!")
	}
	return data
}

func RefreshCache(data any, key string, ttl time.Duration, cache *redis.Client) {
	defer logger.Log.TrackTime()()
	err := storeJSON(data, key, ttl, cache)
	if err != nil {
		logger.Log.Log("Failed cache refresh:", err)
	} else {
		logger.Log.Log("Cache refreshed")
	}
}
