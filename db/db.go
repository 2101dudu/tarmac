package db

import (
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
