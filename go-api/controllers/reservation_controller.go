package controllers

import (
	"fmt"
	"go-api/clients"
	"go-api/dto"
	"go-api/services"
	"net/http"
	"strconv"

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
	var request dto.ReservationRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Verificar si el cliente existe en la base de datos
	_, err := c.ReservationService.DBClient.GetUserByEmail(request.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "client does not exist, must register"})
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

func (c *ReservationController) GetReservations(ctx *gin.Context) {
	reservations, err := c.ReservationService.GetReservations()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reservations)
}

// GetReservationsByUserID busca las reservas por ID de usuario
func (c *ReservationController) GetReservationsByUserID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userID"})
		return
	}

	reservations, err := c.ReservationService.GetReservationsByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get reservations"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"reservations": reservations})
}

func (c *ReservationController) GetReservationsByHotelID(ctx *gin.Context) {
	hotelID, err := strconv.Atoi(ctx.Param("hotelID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotelID"})
		return
	}

	reservations, err := c.ReservationService.GetReservationsByHotelID(hotelID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get reservations"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"reservations": reservations})
}
