package booking

import (
	"testing"
)

func TestMemoryStore_Book(t *testing.T) {
	store := NewMemoryStore()
	service := NewService(store)

	err := service.Book("user1", "show1", "seat1")
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
	service.Book("user1", "show1", "seat1")
	service.Book("user2", "show1", "seat2")
	service.Book("user3", "show2", "seat1")

	bookings := store.ListBookings("show1")
	if len(bookings) != 2 {
		t.Fatalf("expected 2 bookings for show1, got %d", len(bookings))
	}

	bookings = store.ListBookings("show2")
	if len(bookings) != 1 {
		t.Fatalf("expected 1 booking for show2, got %d", len(bookings))
	}
}
