package dto

type ReservationRequest struct {
	HotelID  int    `json:"hotel_id"`
	Checkin  string `json:"checkin"`
	Checkout string `json:"checkout"`
	Email    string `json:"email"`
	UserID   int    `json:"user_id"`
}
