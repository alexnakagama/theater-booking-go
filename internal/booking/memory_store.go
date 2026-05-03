package booking

import "fmt"

type MemoryStore struct {
	bookings map[string]Booking
}

// NewMemoryStore creates a new instance of MemoryStore.
// It initializes the bookings map to store booking data in memory.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

// Book adds a new booking to the in-memory store.
// It checks if a booking with the same seat ID already exists to prevent double booking.
// If a booking with the same seat ID exists, it returns an error.
// Otherwise, it adds the new booking to the bookings map and returns nil.
func (s *MemoryStore) Book(b Booking) error {
	if _, exists := s.bookings[b.SeatID]; exists {
		return fmt.Errorf("booking with seat ID %s already exists", b.SeatID)
	}
	s.bookings[b.ID] = b
	return nil
}

// ListBookings retrieves all bookings for a specific show ID from the in-memory store.
// It iterates through the bookings map and collects bookings that match the provided show ID.
// The function returns a slice of Booking structs that correspond to the specified show ID.
func (s *MemoryStore) ListBookings(showID string) []Booking {
	var result []Booking
	for _, b := range s.bookings {
		if b.ShowID == showID {
			result = append(result, b)
		}
	}
	return result
}
