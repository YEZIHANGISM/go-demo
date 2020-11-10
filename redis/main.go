package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func RedisClient() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("connection failed, err: ", err)
		return
	}
	fmt.Println("connection success")

}

func main() {
	RedisClient()
}
