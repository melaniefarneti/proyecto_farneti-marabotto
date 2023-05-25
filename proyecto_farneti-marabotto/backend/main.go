package main

import (
	"fmt"
	"go-api/handlers"
	"go-api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar los servicios
	hotelService := services.NewHotelService()
	reservationService := services.NewReservationService()
	amenityService := services.NewAmenityService()

	// Crear el enrutador Gin
	router := gin.Default()

	// Definir las rutas y los manejadores
	router.GET("/hotels", handlers.GetHotels(hotelService))
	router.GET("/reservations/:id", handlers.GetReservationByID(reservationService))
	router.POST("/reservations", handlers.CreateReservation(reservationService))
	router.GET("/reservations/hotel/:hotelID/date/:date", handlers.GetReservationsByHotelAndDay(reservationService))
	router.GET("/amenities", handlers.GetAmenities(amenityService))

	// Iniciar el servidor HTTP
	port := 8080
	serverAddress := fmt.Sprintf(":%d", port)
	log.Printf("Server listening on %s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, router))
}
