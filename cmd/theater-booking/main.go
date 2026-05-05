package main

import (
	"fmt"
	"log"
	"net/http"

	h "github.com/alexnakagama/theater-booking-go/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.FileServer(http.Dir("static")))
	mux.HandleFunc("GET /shows/", h.ListShowsHandler)

	fmt.Printf("server is running on port: 8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
