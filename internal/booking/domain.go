package booking

type Booking struct {
	ID     string
	UserID string
	ShowID string
	SeatID string
	Status string
}

type BookingStore interface {
	Book(b Booking) error
	ListBookings(showID string) []Booking
}
