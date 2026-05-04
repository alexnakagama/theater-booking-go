package booking

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const defaultHoltTTL = 2 * time.Minute

type RedisStore struct {
	rdb *redis.Client
}

func NewRedisStore(rdb *redis.Client) *RedisStore {
	return &RedisStore{rdb: rdb}
}

func sessionKey(id string) string {
	return fmt.Sprintf("session: %s", id)
}

func (s *RedisStore) Book(b Booking) error {
	session, err := s.hold(b)
	if err != nil {
		return err
	}

	log.Printf("session booked %v", session)

	return nil
}

func (s *RedisStore) ListBookings(movieID string) []Booking {
	return []Booking{}
}

func (s *RedisStore) hold(b Booking) (Booking, error) {
	id := uuid.New().String()
	now := time.Now()
	ctx := context.Background()
	key := fmt.Sprintf("seat:%s:%s", b.ShowID, b.SeatID)

	b.ID = id
	value, _ := json.Marshal(b)

	res := s.rdb.SetArgs(ctx, key, value, redis.SetArgs{
		Mode: "NX",
		TTL:  defaultHoltTTL,
	})

	result, err := res.Result()
	ok := err == nil && result == "OK"
	if !ok {
		return Booking{}, errors.New("seat is already booked")
	}

	return Booking{
		ID:        id,
		UserID:    b.UserID,
		ShowID:    b.ShowID,
		SeatID:    b.SeatID,
		Status:    "held",
		ExpiresAt: now.Add(defaultHoltTTL),
	}, nil
}
