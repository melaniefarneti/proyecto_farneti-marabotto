package domain

type AmenityRequest struct {
	Name    string `json:"nombre"`
	HotelID int    `json:"hotel_id"`
}
