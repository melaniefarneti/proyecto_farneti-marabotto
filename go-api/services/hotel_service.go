package services

import (
	"go-api/clients"
	"go-api/dao"
	"go-api/dto"
)

// HotelServiceInterface define la interfaz para el servicio de hoteles
type HotelServiceInterface interface {
	GetHotels() ([]dao.Hotel, error)
	CreateHotel(hotel *dao.Hotel) (*dao.Hotel, error)
	DeleteHotel(hotelID int) error
}

// HotelService es una implementación del servicio de hoteles
type HotelService struct {
	DBClient clients.DBClientInterface
}

// GetHotels devuelve la lista de hoteles desde la base de datos
func (s *HotelService) GetHotels() ([]dao.Hotel, error) {
	// Llamar al cliente de base de datos para obtener los hoteles
	hotels, err := s.DBClient.GetHotels()
	if err != nil {
		return nil, err
	}

	return hotels, nil
}

// CreateHotel crea un nuevo hotel en la base de datos
func (s *HotelService) CreateHotel(hotel *dao.Hotel) (*dao.Hotel, error) {
	// Llamar al cliente de base de datos para crear el hotel
	createdHotel, err := s.DBClient.CreateHotel(hotel)
	if err != nil {
		return nil, err
	}

	return createdHotel, nil
}

func NewHotelService() *HotelService {
	dbClient := clients.NewDBClient() // Obtener una instancia válida del cliente de base de datos
	return &HotelService{
		DBClient: dbClient,
	}
}

func (s *HotelService) DeleteHotel(hotelID int) error {
	// Llamar al cliente de base de datos para eliminar el hotel
	err := s.DBClient.DeleteHotel(hotelID)
	if err != nil {
		return err
	}

	return nil
}

func (s *HotelService) UploadHotelPhoto(photoDTO dto.HotelPhoto) error {
	// Crear una instancia del DAO HotelPhoto
	photo := dao.HotelPhoto{
		HotelID:  photoDTO.HotelID,
		Filename: photoDTO.Filename,
	}

	// Llamar al cliente de base de datos para crear la foto del hotel
	err := s.DBClient.CreateHotelPhoto(&photo)
	if err != nil {
		return err
	}

	return nil
}
