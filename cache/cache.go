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

func Start(db *Client) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: db.Addr,
	})
	return client
}

func CheckCacheHit[T any](key string, db *redis.Client) *T {
	defer logger.Log.TrackTime()()
	data, err := loadJSON[T](key, db)
	if err != nil {
		logger.Log.Log("Failed cache check:", err)
	} else if data == nil {
		logger.Log.Log("No Cache Hit")
	} else {
		logger.Log.Log("Cache Hit!")
	}
	return data
}

func RefreshCache(data any, key string, ttl time.Duration, db *redis.Client) {
	defer logger.Log.TrackTime()()
	err := storeJSON(data, key, ttl, db)
	if err != nil {
		logger.Log.Log("Failed cache refresh:", err)
	} else {
		logger.Log.Log("Cache refreshed")
	}
}
