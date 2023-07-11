package clients

import (
	"errors"
	"fmt"
	"go-api/dao"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBClientInterface interface {
	CreateReservation(reservation dao.Reservation) error
	GetHotelByID(hotelID int) (dao.Hotel, error)
	CountReservations(hotelID int, checkin string, checkout string) (int, error)
	GetHotels() ([]dao.Hotel, error)
	CreateHotel(hotel *dao.Hotel) (*dao.Hotel, error)
	DeleteHotel(hotelID int) error

	GetUserByID(userID int) (*dao.User, error)
	GetUserByEmail(email string) (*dao.User, error)
	CreateUser(user *dao.User) (*dao.User, error)
	GetReservations() ([]dao.Reservation, error)

	GetReservationsByUserID(userID int) ([]*dao.Reservation, error)
	GetReservationsByHotelID(hotelID int) ([]*dao.Reservation, error)

	GetReservedRoomsByHotelIDAndDates(hotelID int, checkin, checkout string) (int, error)
	CreateAmenity(amenity *dao.Amenity) error
	CreateHotelPhoto(photo *dao.HotelPhoto) error

	GetAmenityByHotelID(hotelID int64) ([]dao.Amenity, error)
}

type DBClient struct {
	DB *gorm.DB
}

func NewDBClient() DBClient {
	// Configura la conexión a la base de datos
	dsn := "root:mel1@tcp(localhost:3306)/proyectohotel?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Obtiene la instancia de *sql.DB desde el objeto *gorm.DB
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	// Prueba la conexión
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&dao.User{}, &dao.Hotel{}, &dao.Reservation{})
	if err != nil {
		log.Fatal(err)
	}
	/*
		if err := db.AutoMigrate(&dao.User{}); err != nil {
			panic(err)
		}

		if err := db.AutoMigrate(&dao.Hotel{}); err != nil {
			panic(err)
		}

		if err := db.AutoMigrate(&dao.Reservation{}); err != nil {
			panic(err)
		}*/

	return DBClient{
		DB: db,
	}
}

func (c DBClient) CreateReservation(reservation dao.Reservation) error {
	if err := c.DB.Create(&reservation).Error; err != nil {
		return fmt.Errorf("error creating reservation: %w", err)
	}
	return nil
}

func (c DBClient) GetHotelByID(hotelID int) (dao.Hotel, error) {
	var hotel dao.Hotel
	query := c.DB.Where("id = ?", hotelID).First(&hotel)
	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return dao.Hotel{}, fmt.Errorf("Hotel with ID %d not found", hotelID)
		}
		return dao.Hotel{}, query.Error
	}
	return hotel, nil
}

// contar el número de reservas que se superponen con el rango de fechas especificado para un hotel dado
func (c DBClient) CountReservations(hotelID int, checkin string, checkout string) (int, error) {
	var count int64
	err := c.DB.Model(&dao.Reservation{}).
		Where("hotel_id = ?", hotelID).
		Where(c.DB.Where(c.DB.Where("checkin <= ?", checkin).Where("? <= checkin", checkin)).
			Or(c.DB.Where(c.DB.Where("checkin <= ?", checkout).Where("? <= checkin", checkout))).
			Or(c.DB.Where(c.DB.Where("? <= checkin", checkin).Where("checkin <= ?", checkout)))).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (c DBClient) GetHotels() ([]dao.Hotel, error) {
	var hotels []dao.Hotel
	err := c.DB.Model(&dao.Hotel{}).Find(&hotels).Error
	if err != nil {
		return nil, err
	}
	return hotels, nil
}

func (c DBClient) CreateHotel(hotel *dao.Hotel) (*dao.Hotel, error) {
	if err := c.DB.Create(hotel).Error; err != nil {
		return nil, fmt.Errorf("error creating hotel: %w", err)
	}
	return hotel, nil
}

func (c DBClient) DeleteHotel(hotelID int) error {
	var hotel dao.Hotel
	err := c.DB.Model(&dao.Hotel{}).Where("id = ?", hotelID).First(&hotel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("hotel not found")
		}
		return err
	}

	err = c.DB.Delete(&hotel).Error
	if err != nil {
		return fmt.Errorf("error deleting hotel: %w", err)
	}

	return nil
}

// metodos user
func (c DBClient) GetUserByID(userID int) (*dao.User, error) {
	var user dao.User
	err := c.DB.Model(&dao.User{}).Where("id = ?", userID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dao.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (c DBClient) GetUserByEmail(email string) (*dao.User, error) {
	var user dao.User
	err := c.DB.Model(&dao.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dao.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (c DBClient) CreateUser(user *dao.User) (*dao.User, error) {
	if err := c.DB.Create(user).Error; err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}
	return user, nil
}

func (c DBClient) GetReservations() ([]dao.Reservation, error) {
	var reservations []dao.Reservation
	err := c.DB.Model(&dao.Reservation{}).Find(&reservations).Error
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

// GetReservationsByUserID busca las reservas por ID de usuario en la base de datos
func (c DBClient) GetReservationsByUserID(userID int) ([]*dao.Reservation, error) {
	var reservations []*dao.Reservation
	err := c.DB.Model(&dao.Reservation{}).Where("user_id = ?", userID).Find(&reservations).Error
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

// GetReservationsByHotelID busca las reservas por ID de hotel en la base de datos
func (c DBClient) GetReservationsByHotelID(hotelID int) ([]*dao.Reservation, error) {
	var reservations []*dao.Reservation
	err := c.DB.Model(&dao.Reservation{}).Where("hotel_id = ?", hotelID).Find(&reservations).Error
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (c DBClient) GetReservedRoomsByHotelIDAndDates(hotelID int, checkin, checkout string) (int, error) {
	var count int64
	err := c.DB.Model(&dao.Reservation{}).
		Where("hotel_id = ?", hotelID).
		Where("checkin <= ?", checkout).
		Where("checkout >= ?", checkin).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (c DBClient) CreateAmenity(amenity *dao.Amenity) error {
	if err := c.DB.Create(amenity).Error; err != nil {
		return fmt.Errorf("error creating amenity: %w", err)
	}
	return nil
}

func (c DBClient) CreateHotelPhoto(photo *dao.HotelPhoto) error {
	if err := c.DB.Create(photo).Error; err != nil {
		return fmt.Errorf("error creating hotel photo: %w", err)
	}
	return nil
}

func (c DBClient) GetAmenityByHotelID(hotelID int64) ([]dao.Amenity, error) {
	var amenities []dao.Amenity
	err := c.DB.Model(&dao.Amenity{}).Where("hotel_id = ?", hotelID).Find(&amenities).Error
	if err != nil {
		return nil, err
	}
	return amenities, nil
}
