package cache

import (
	"log"

	"github.com/go-redis/redis"
	goccyjson "github.com/goccy/go-json"
)

func CheckCacheHit[Out any](key string, db *redis.Client) *Out {
	val, err := db.Get(key).Result()
	if err == nil { // chache hit
		var cachedResp Out
		if err := goccyjson.Unmarshal([]byte(val), &cachedResp); err == nil {
			return &cachedResp
		}
		log.Printf("Cache unmarshal failed for key %s: %v", key, err)
		db.Del(key)
	}

	return nil
}
