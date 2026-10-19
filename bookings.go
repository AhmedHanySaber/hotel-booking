package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var input BookingRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Could Not decode request", http.StatusBadRequest)
		return
	}
	req, err := json.Marshal(input)
	if err != nil {
		http.Error(w, "Could not encode rooms", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(req))
}
