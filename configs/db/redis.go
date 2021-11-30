package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func Client() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return client
}

var Ctx = context.Background()

func Connect(client *redis.Client) error {
	pong, err := client.Ping(Ctx).Result()
	if err != nil && pong != "PONG" {
		return err
	}
	fmt.Println("Success connect Redis")

	return nil
}
