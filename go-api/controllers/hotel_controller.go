package controllers

import (
	//"go-api/dao"
	"go-api/services"
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
)

func GetHotels(ctx *gin.Context) {
	hotelService := services.NewHotelService() // Crear una instancia del servicio de hoteles

	// Llama al servicio para obtener el listado de hoteles
	hotels, err := hotelService.GetHotels()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error getting hotels",
		})
		return
	}

	ctx.JSON(http.StatusOK, hotels)
}

/*
func CreateHotel(ctx *gin.Context) {
	var hotelRequest dao.Hotel
	if err := ctx.ShouldBindJSON(&hotelRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	hotelService := services.NewHotelService()
	hotel, err := hotelService.CreateHotel(&hotelRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create hotel",
		})
		return
	}

	ctx.JSON(http.StatusCreated, hotel)
}

func DeleteHotel(c *gin.Context) {
	// Obtener el ID del hotel desde los parámetros de la solicitud
	hotelID := c.Param("id")

	// Convertir el ID del hotel a entero
	id, err := strconv.Atoi(hotelID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hotel ID"})
		return
	}

	// Crear una instancia del servicio de hoteles
	hotelService := services.NewHotelService()

	// Llamar al servicio de hoteles para eliminar el hotel
	err = hotelService.DeleteHotel(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete hotel"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hotel deleted successfully"})
}
*/