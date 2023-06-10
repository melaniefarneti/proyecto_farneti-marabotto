package controllers

import (
	"fmt"
	"go-api/domain"
	"go-api/services"
	"go-api/services/clients"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReservationController struct {
	ReservationService services.ReservationService
}

func NewReservationController(dbClient clients.DBClientInterface) *ReservationController {
	reservationService := services.ReservationService{
		DBClient: dbClient,
	}

	return &ReservationController{
		ReservationService: reservationService,
	}
}

func (c *ReservationController) CreateReservation(ctx *gin.Context) {
	// Obtener los parámetros del hotel, fechas de entrada y salida, y nombre del cliente del cuerpo de la solicitud
	var request domain.ReservationRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Verificar si el cliente existe en la base de datos
	_, err := c.ReservationService.DBClient.GetUserByEmail(request.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "client does not exist"})
		return
	}

	// Llamar al servicio para crear la reserva
	err = c.ReservationService.CreateReservation(request.HotelID, request.Checkin, request.Checkout, request.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error creating reservation: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "reservation created"})
}
