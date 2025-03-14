package clients

import (
	"errors"
	"fmt"
	"go-api/dao"
)

type DBClientMock struct {
	GetUserByIDFunc                       func(userID int) (*dao.User, error)
	GetUserByEmailFunc                    func(email string) (*dao.User, error)
	CreateUserFunc                        func(user *dao.User) (*dao.User, error)
	CreateReservationFunc                 func(reservation dao.Reservation) error
	CountReservationsFunc                 func(hotelID int, checkin string, checkout string) (int, error)
	GetReservationsFunc                   func() ([]dao.Reservation, error)
	GetHotelByIDFunc                      func(hotelID int) (dao.Hotel, error)
	GetReservationsByUserIDFunc           func(userID int) ([]*dao.Reservation, error)
	GetReservationsByHotelIDFunc          func(hotelID int) ([]*dao.Reservation, error)
	GetReservedRoomsByHotelIDAndDatesFunc func(hotelID int, checkin, checkout string) (int, error)
	CreateAmenityFunc                     func(amenity *dao.Amenity) error
	GetAmenityByHotelIDFunc               func(hotelID int64) ([]dao.Amenity, error)
}

func (DBClientMock) CreateReservation(reservation dao.Reservation) error {
	fmt.Println(fmt.Sprintf("Creating reservation: %v", reservation))
	return nil
}

func (DBClientMock) GetHotelByID(hotelID int) (dao.Hotel, error) {
	fmt.Println(fmt.Sprintf("Getting hotel by ID: %d", hotelID))
	return dao.Hotel{
		ID:          1,
		Name:        "Test Hotel",
		Photo:       "Test Photo.jpg",
		Description: "Test Description",
		Location:    "Test Location",
		Rooms:       10,
	}, nil
}

func (DBClientMock) CountReservations(hotelID int, checkin string, checkout string) (int, error) {
	fmt.Println(fmt.Sprintf("Counting reservations for hotel: %d", hotelID))
	return 5, nil
}

func (DBClientMock) GetHotels() ([]dao.Hotel, error) {
	fmt.Println("Getting hotels")
	hotels := []dao.Hotel{
		{
			ID:          1,
			Name:        "Hotel 1",
			Photo:       "photo1.jpg",
			Description: "Description 1",
			Location:    "Location 1",
			Rooms:       10,
		},
		{
			ID:          2,
			Name:        "Hotel 2",
			Photo:       "photo2.jpg",
			Description: "Description 2",
			Location:    "Location 2",
			Rooms:       20,
		},
	}
	return hotels, nil
}

func (m DBClientMock) CreateHotel(hotel *dao.Hotel) (*dao.Hotel, error) {
	hotel.ID = 1 // Asignar un ID ficticio para el hotel
	return hotel, nil
}

func (m DBClientMock) DeleteHotel(hotelID int) error {
	//verificar si el hotelID es igual a un valor específico y devolver un error simulado en caso contrario.
	if hotelID != 123 {
		return fmt.Errorf("hotel with ID %d does not exist", hotelID)
	}

	// Si el hotelID es válido, se ha eliminado correctamente y devolver nil (sin error).
	return nil
}

// metodos user
func (m DBClientMock) GetUserByID(userID int) (*dao.User, error) {
	if m.GetUserByIDFunc != nil {
		return m.GetUserByIDFunc(userID)
	}
	return nil, fmt.Errorf("GetUserByIDFunc not implemented")
}

func (m DBClientMock) GetUserByEmail(email string) (*dao.User, error) {
	if m.GetUserByEmailFunc != nil {
		return m.GetUserByEmailFunc(email)
	}
	return nil, fmt.Errorf("GetUserByEmailFunc not implemented")
}

func (m DBClientMock) CreateUser(user *dao.User) (*dao.User, error) {
	if m.CreateUserFunc != nil {
		return m.CreateUserFunc(user)
	}
	return nil, errors.New("not implemented")
}

func (m *DBClientMock) GetReservations() ([]dao.Reservation, error) {
	if m.GetReservationsFunc != nil {
		return m.GetReservationsFunc()
	}
	return nil, errors.New("not implemented")
}

func (m *DBClientMock) GetReservationsByUserID(userID int) ([]*dao.Reservation, error) {
	//if m.GetReservationsByUserID != nil {
	//	return m.GetReservationsByUserID(userID)
	//}
	return nil, errors.New("not implemented")
}

func (m *DBClientMock) GetReservationsByHotelID(hotelID int) ([]*dao.Reservation, error) {
	//if m.GetReservationsByHotelID != nil {
	//	return m.GetReservationsByHotelID(hotelID)
	//}
	return nil, errors.New("not implemented")
}

func (m *DBClientMock) GetReservedRoomsByHotelIDAndDates(hotelID int, checkin, checkout string) (int, error) {
	if m.GetReservedRoomsByHotelIDAndDatesFunc != nil {
		return m.GetReservedRoomsByHotelIDAndDatesFunc(hotelID, checkin, checkout)
	}
	return 0, errors.New("not implemented")
}

func (m *DBClientMock) CreateAmenity(amenity *dao.Amenity) error {
	if m.CreateAmenityFunc != nil {
		return m.CreateAmenityFunc(amenity)
	}
	return errors.New("not implemented")
}

func (DBClientMock) CreateHotelPhoto(photo *dao.HotelPhoto) error {
	fmt.Printf("Creando foto para el ID del hotel %d\n", photo.HotelID)
	fmt.Println("Foto creada exitosamente")
	return nil
}

func (m *DBClientMock) GetAmenityByHotelID(hotelID int64) ([]dao.Amenity, error) {
	if m.GetAmenityByHotelIDFunc != nil {
		return m.GetAmenityByHotelIDFunc(hotelID)
	}
	return nil, errors.New("not implemented")
}

func (m *DBClientMock) GetHotelPhotos(hotelID int) ([]dao.HotelPhoto, error) {
	// Aquí puedes proporcionar la lógica de simulación para obtener las fotos del hotel
	// según el ID del hotel.
	// Por ejemplo, puedes devolver un conjunto de fotos predefinidas para el hotel con el ID proporcionado.

	photos := []dao.HotelPhoto{
		{HotelID: 1, Filename: "photo1.jpg"},
		{HotelID: 1, Filename: "photo2.jpg"},
		{HotelID: 1, Filename: "photo3.jpg"},
	}

	return photos, nil
}
