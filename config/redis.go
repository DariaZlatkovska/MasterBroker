package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Nie można połączyć z Redis: %v", err)
	}
	log.Println("Połączono z Redis")
}
