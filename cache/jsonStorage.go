package cache

import (
	"tarmac/json"
	"tarmac/logger"
	"time"

	"github.com/go-redis/redis"
)

func storeJSON(data any, key string, ttl time.Duration, db *redis.Client) error {
	defer logger.Log.TrackTime()()
	compressedData, err := json.Compress(data)
	if err != nil {
		return err
	}
	return db.Set(key, compressedData, ttl).Err()
}

func loadJSON[T any](key string, db *redis.Client) (*T, error) {
	defer logger.Log.TrackTime()()
	compressedData, err := db.Get(key).Bytes()
	if err != nil { // "It returns redis.Nil error when key does not exist."
		return nil, nil
	}
	return json.Uncompress[T](compressedData)
}
