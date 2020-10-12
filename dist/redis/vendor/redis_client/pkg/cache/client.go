package cache

import (
	"context"
	"github.com/go-redis/redis"
	"log"
	"time"
)

func NewClient(db int, url string) *redis.Client {
	log.Println("connecting to the redis", url, "db", db)
	cl := redis.NewClient(&redis.Options{
		Addr: url,
		DB:   db,
	})
	log.Println("connected to redis")
	return cl
}

var get = func(client *redis.Client, ctx context.Context, params ...string) string {
	res := client.Get(ctx, params[0])
	return res.String()
}
var set = func(client *redis.Client, ctx context.Context, params ...string) string {
	var duration time.Duration
	if len(params) == 3 {
		converted, err := time.ParseDuration(params[2])
		if err != nil {
			log.Println("ignoring string parsing, default duration will be 0", err)
		} else {
			duration = converted
		}
	}
	res := client.Set(ctx, params[0], params[1], duration)
	return res.String()
}
var del = func(client *redis.Client, ctx context.Context, params ...string) string {
	res := client.Del(ctx, params...)
	return res.String()
}

var mget = func(client *redis.Client, ctx context.Context, params ...string) string {
	res := client.MGet(ctx, params...)
	return res.String()
}
var keys = func(client *redis.Client, ctx context.Context, params ...string) string {
	res := client.Keys(ctx, params[0])
	return res.String()
}

const (
	Get  = "get"
	Set  = "set"
	Del  = "del"
	Mget = "mget"
	Keys = "keys"
)

var Commands = make(map[string]func(client *redis.Client, ctx context.Context, params ...string) string)

func init() {
	Commands[Get] = get
	Commands[Set] = set
	Commands[Del] = del
	Commands[Mget] = mget
	Commands[Keys] = keys
}
