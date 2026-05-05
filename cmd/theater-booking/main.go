package main

import (
	"log"
	"net/http"

	"github.com/alexnakagama/theater-booking-go/internal/adapters/redis"
	"github.com/alexnakagama/theater-booking-go/internal/booking"
	"github.com/alexnakagama/theater-booking-go/util"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /shows", listShows)

	mux.Handle("GET /", http.FileServer(http.Dir("static")))

	store := booking.NewRedisStore(redis.NewClient("localhost:6379"))
	svc := booking.NewService(store)

	bookingHandler := booking.NewHandler(svc)

	mux.HandleFunc("GET /shows/{showID}/seats", bookingHandler.ListSeats)
	mux.HandleFunc("POST /shows/{showID}/seats/{seatID}/hold", bookingHandler.HoldSeat)

	mux.HandleFunc("PUT /sessions/{sessionID}/confirm", bookingHandler.ConfirmSession)
	mux.HandleFunc("DELETE /sessions/{sessionID}", bookingHandler.ReleaseSession)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

var shows = []showsResponse{
	{ID: "inception", Title: "Inception", Rows: 5, SeatsPerRow: 8},
	{ID: "dune", Title: "Dune: Part Two", Rows: 4, SeatsPerRow: 6},
}

func listShows(w http.ResponseWriter, r *http.Request) {
	util.WriteJSON(w, http.StatusOK, shows)
}

type showsResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Rows        int    `json:"rows"`
	SeatsPerRow int    `json:"seats_per_row"`
}
