package services

import (
	"go-api/clients"
	"go-api/dao"
	"go-api/dto"

	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
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

func (s *HotelService) UploadHotelPhoto(photoDTO dto.HotelPhoto, file multipart.File, header *multipart.FileHeader) error {
	// Generar un nombre único para el archivo
	filename := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(header.Filename))

	// Guardar el archivo en el sistema de archivos
	destinationPath := filepath.Join("uploads", filename)
	out, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	// Guardar el nombre del archivo en la base de datos junto con el ID del hotel
	photo := dao.HotelPhoto{
		HotelID:  photoDTO.HotelID,
		Filename: filename,
	}

	err = s.DBClient.CreateHotelPhoto(&photo)
	if err != nil {
		return err
	}

	return nil
}

func (s *HotelService) GetHotelPhotos(hotelID int) ([]dao.HotelPhoto, error) {
	// Llamar al cliente de base de datos para obtener las fotos del hotel
	photos, err := s.DBClient.GetHotelPhotos(hotelID)
	if err != nil {
		return nil, err
	}

	return photos, nil
}
