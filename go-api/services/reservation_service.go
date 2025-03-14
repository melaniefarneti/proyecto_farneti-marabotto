package services

import (
	"errors"
	"fmt"
	"go-api/clients"
	"go-api/dao"
	"time"
)

type ServiceInterface interface {
	CreateReservation(hotelID int, checkin, checkout, clientName string) error
	GetReservationsByUserID(userID int) ([]*dao.Reservation, error)
	GetAvailableRoomsByHotelID(hotelID int, checkin, checkout string) (int, error)
}

type ReservationService struct {
	DBClient clients.DBClientInterface
}

func (s ReservationService) CreateReservation(hotelID int, checkin, checkout, clientName string, userID int) error {
	// Validar el formato de las fechas de check-in y check-out
	_, err := time.Parse("2006-01-02", checkin)
	if err != nil {
		return fmt.Errorf("invalid check-in date format: %w", err)
	}

	_, err = time.Parse("2006-01-02", checkout)
	if err != nil {
		return fmt.Errorf("invalid check-out date format: %w", err)
	}

	// Validar la disponibilidad de habitaciones del hotel para las fechas especificadas
	availableRooms, err := s.calculateAvailableRooms(hotelID, checkin, checkout)
	if err != nil {
		return err
	}

	// Verificar si hay suficientes habitaciones disponibles
	if availableRooms <= 0 {
		return errors.New("no available rooms for the specified dates")
	}

	// Parsea las fechas de check-in y check-out
	checkInDate, err := time.Parse("2006-01-02", checkin)
	if err != nil {
		return fmt.Errorf("error parsing check-in date: %w", err)
	}

	checkOutDate, err := time.Parse("2006-01-02", checkout)
	if err != nil {
		return fmt.Errorf("error parsing check-out date: %w", err)
	}

	// Crea una nueva instancia de la reserva
	reservation := dao.Reservation{
		HotelID:    hotelID,
		CheckIn:    checkInDate.Format("2006-01-02"),
		CheckOut:   checkOutDate.Format("2006-01-02"),
		ClientName: clientName,
		UserID:     userID,
	}

	// Realiza las operaciones necesarias para almacenar la reserva en la base de datos
	if err := s.DBClient.CreateReservation(reservation); err != nil {
		return fmt.Errorf("error creating reservation: %w", err)
	}

	return nil
}

func (s ReservationService) calculateAvailableRooms(hotelID int, checkin string, checkout string) (int, error) {
	// Obtener la cantidad total de habitaciones del hotel
	totalRooms, err := s.getTotalRoomsFromDB(hotelID)
	if err != nil {
		return 0, err
	}

	// Obtener la cantidad de habitaciones reservadas para el rango de fechas dado
	reservedRooms, err := s.DBClient.CountReservations(hotelID, checkin, checkout)
	if err != nil {
		return 0, err
	}

	// Calcular la cantidad de habitaciones disponibles
	availableRooms := totalRooms - reservedRooms

	if availableRooms < 0 {
		return 0, errors.New("no available rooms for the specified dates")
	}

	return availableRooms, nil
}

func (s ReservationService) getTotalRoomsFromDB(hotelID int) (int, error) {
	hotel, err := s.DBClient.GetHotelByID(hotelID)
	if err != nil {
		return 0, err
	}

	/*
		rooms := hotel.Rooms
		if err != nil {
			return 0, fmt.Errorf("error converting rooms to int: %w", err)
		}*/

	return hotel.Rooms, nil
}

func (s *ReservationService) GetReservations() ([]dao.Reservation, error) {
	reservations, err := s.DBClient.GetReservations()
	if err != nil {
		return nil, err
	}

	return reservations, nil
}

// GetReservationsByUserID busca las reservas por ID de usuario
func (s *ReservationService) GetReservationsByUserID(userID int) ([]*dao.Reservation, error) {
	reservations, err := s.DBClient.GetReservationsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

// GetReservationsByHotelID busca las reservas por ID de hotel
func (s *ReservationService) GetReservationsByHotelID(hotelID int) ([]*dao.Reservation, error) {
	reservations, err := s.DBClient.GetReservationsByHotelID(hotelID)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (s ReservationService) GetAvailableRoomsByHotelID(hotelID int, checkin, checkout string) (int, error) {
	// Validar el formato de las fechas de check-in y check-out
	_, err := time.Parse("2006-01-02", checkin)
	if err != nil {
		return 0, fmt.Errorf("invalid check-in date format: %w", err)
	}

	_, err = time.Parse("2006-01-02", checkout)
	if err != nil {
		return 0, fmt.Errorf("invalid check-out date format: %w", err)
	}

	// Obtener la cantidad total de habitaciones del hotel
	totalRooms, err := s.getTotalRoomsFromDB(hotelID)
	if err != nil {
		return 0, err
	}

	// Obtener la cantidad de habitaciones reservadas para el hotel y las fechas especificadas
	reservedRooms, err := s.DBClient.GetReservedRoomsByHotelIDAndDates(hotelID, checkin, checkout)
	if err != nil {
		return 0, err
	}

	// Calcular la cantidad de habitaciones disponibles
	availableRooms := totalRooms - reservedRooms

	if availableRooms < 0 {
		return 0, errors.New("no available rooms for the specified hotel and dates")
	}

	return availableRooms, nil
}
