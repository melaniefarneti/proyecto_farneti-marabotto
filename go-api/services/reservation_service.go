package services

import (
	"errors"
	"go-api/domain"
)

var reservations []domain.Reservation

// GetReservationByID obtiene los datos de una reserva por su ID
func GetReservationByID(id int) (domain.Reservation, error) {
	for _, reservation := range reservations {
		if reservation.ID == id {
			return reservation, nil
		}
	}
	return domain.Reservation{}, domain.ErrReservationNotFound
}

// CreateReservation crea una nueva reserva
func CreateReservation(reservation domain.Reservation, user domain.User) (domain.Reservation, error) {
	if !user.IsRegistered() {
		return domain.Reservation{}, errors.New("user must be registered to make a reservation")
	}

	// Verificar disponibilidad de habitaciones
	availableRooms := getAvailableRooms(reservation.HotelID, reservation.CheckIn, reservation.CheckOut)
	if availableRooms < reservation.Rooms {
		return domain.Reservation{}, errors.New("not enough available rooms for the selected period")
	}

	// Asignar ID único a la reserva
	reservation.ID = generateReservationID()

	// Actualizar disponibilidad de habitaciones
	updateRoomAvailability(reservation.HotelID, reservation.CheckIn, reservation.CheckOut, reservation.Rooms)

	// Agregar reserva a la lista
	reservations = append(reservations, reservation)

	return reservation, nil
}

// GetReservationsByHotelAndDay obtiene las reservas filtradas por hotel y día para el administrador
func GetReservationsByHotelAndDay(hotelID int, date string) ([]domain.Reservation, error) {
	var filteredReservations []domain.Reservation
	for _, reservation := range reservations {
		if reservation.HotelID == hotelID && (reservation.CheckIn == date || reservation.CheckOut == date) {
			filteredReservations = append(filteredReservations, reservation)
		}
	}

	if len(filteredReservations) == 0 {
		return nil, domain.ErrReservationNotFound
	}

	return filteredReservations, nil
}

// getAvailableRooms retorna la cantidad de habitaciones disponibles para un hotel y período determinados
func getAvailableRooms(hotelID int, checkIn string, checkOut string) (int, error) {
	// Paso 1: Obtener el hotel por su ID desde la base de datos u otro sistema de almacenamiento
	hotel, err := GetHotelByID(hotelID)
	if err != nil {
		return 0, fmt.Errorf("error obteniendo el hotel: %w", err)
	}

	// Paso 2: Verificar la disponibilidad de habitaciones para el período especificado
	// Realiza una consulta en la base de datos para contar la cantidad de reservas existentes
	// para el hotel y el período especificados
	reservedRooms, err := CountReservedRooms(hotelID, checkIn, checkOut)
	if err != nil {
		return 0, fmt.Errorf("error contando las habitaciones reservadas: %w", err)
	}

	// Resta la cantidad de reservas al número total de habitaciones del hotel
	// para obtener la cantidad de habitaciones disponibles
	availableRooms := hotel.Rooms - reservedRooms

	// Paso 3: Retornar la cantidad de habitaciones disponibles
	return availableRooms, nil
}


// updateRoomAvailability actualiza la disponibilidad de habitaciones para un hotel y período determinados
func updateRoomAvailability(hotelID int, checkIn string, checkOut string, reservedRooms int) error {
	// 1. Obtener el hotel por su ID desde la base de datos u otro sistema de almacenamiento
	hotel, err := getHotelByID(hotelID)
	if err != nil {
		return fmt.Errorf("error obteniendo el hotel: %w", err)
	}

	// 2. Verificar la disponibilidad de habitaciones para el período especificado
	availableRooms := getAvailableRooms(hotelID, checkIn, checkOut)

	// 3. Actualizar la disponibilidad de habitaciones
	if availableRooms >= reservedRooms {
		// Hay suficientes habitaciones disponibles, se puede realizar la reserva
		updatedRooms := availableRooms - reservedRooms

		// 3.1. Actualizar la disponibilidad de habitaciones en el hotel
		hotel.AvailableRooms = updatedRooms

		// 3.2. Guardar los cambios en el hotel en la base de datos u otro sistema de almacenamiento
		err := updateHotel(hotel)
		if err != nil {
			return fmt.Errorf("error actualizando la disponibilidad de habitaciones: %w", err)
		}
	} else {
		// No hay suficientes habitaciones disponibles, la reserva no puede ser realizada
		return fmt.Errorf("no hay suficientes habitaciones disponibles para realizar la reserva")
	}

	return nil
}

// generateReservationID genera un ID único para una reserva
func generateReservationID() int {
	// Implementar la lógica para generar un ID único para la reserva, por ejemplo, incrementando un contador o utilizando un generador de IDs único
	// generamos un ID incrementando un contador
	return len(reservations) + 1
}
