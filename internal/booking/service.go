package booking

import "github.com/alexnakagama/theater-booking-go/util"

type Service struct {
	store BookingStore
}

func NewService(store BookingStore) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) Book(b Booking) error {
	b.ID = util.GenerateRandomString()
	b.Status = "booked"
	return s.store.Book(b)
}

func (s *Service) ListBookings(showID string) []Booking {
	return s.store.ListBookings(showID)
}
