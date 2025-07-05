package cache

import (
	"tarmac/json"
	"time"
)

func (c *Service) storeJSON(key string, data any, ttl time.Duration) error {
	// defer logger.Log.TrackTime()()
	compressedData, err := json.Compress(data)
	if err != nil {
		return err
	}
	return c.redisClient.Set(key, compressedData, ttl).Err()
}

func loadJSON[T any](cacheService *Service, key string) (*T, error) {
	// defer logger.Log.TrackTime()()
	compressedData, err := cacheService.redisClient.Get(key).Bytes()
	if err != nil { // "It returns redis.Nil error when key does not exist."
		return nil, nil
	}
	return json.Uncompress[T](compressedData)
}
