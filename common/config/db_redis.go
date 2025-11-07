package config

import (
	"context"
	"log"
	"net/url"
	"os"

	"github.com/go-redis/redis/v8"
)

func InitRedis() *redis.Client {
	username := os.Getenv("MEMSTORE_USER")
	password := os.Getenv("MEMSTORE_PASS")
	cluster := os.Getenv("MEMSTORE_INSTANCE")
	uri := "redis://" + url.QueryEscape(username) + ":" +
		url.QueryEscape(password) + "@" + cluster
	opt, err := redis.ParseURL(uri)
	if err != nil {
		log.Printf("Error parsing Redis uri: %s", err.Error())
	}

	rdb := redis.NewClient(opt)
	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Error pinging memstore: %s", err.Error())
	}

	return rdb
}
