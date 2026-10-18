package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/rooms/", getRoom)
	http.HandleFunc("/api/rooms", listRooms)
	http.HandleFunc("/health", healthHandler)
	serverErr := http.ListenAndServe(":8080", nil)
	fmt.Println(serverErr)

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Some message")
}
