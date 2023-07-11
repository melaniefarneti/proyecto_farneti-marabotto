package services

import (
	"go-api/clients"
	"go-api/dao"
)

// AmenityServiceInterface define la interfaz para el servicio de amenities
type AmenityServiceInterface interface {
	CreateAmenity(amenity *dao.Amenity) error
	GetAmenityByHotelID(hotelID int64) ([]dao.Amenity, error)
}

// AmenityService es una implementación del servicio de amenities
type AmenityService struct {
	DBClient clients.DBClientInterface
}

func NewAmenityService() *AmenityService {
	dbClient := clients.NewDBClient() // Obtener una instancia válida del cliente de base de datos
	return &AmenityService{
		DBClient: dbClient,
	}
}

// CreateAmenity crea un nuevo amenity en la base de datos
func (s *AmenityService) CreateAmenity(amenity *dao.Amenity) error {
	// Llamar al cliente de base de datos para crear el amenity
	err := s.DBClient.CreateAmenity(amenity)
	if err != nil {
		return err
	}

	return nil
}

// GetAmenityByHotelID obtiene las amenidades por ID de hotel
func (s *AmenityService) GetAmenityByHotelID(hotelID int64) ([]dao.Amenity, error) {
	// Llamar al cliente de base de datos para obtener las amenidades por ID de hotel
	amenities, err := s.DBClient.GetAmenityByHotelID(hotelID)
	if err != nil {
		return nil, err
	}

	return amenities, nil
}
