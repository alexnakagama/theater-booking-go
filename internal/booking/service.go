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

func (s *Service) Book(userID, showID, seatID string) error {
	booking := Booking{
		ID:     util.GenerateRandomString(),
		UserID: userID,
		ShowID: showID,
		SeatID: seatID,
		Status: "booked",
	}
	return s.store.Book(booking)
}
