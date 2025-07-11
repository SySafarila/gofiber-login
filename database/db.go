package database

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Redis *redis.Client
var Ctx = context.Background()
var IsRedisConnected bool

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := Redis.Ping(Ctx).Result()
	if err != nil {
		fmt.Println("Redis not connected")
		IsRedisConnected = false
		return
	}
	IsRedisConnected = true
}
