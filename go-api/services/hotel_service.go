package services

import (
	"go-api/dao"
	"go-api/services/clients"
)

// HotelServiceInterface define la interfaz para el servicio de hoteles
type HotelServiceInterface interface {
	GetHotels() ([]dao.Hotel, error)
}

// HotelService es una implementación del servicio de hoteles
type HotelService struct {
	DBClient clients.DBClientInterface
}

// GetHotels devuelve la lista de hoteles desde la base de datos
func (s HotelService) GetHotels() ([]dao.Hotel, error) {
	// Llamar al cliente de base de datos para obtener los hoteles
	hotels, err := s.DBClient.GetHotels()
	if err != nil {
		return nil, err
	}

	return hotels, nil
}
