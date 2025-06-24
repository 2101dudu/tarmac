package cache

import (
	"log"
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
	data, err := loadJSON[T](key, db)
	if err != nil {
		log.Println("Failed cache check:", err)
	} else if data == nil {
		log.Println("No Cache Hit")
	} else {
		log.Println("Cache Hit!")
	}
	return data
}

func RefreshCache(data any, key string, ttl time.Duration, db *redis.Client) {
	err := storeJSON(data, key, ttl, db)
	if err != nil {
		log.Println("Failed cache refresh:", err)
	} else {
		log.Println("Cache refreshed")
	}
}
