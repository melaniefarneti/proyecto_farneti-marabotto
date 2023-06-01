package controllers

import (
	"go-api/services"
	"gorm.io/gorm"
	"net/http"

	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
)

func CreateReservation(ctx *gin.Context) {
	// Obtener el token del encabezado de autorización
	token := ctx.GetHeader("Authorization")

	// Validar el token
	if !isValidToken(token) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	// Obtener los demás parámetros del hotel, fechas de entrada y salida del cuerpo de la solicitud
	var request struct {
		HotelID  int    `json:"hotel_id"`
		Checkin  string `json:"checkin"`
		Checkout string `json:"checkout"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// Obtener la instancia de la base de datos desde el contexto
	db := ctx.MustGet("db").(*gorm.DB)

	// Llamar al servicio para crear la reserva
	err := services.CreateReservation(db, request.HotelID, request.Checkin, request.Checkout)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error creating reservation"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "reservation created"})
}
