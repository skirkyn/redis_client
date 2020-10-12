package cache

import (
	"github.com/go-redis/redis"
	"log"
)

func NewClient(db int, url string) *redis.Client {
	log.Println("connecting to the redis", url, "db", db)
	cl := redis.NewClient(&redis.Options{
		Addr: url,
		DB:   db,
	})
	log.Println("connected")
	return cl
}
