package booking

import (
	"fmt"
	"sync"
)

type ConcurrentStore struct {
	bookings map[string]Booking
	sync.RWMutex
}

func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: map[string]Booking{},
	}
}

func (s *ConcurrentStore) Book(b Booking) error {
	s.Lock()
	defer s.Unlock()

	if _, exists := s.bookings[b.SeatID]; exists {
		return fmt.Errorf("booking with seat ID %s already exists", b.SeatID)
	}
	s.bookings[b.ID] = b
	return nil
}

func (s *ConcurrentStore) ListBookings(showID string) []Booking {
	s.RLock()
	defer s.RUnlock()

	var result []Booking
	for _, b := range s.bookings {
		if b.ShowID == showID {
			result = append(result, b)
		}
	}
	return result
}
