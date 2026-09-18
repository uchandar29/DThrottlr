package limiter

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

func TestAllow_DepletesAndRefills(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	rdb.FlushDB(context.Background()) // clean slate BEFORE
	defer rdb.FlushDB(context.Background())

	lim := New(rdb, 3, 1)

	for i := 0; i < 3; i++ {
		allowed, _, err := lim.Allow(context.Background(), "test")
		require.NoError(t, err)
		require.True(t, allowed)
	}

	allowed, _, _ := lim.Allow(context.Background(), "test")
	require.False(t, allowed) // 4th request should be denied
}
