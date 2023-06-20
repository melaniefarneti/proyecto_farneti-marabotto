package controllers

import (
	"errors"
	"go-api/clients"
	"go-api/dao"
	"go-api/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateReservation(t *testing.T) {
	// Crear un enrutador Gin y configurar el controlador
	router := gin.Default()
	dbClientMock := &clients.DBClientMock{}
	reservationServiceMock := &services.ReservationServiceMock{
		DBClientMock: dbClientMock,
		CreateReservationFunc: func(reservation dao.Reservation) error {
			return nil // Simulación de creación exitosa de reserva
		},
	}
	//reservationController := NewReservationController(reservationServiceMock)
	//router.POST("/reservations", reservationController.CreateReservation)

	// Preparar una solicitud HTTP de prueba
	requestBody := strings.NewReader(`{"hotelID": 1, "checkin": "2023-06-21", "checkout": "2023-06-23", "email": "john.doe@example.com"}`)
	req, _ := http.NewRequest("POST", "/reservations", requestBody)
	req.Header.Set("Content-Type", "application/json")

	// Caso de prueba 1: Creación exitosa de reserva
	reservationServiceMock.CreateReservationFunc = func(reservation dao.Reservation) error {
		return nil // Simulación de creación exitosa de reserva
	}

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code, "Código de estado de respuesta incorrecto")

	// Caso de prueba 2: Error en la creación de reserva
	reservationServiceMock.CreateReservationFunc = func(reservation dao.Reservation) error {
		return errors.New("error creating reservation") // Simulación de error en la creación de reserva
	}

	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code, "Código de estado de respuesta incorrecto")
}
