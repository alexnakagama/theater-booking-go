package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexnakagama/theater-booking-go/internal/booking"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.FileServer(http.Dir("static")))
	mux.HandleFunc("GET /shows/", booking.ListShowsHandler)

	fmt.Printf("server is running on port: 8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
