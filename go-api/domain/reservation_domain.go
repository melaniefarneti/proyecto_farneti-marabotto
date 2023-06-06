package domain

type ReservationRequest struct {
	HotelID    int    `json:"hotel_id"`
	Checkin    string `json:"checkin"`
	Checkout   string `json:"checkout"`
	ClientName string `json:"client_name"`
}
