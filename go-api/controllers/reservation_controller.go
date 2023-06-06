package controllers

import (
	"fmt"
	"go-api/domain"
	"go-api/services"
	"go-api/services/clients"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
)

func CreateReservation(ctx *gin.Context) {
	// Obtener el token del encabezado de autorización
	token := ctx.GetHeader("Authorization")

	// Validar el token
	// if !isValidToken(token) {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
	// 	return
	// }

	// Obtener los demás parámetros del hotel, fechas de entrada y salida, y nombre del cliente del cuerpo de la solicitud
	var request domain.ReservationRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	// client
	dbClient := clients.NewDBClient()

	// Llamar al servicio para crear la reserva
	service := services.ReservationService{
		DBClient: dbClient,
	}
	err := service.CreateReservation(request.HotelID, request.Checkin, request.Checkout, token, request.ClientName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error creating reservation: %s", err.Error())})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "reservation created"})
}

func isValidToken(tokenString string) bool {
	// Paso 1: Define la estructura de la clave secreta
	var secretKey = []byte("mi-clave-secreta")

	// Paso 2: Parsea y valida el token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifica el método de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Devuelve la clave secreta para validar la firma del token
		return secretKey, nil
	})

	// Paso 3: Verifica si hubo algún error durante el parsing o la validación del token
	if err != nil {
		return false
	}

	// Paso 4: Verifica si el token es válido
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}

	return false
}
