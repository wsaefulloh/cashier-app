package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

func Client() *redis.Client {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error on loading .env file")
	}
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
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
