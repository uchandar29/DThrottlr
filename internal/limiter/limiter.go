package limiter

import (
	"context"
	_ "embed"

	"github.com/redis/go-redis/v9"
)

//go:embed scripts/token_bucket.lua
var tokenBucketScript string

type Limiter struct {
	client     *redis.Client
	script     *redis.Script
	capacity   int
	refillRate int
}

func New(client *redis.Client, capacity, refillRate int) *Limiter {
	return &Limiter{
		client:     client,
		script:     redis.NewScript(tokenBucketScript),
		capacity:   capacity,
		refillRate: refillRate,
	}
}

func (l *Limiter) Allow(ctx context.Context, clientID string) (bool, int, error) {
	key := "ratelimit:" + clientID
	res, err := l.script.Run(ctx, l.client, []string{key}, l.capacity, l.refillRate, 1).Result()
	if err != nil {
		return false, 0, err
	}

	vals := res.([]interface{})
	allowed := vals[0].(int64) == 1
	remaining := int(vals[1].(int64))

	return allowed, remaining, nil
}
