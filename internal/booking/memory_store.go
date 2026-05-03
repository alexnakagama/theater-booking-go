package booking

import "fmt"

type MemoryStore struct {
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

func (s *MemoryStore) Book(b Booking) error {
	if _, exists := s.bookings[b.SeatID]; exists {
		return fmt.Errorf("booking with seat ID %s already exists", b.SeatID)
	}
	s.bookings[b.ID] = b
	return nil
}

func (s *MemoryStore) ListBookings(showID string) []Booking {
	var result []Booking
	for _, b := range s.bookings {
		if b.ShowID == showID {
			result = append(result, b)
		}
	}
	return result
}
