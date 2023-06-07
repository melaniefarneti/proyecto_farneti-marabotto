package controllers

import (
	"net/http"

	"go-api/services"

	"github.com/gin-gonic/gin"
)

func GetHotels(ctx *gin.Context) {
	var hotelService services.HotelService
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
