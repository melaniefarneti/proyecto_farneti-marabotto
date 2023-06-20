package dto

type AmenityRequest struct {
	Nombre  string `json:"nombre"`
	HotelID int64  `json:"hotel_id"`
}
