package services

import (
	"go-api/dao"
	"go-api/dto"
)

// HotelServiceMock es una implementación simulada del servicio de hoteles para pruebas.
type HotelServiceMock struct {
	Hotels []dto.HotelRequest
}

// NewHotelServiceMock crea una nueva instancia de HotelServiceMock.
func NewHotelServiceMock() *HotelServiceMock {
	return &HotelServiceMock{
		Hotels: []dto.HotelRequest{},
	}
}

// GetHotels devuelve el listado de hoteles simulados.
func (m *HotelServiceMock) GetHotels() ([]dto.HotelRequest, error) {
	return m.Hotels, nil
}

// CreateHotel simula la creación de un hotel y lo agrega a la lista simulada.
func (m *HotelServiceMock) CreateHotel(hotel *dao.Hotel) (*dto.HotelRequest, error) {
	newHotel := dto.HotelRequest{
		Name:        hotel.Name,
		Photo:       hotel.Photo,
		Description: hotel.Description,
		Location:    hotel.Location,
		Rooms:       hotel.Rooms,
	}

	m.Hotels = append(m.Hotels, newHotel)
	return &newHotel, nil
}
