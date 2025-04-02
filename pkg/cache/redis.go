package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/iqthuc/sport-shop/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient() (*redis.Client, error) {
	redisAddr := os.Getenv(config.Env.RedisAddr)
	redisPassword := os.Getenv(config.Env.RedisPassword)
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         redisAddr,
		Password:     redisPassword,
		DB:           0,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 2,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return rdb, nil
}
