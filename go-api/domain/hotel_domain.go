package domain

type HotelRequest struct {
	Name        string `json:"nombre"`
	Photo       string `json:"foto"`
	Description string `json:"descripcion"`
	Location    string `json:"ubicacion"`
	Rooms       int    `json:"habitaciones"`
}
