package booking

import (
	"testing"

	"github.com/google/uuid"
)

func TestMemoryStore_Book(t *testing.T) {
	store := NewMemoryStore()
	service := NewService(store)

	err := service.Book(Booking{
		UserID: "user1",
		ShowID: "show1",
		SeatID: "seat1",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	bookings := store.ListBookings("show1")
	if len(bookings) != 1 {
		t.Fatalf("expected 1 booking, got %d", len(bookings))
	}

	if bookings[0].UserID != "user1" || bookings[0].SeatID != "seat1" {
		t.Fatalf("booking details do not match expected values")
	}
}

func TestMemoryStore_ListBookings(t *testing.T) {
	store := NewMemoryStore()
	service := NewService(store)

	// Create multiple bookings for the same show
	err := service.Book(Booking{
		UserID: uuid.New().String(),
		ShowID: "show1",
		SeatID: "A1",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err2 := service.Book(Booking{
		UserID: uuid.New().String(),
		ShowID: "show1",
		SeatID: "A2",
	})
	if err2 != nil {
		t.Fatalf("expected no error, got %v", err2)
	}

	err3 := service.Book(Booking{
		UserID: uuid.New().String(),
		ShowID: "show2",
		SeatID: "B1",
	})
	if err3 != nil {
		t.Fatalf("expected no error, got %v", err3)
	}

	bookings := store.ListBookings("show1")
	if len(bookings) != 2 {
		t.Fatalf("expected 2 bookings for show1, got %d", len(bookings))
	}

	bookings = store.ListBookings("show2")
	if len(bookings) != 1 {
		t.Fatalf("expected 1 booking for show2, got %d", len(bookings))
	}
}
