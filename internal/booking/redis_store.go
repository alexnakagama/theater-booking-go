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

func parseSession(val string) (Booking, error) {
	var data Booking
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		return Booking{}, err
	}
	return Booking{
		ID:     data.ID,
		ShowID: data.ShowID,
		SeatID: data.SeatID,
		UserID: data.UserID,
		Status: data.Status,
	}, nil
}

func (s *RedisStore) ListBookings(movieID string) []Booking {
	pattern := fmt.Sprintf("seat:%s:*", movieID)
	var sessions []Booking

	ctx := context.Background()
	iter := s.rdb.Scan(ctx, 0, pattern, 0).Iterator()

	for iter.Next(ctx) {
		key := iter.Val()
		val, err := s.rdb.Get(ctx, key).Result()
		if err != nil {
			log.Printf("error getting booking for key %s: %v", key, err)
			continue
		}

		session, err := parseSession(val)
		if err != nil {
			log.Printf("error parsing booking for key %s: %v", key, err)
			continue
		}

		sessions = append(sessions, session)
	}

	return sessions
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
