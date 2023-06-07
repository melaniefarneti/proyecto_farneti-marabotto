package services

import "go-api/dao"

// HotelServiceMock es una implementación mock del servicio de hoteles
type HotelServiceMock struct {
	hotels []dao.Hotel
}

// NewHotelServiceMock crea una nueva instancia del servicio de hoteles mock
func NewHotelServiceMock() *HotelServiceMock {
	// Inicializar los hoteles mock
	hotels := []dao.Hotel{
		{
			ID:   1,
			Name: "Hotel A",
		},
		{
			ID:   2,
			Name: "Hotel B",
		},
		{
			ID:   3,
			Name: "Hotel C",
		},
	}

	return &HotelServiceMock{
		hotels: hotels,
	}
}

// GetHotels devuelve la lista de hoteles mock
func (s *HotelServiceMock) GetHotels() ([]dao.Hotel, error) {
	return s.hotels, nil
}
