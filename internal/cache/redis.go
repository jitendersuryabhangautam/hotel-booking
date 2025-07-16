package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"), // e.g., "localhost:6379"
		Password: "",                      // set if needed
		DB:       0,                     // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := Client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("❌ Redis connection failed: %v", err))
	}
	fmt.Println("✅ Connected to Redis")
}

func Close() {
	if Client != nil {
		_ = Client.Close()
	}
}
