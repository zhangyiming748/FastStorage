package storage

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func SetRedis(host, port, password string) {
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password, // 没有密码则为空字符串
		DB:       0,        // 使用默认 DB
	})

	// 测试连接
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	log.Println("Successfully connected to Redis")
	redisClient = rdb
}

func GetRedis() *redis.Client {
	return redisClient
}

