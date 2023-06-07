package services

import (
	"go-api/services/clients"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateReservation(t *testing.T) {
	// Configurar el entorno de prueba
	token := "valid-token"
	hotelID := 1
	checkin := "2023-06-01"
	checkout := "2023-06-03"
	clientName := "John Doe"

	service := ReservationService{
		DBClient: clients.DBClientMock{},
	}

	// Ejecutar la función a probar
	err := service.CreateReservation(hotelID, checkin, checkout, token, clientName)

	// Verificar el resultado esperado
	assert.NoError(t, err, "error creating reservation should be nil")
}

func TestCalculateAvailableRooms(t *testing.T) {
	// Configurar el entorno de prueba
	hotelID := 1
	checkin := "2023-06-01"
	checkout := "2023-06-03"

	// Ejecutar la función a probar
	service := ReservationService{
		DBClient: clients.DBClientMock{},
	}
	availableRooms, err := service.calculateAvailableRooms(hotelID, checkin, checkout)

	// Verificar el resultado esperado
	assert.NoError(t, err, "error calculating available rooms should be nil")
	assert.Equal(t, 5, availableRooms, "available rooms should be 5")
}

func TestGetTotalRoomsFromDB(t *testing.T) {
	// Configurar el entorno de prueba
	hotelID := 1

	// Ejecutar la función a probar
	service := ReservationService{
		DBClient: clients.DBClientMock{},
	}
	totalRooms, err := service.getTotalRoomsFromDB(hotelID)

	// Verificar el resultado esperado
	assert.NoError(t, err, "error getting total rooms should be nil")
	assert.Equal(t, 5, totalRooms, "total rooms should be 5")
}
