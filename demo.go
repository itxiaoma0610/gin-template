package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	var rds *redis.Client
	rds = redis.NewClient(&redis.Options{
		Addr:     "118.89.66.104:6379",
		Password: "rootpassword",
	})
	pong, err := rds.Get(context.Background(), "hello").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("key does not exist")
	}
	println(pong)
}
