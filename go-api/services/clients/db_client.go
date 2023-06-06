package clients

import (
	"fmt"
	"go-api/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DBClientInterface interface {
	CreateReservation(reservation dao.Reservation) error
	GetHotelByID(hotelID int) (dao.Hotel, error)
	CountReservations(hotelID int, checkin string, checkout string) (int, error)
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

	if err := db.AutoMigrate(&dao.User{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&dao.Amenity{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&dao.Hotel{}); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&dao.Reservation{}); err != nil {
		panic(err)
	}

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
	err := c.DB.Model(&dao.Hotel{}).Where("id = ?", hotelID).First(&hotel).Error
	if err != nil {
		return dao.Hotel{}, err
	}
	return hotel, nil
}

func (c DBClient) CountReservations(hotelID int, checkin string, checkout string) (int, error) {
	var count int64
	err := c.DB.Model(&dao.Reservation{}).
		Where("hotel_id = ?", hotelID).
		Where("checkin >= ? AND checkout <= ?", checkin, checkout).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
