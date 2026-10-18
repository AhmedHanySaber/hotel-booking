package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var rooms = []Room{
	{
		ID:                 101,
		Number:             "101",
		Type:               "single",
		PricePerNightMinor: 50000,
	},
	{
		ID:                 102,
		Number:             "102",
		Type:               "double",
		PricePerNightMinor: 75000,
	},

	{
		ID:                 103,
		Number:             "103",
		Type:               "suite",
		PricePerNightMinor: 100000,
	},
}

func listRooms(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	listRooms, err := json.Marshal(rooms)
	if err != nil {
		http.Error(w, "Could not encode rooms", http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(listRooms))
	}
}
func getRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idText := strings.TrimPrefix(r.URL.Path, "/api/rooms/")
	id, err := strconv.Atoi(idText)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid room ID", http.StatusBadRequest)
		return
	}

	room, found := findRoomByID(rooms, id)
	if !found {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	data, err := json.Marshal(room)
	if err != nil {
		http.Error(w, "Could not encode room", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(data))
}
func findRoomByID(rooms []Room, id int) (Room, bool) {
	for _, room := range rooms {
		if room.ID == id {
			return room, true
		}
	}
	return Room{}, false
}
