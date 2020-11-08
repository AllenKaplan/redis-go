package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Go Redis Tutorial")

	redisAddr := os.Getenv("REDIS")
	if len(redisAddr) == 0 {
		redisAddr = "localhost:6379"
	}

	fmt.Println("Redis Addr: ", redisAddr)

	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

}
