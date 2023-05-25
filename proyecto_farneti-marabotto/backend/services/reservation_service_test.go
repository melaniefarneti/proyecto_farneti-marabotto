package services

import (
	"errors"
	"fmt"
	"go-api/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReservationByID(t *testing.T) {
	reservations = []domain.Reservation{
		{ID: 1, HotelID: 1, CheckIn: "2023-06-01", CheckOut: "2023-06-05", Rooms: 2},
		{ID: 2, HotelID: 2, CheckIn: "2023-06-01", CheckOut: "2023-06-03", Rooms: 1},
	}

	// Existing reservation
	reservation, err := GetReservationByID(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, reservation.ID)

	// Non-existing reservation
	_, err = GetReservationByID(3)
	assert.NotNil(t, err)
	assert.Equal(t, domain.ErrReservationNotFound, err)
}

func TestCreateReservation(t *testing.T) {
	reservations = []domain.Reservation{
		{ID: 1, HotelID: 1, CheckIn: "2023-06-01", CheckOut: "2023-06-05", Rooms: 2},
	}

	user := domain.User{Registered: true}

	// Valid reservation
	reservation := domain.Reservation{HotelID: 2, CheckIn: "2023-06-01", CheckOut: "2023-06-03", Rooms: 1}
	createdReservation, err := CreateReservation(reservation, user)
	assert.Nil(t, err)
	assert.NotZero(t, createdReservation.ID)
	assert.Equal(t, reservation.HotelID, createdReservation.HotelID)
	assert.Equal(t, reservation.CheckIn, createdReservation.CheckIn)
	assert.Equal(t, reservation.CheckOut, createdReservation.CheckOut)
	assert.Equal(t, reservation.Rooms, createdReservation.Rooms)

	// Invalid reservation with unregistered user
	reservation = domain.Reservation{HotelID: 1, CheckIn: "2023-06-01", CheckOut: "2023-06-05", Rooms: 2}
	_, err = CreateReservation(reservation, domain.User{})
	assert.NotNil(t, err)
	assert.Equal(t, "user must be registered to make a reservation", err.Error())

	// Invalid reservation with not enough available rooms
	reservation = domain.Reservation{HotelID: 1, CheckIn: "2023-06-01", CheckOut: "2023-06-05", Rooms: 3}
	_, err = CreateReservation(reservation, user)
	assert.NotNil(t, err)
	assert.Equal(t, "not enough available rooms for the selected period", err.Error())
}

func TestGetReservationsByHotelAndDay(t *testing.T) {
	reservations = []domain.Reservation{
		{ID: 1, HotelID: 1, CheckIn: "2023-06-01", CheckOut: "2023-06-05", Rooms: 2},
		{ID: 2, HotelID: 2, CheckIn: "2023-06-01", CheckOut: "2023-06-03", Rooms: 1},
		{ID: 3, HotelID: 1, CheckIn: "2023-06-05", CheckOut: "2023-06-08", Rooms: 1},
	}

	// Existing reservations for a hotel and date
	filteredReservations, err := GetReservationsByHotelAndDay(1, "2023-06-01")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(filteredReservations))

	// Non-existing reservations for a hotel and date
	_, err = GetReservationsByHotelAndDay(2, "2023-06-02")
	assert.NotNil(t, err)
	assert.Equal(t, domain.ErrReservationNotFound, err)
}

func TestGetAvailableRooms(t *testing.T) {
	// Mocking the GetHotelByID function
	getHotelByID = func(id int) (domain.Hotel, error) {
		if id == 1 {
			return domain.Hotel{ID: 1, Rooms: 10}, nil
		}
		return domain.Hotel{}, errors.New("hotel not found")
	}

	// Mocking the CountReservedRooms function
	countReservedRooms = func(hotelID int, checkIn string, checkOut string) (int, error) {
		if hotelID == 1 && checkIn == "2023-06-01" && checkOut == "2023-06-05" {
			return 2, nil
		}
		return 0, fmt.Errorf("error counting reserved rooms")
	}

	// Existing hotel and reserved rooms
	availableRooms, err := getAvailableRooms(1, "2023-06-01", "2023-06-05")
	assert.Nil(t, err)
	assert.Equal(t, 8, availableRooms)

	// Non-existing hotel
	_, err = getAvailableRooms(2, "2023-06-01", "2023-06-05")
	assert.NotNil(t, err)
	assert.Equal(t, "error obteniendo el hotel: hotel not found", err.Error())

	// Error counting reserved rooms
	_, err = getAvailableRooms(1, "2023-06-02", "2023-06-05")
	assert.NotNil(t, err)
	assert.Equal(t, "error contando las habitaciones reservadas: error counting reserved rooms", err.Error())
}

func TestUpdateRoomAvailability(t *testing.T) {
	// Mocking the getHotelByID function
	getHotelByID = func(id int) (domain.Hotel, error) {
		if id == 1 {
			return domain.Hotel{ID: 1, Rooms: 10, AvailableRooms: 10}, nil
		}
		return domain.Hotel{}, errors.New("hotel not found")
	}

	// Mocking the getAvailableRooms function
	getAvailableRooms = func(hotelID int, checkIn string, checkOut string) (int, error) {
		if hotelID == 1 && checkIn == "2023-06-01" && checkOut == "2023-06-05" {
			return 10, nil
		}
		return 0, fmt.Errorf("error getting available rooms")
	}

	// Mocking the updateHotel function
	updateHotel = func(hotel domain.Hotel) error {
		return nil
	}

	// Valid case: enough available rooms
	err := updateRoomAvailability(1, "2023-06-01", "2023-06-05", 2)
	assert.Nil(t, err)

	// Non-existing hotel
	err = updateRoomAvailability(2, "2023-06-01", "2023-06-05", 2)
	assert.NotNil(t, err)
	assert.Equal(t, "error obteniendo el hotel: hotel not found", err.Error())

	// Invalid case: not enough available rooms
	err = updateRoomAvailability(1, "2023-06-01", "2023-06-05", 11)
	assert.NotNil(t, err)
	assert.Equal(t, "no hay suficientes habitaciones disponibles para realizar la reserva", err.Error())

	// Error getting available rooms
	err = updateRoomAvailability(1, "2023-06-02", "2023-06-05", 2)
	assert.NotNil(t, err)
	assert.Equal(t, "error contando las habitaciones reservadas: error getting available rooms", err.Error())

	// Error updating hotel
	updateHotel = func(hotel domain.Hotel) error {
		return errors.New("error updating hotel")
	}
	err = updateRoomAvailability(1, "2023-06-01", "2023-06-05", 2)
	assert.NotNil(t, err)
	assert.Equal(t, "error actualizando la disponibilidad de habitaciones: error updating hotel", err.Error())
}

func TestGenerateReservationID(t *testing.T) {
	reservations = []domain.Reservation{
		{ID: 1, HotelID: 1, CheckIn: "2023-06-01", CheckOut: "2023-06-05", Rooms: 2},
		{ID: 2, HotelID: 2, CheckIn: "2023-06-01", CheckOut: "2023-06-03", Rooms: 1},
	}

	// Existing reservations
	reservationID := generateReservationID()
	assert.Equal(t, 3, reservationID)

	// No existing reservations
	reservations = nil
	reservationID = generateReservationID()
	assert.Equal(t, 1, reservationID)
}
