package booking

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
	{ID: "show1", Name: "The Phantom of the Opera", Rows: 10, SeatsPerRow: 20},
	{ID: "show2", Name: "Hamilton", Rows: 12, SeatsPerRow: 25},
	{ID: "show3", Name: "Les Misérables", Rows: 15, SeatsPerRow: 30},
}

func ListShowsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shows)
}

func ListSeatsHandler(w http.ResponseWriter, r *http.Request) {

}
