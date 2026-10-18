package main

type Room struct {
	ID                 int    `json:"id"`
	Number             string `json:"number"`
	Type               string `json:"type"`
	PricePerNightMinor int    `json:"price_per_night_minor"`
}

type BookingRequest struct {
	RoomID     int    `json:"room_id"`
	GuestName  string `json:"guest_name"`
	GuestEmail string `json:"guest_email"`
	CheckIn    string `json:"check_in"`
	CheckOut   string `json:"check_out"`
}

type Booking struct {
	ID               string `json:"id"`
	RoomID           int    `json:"room_id"`
	GuestName        string `json:"guest_name"`
	GuestEmail       string `json:"guest_email"`
	CheckIn          string `json:"check_in"`
	CheckOut         string `json:"check_out"`
	NightlyRateMinor int    `json:"nightly_rate_minor"`
	TotalPriceMinor  int    `json:"total_price_minor"`
	Status           string `json:"status"`
	CreatedAt       string `json:"created_at"`
}
