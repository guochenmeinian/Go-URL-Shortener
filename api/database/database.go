package database

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

// create context
var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDRESS"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       dbNo,
	})
	return rdb
}
