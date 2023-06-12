package controllers

import (
	"go-api/dao"
	"go-api/services"
	"net/http"
	"strconv"

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

func CreateHotel(ctx *gin.Context) {
	// Parsear los datos del cuerpo de la solicitud JSON en una estructura Hotel
	var hotel dao.Hotel
	if err := ctx.ShouldBindJSON(&hotel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}

	// Crear una instancia del servicio de hoteles
	hotelService := services.NewHotelService()

	// Llamar al servicio para crear el hotel
	createdHotel, err := hotelService.CreateHotel(&hotel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error creating hotel",
		})
		return
	}

	ctx.JSON(http.StatusOK, createdHotel)
}

func DeleteHotel(ctx *gin.Context) {
	hotelIDStr := ctx.Param("hotelID")

	// Convertir hotelIDStr a un valor entero
	hotelID, err := strconv.Atoi(hotelIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid hotel ID",
		})
		return
	}

	hotelService := services.NewHotelService()
	err = hotelService.DeleteHotel(hotelID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error deleting hotel",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "hotel deleted successfully",
	})
}
