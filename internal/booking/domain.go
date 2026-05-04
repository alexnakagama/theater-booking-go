package booking

import "time"

type Booking struct {
	ID        string
	UserID    string
	ShowID    string
	SeatID    string
	Status    string
	ExpiresAt time.Time
}

type BookingStore interface {
	Book(b Booking) error
	ListBookings(showID string) []Booking
}
