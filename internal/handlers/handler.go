package handlers

import (
	"encoding/json"
	"net/http"
)

type showsResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Rows        int    `json:"rows"`
	SeatsPerRow int    `json:"seats_per_row"`
}

var shows = []showsResponse{
	{ID: "show1", Name: "The Phantom of the Opera"},
	{ID: "show2", Name: "Hamilton"},
	{ID: "show3", Name: "Les Misérables"},
}

func ListShowsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shows)
}
