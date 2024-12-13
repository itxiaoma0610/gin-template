package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "118.89.66.104:6379",
		Password: "rootpassword",
		DB:       0,
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalf("连接 Redis 失败: %v", err)
	}
	fmt.Println(client.Get(context.Background(), "biz#verification#19129211344").Err())
	fmt.Println(client.Get(context.Background(), "biz#verification#19129211344").Result())
}
