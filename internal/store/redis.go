package store

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var RDB *redis.Client

func InitRedis() {
    RDB = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })
}

func SaveIP(ip string, score int, ttl time.Duration) error {
    return RDB.Set(ctx, "ip:"+ip, score, ttl).Err()
}

func GetIPScore(ip string) (string, error) {
    return RDB.Get(ctx, "ip:"+ip).Result()
}