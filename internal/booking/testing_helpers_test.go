package booking

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	goredis "github.com/redis/go-redis/v9" // ← alias to avoid collision
)

func newTestRedisStore(t *testing.T) BookingStore {
	t.Helper()

	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("failed to start miniredis: %v", err)
	}
	t.Cleanup(mr.Close)

	client := goredis.NewClient(&goredis.Options{
		Addr: mr.Addr(),
	})
	t.Cleanup(func() { client.Close() })

	return NewRedisStore(client) // returns *RedisStore which satisfies BookingStore
}
