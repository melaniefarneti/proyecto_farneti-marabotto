package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-api/services"
)

func GetHotels(ctx *gin.Context) {
	// Llama al servicio para obtener el listado de hoteles
	hotels, err := services.GetHotels()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error getting hotels",
		})
		return
	}

	ctx.JSON(http.StatusOK, hotels)
}
